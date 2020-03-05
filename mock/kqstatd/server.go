package kqstatd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Replay replays stats from an io.Reader on repeat.
type Replay struct {
	*websocket.Conn
	reader          *bufio.Reader
	buffer          []byte
	failedKeepAlive chan struct{}
	wmutex          *sync.Mutex
	rmutex          *sync.Mutex
}

// NewReplay constructs a *Replay object using an io.Reader as its input source.
func NewReplay(r io.Reader) (*Replay, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	re := &Replay{
		buffer:          buf,
		failedKeepAlive: make(chan struct{}),
		wmutex:          new(sync.Mutex),
		rmutex:          new(sync.Mutex),
	}
	re.resetReader()
	return re, nil
}

// ServeHTTP does http.Handler.
func (r *Replay) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	r.Conn = c
	done := make(chan struct{})
	defer func() { done <- struct{}{} }()
	go r.doKeepAlives(done)
	for {
		select {
		case <-r.failedKeepAlive:
			return
		default:
			line, err := r.reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					r.resetReader()
					if len(line) < 1 {
						continue
					}
				} else {
					log.Println(err)
					return
				}
			}
			fmt.Printf("writing line: <%s>\n", string(line))
			err = r.WriteMessage(websocket.TextMessage, line)
			if err != nil {
				log.Println("write: ", err)
				return
			}
		}
	}
}

func (r *Replay) resetReader() {
	cp := make([]byte, len(r.buffer))
	n := copy(cp, r.buffer)
	if n < len(r.buffer) {
		log.Printf("Failed to copy buffer: got %d want %d\n", n, len(r.buffer))
		return
	}
	r.reader = bufio.NewReader(bytes.NewBuffer(cp))
}

// ReadMessage wraps websocket.Conn.ReadMessage with a mutex.
func (r *Replay) ReadMessage() (int, []byte, error) {
	r.rmutex.Lock()
	defer r.rmutex.Unlock()
	return r.Conn.ReadMessage()
}

// WriteMessage wraps websocket.Conn.WriteMessage with a mutex.
func (r *Replay) WriteMessage(messageType int, data []byte) error {
	r.wmutex.Lock()
	defer r.wmutex.Unlock()
	return r.Conn.WriteMessage(messageType, data)
}

func (r *Replay) doKeepAlives(done <-chan struct{}) {
	ticker := time.NewTicker(600 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := r.sendKeepAlive()
			if err != nil {
				log.Println("sendKeepAlive: ", err)
				return
			}
			_, message, err := r.ReadMessage()
			if err != nil {
				log.Println("ReadMessage: ", err)
				return
			}
			if !bytes.Equal(message, []byte(keepAliveMsg)) {
				log.Println("Did not receive expected keep alive response, closing connection.")
				r.failedKeepAlive <- struct{}{}
			}
		}
	}
}

func (r *Replay) sendKeepAlive() error {
	now := time.Now()
	ts := now.Format("15:04:05")
	message := fmt.Sprintf("![k[alive],v[%s]!", ts)
	err := r.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

const keepAliveMsg = "![k[im alive],v[]]!"
