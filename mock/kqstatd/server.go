package kqstatd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Logger interface {
	Logf(format string, a ...interface{})
}

// Replay replays stats from an io.Reader on repeat.
type Replay struct {
	buffer []byte
	log    Logger
}

// NewReplay constructs a *Replay object using an io.Reader as its input source.
func NewReplay(r io.Reader, l Logger) (*Replay, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	re := &Replay{
		buffer: buf,
		log:    l,
	}
	return re, nil
}

// ServeHTTP does http.Handler.
func (r *Replay) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	ws, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		r.log.Logf("websocket Upgrade: %s", err)
		return
	}
	con := newConn(ws, r.log)
	defer con.Close()
	done := make(chan struct{})
	defer func() { done <- struct{}{} }()
	go con.doKeepAlives(done)
	reader := r.getReader()
	for {
		select {
		case <-con.failedKeepAlive:
			return
		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					reader = r.getReader()
					if len(line) < 1 {
						continue
					}
				} else {
					r.log.Logf("%s", err)
					return
				}
			}
			err = con.WriteMessage(websocket.TextMessage, line)
			if err != nil {
				r.log.Logf("websocket WriteMessage: %s", err)
				return
			}
		}
	}
}

// getReader returns a copy of the original input as an io.Reader for each connection.
func (r *Replay) getReader() *bufio.Reader {
	cp := make([]byte, len(r.buffer))
	n := copy(cp, r.buffer)
	if n < len(r.buffer) {
		r.log.Logf("Failed to copy buffer: got %d want %d\n", n, len(r.buffer))
	}
	return bufio.NewReader(bytes.NewBuffer(cp))
}

// conn manages the websocket connection.
type conn struct {
	*websocket.Conn
	failedKeepAlive chan struct{}
	wmutex          *sync.Mutex
	rmutex          *sync.Mutex
	log             Logger
}

// newConn creates a *conn from a websocket.
func newConn(ws *websocket.Conn, l Logger) *conn {
	return &conn{
		Conn:            ws,
		failedKeepAlive: make(chan struct{}),
		wmutex:          new(sync.Mutex),
		rmutex:          new(sync.Mutex),
		log:             l,
	}
}

// ReadMessage wraps websocket.Conn.ReadMessage with a mutex.
func (c *conn) ReadMessage() (int, []byte, error) {
	c.rmutex.Lock()
	defer c.rmutex.Unlock()
	return c.Conn.ReadMessage()
}

// WriteMessage wraps websocket.Conn.WriteMessage with a mutex.
func (c *conn) WriteMessage(messageType int, data []byte) error {
	c.wmutex.Lock()
	defer c.wmutex.Unlock()
	return c.Conn.WriteMessage(messageType, data)
}

// doKeepAlives sends keep alive messages, and expects the proper response from the peer.
func (c *conn) doKeepAlives(done <-chan struct{}) {
	ticker := time.NewTicker(600 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.sendKeepAlive()
			if err != nil {
				c.log.Logf("sendKeepAlive: %s", err)
				return
			}
			_, message, err := c.ReadMessage()
			if err != nil {
				c.log.Logf("ReadMessage: %s", err)
				return
			}
			if !bytes.Equal(message, []byte(keepAliveMsg)) {
				c.log.Logf("Did not receive expected keep alive response, closing connection.")
				c.failedKeepAlive <- struct{}{}
			}
		}
	}
}

// sendKeepAlive sends a keep alive message to its peer.
func (c *conn) sendKeepAlive() error {
	now := time.Now()
	ts := now.Format("15:04:05")
	message := fmt.Sprintf("![k[alive],v[%s]!", ts)
	err := c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

const keepAliveMsg = "![k[im alive],v[]]!"
