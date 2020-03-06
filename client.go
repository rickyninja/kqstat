// Package kqstat is used to connect to a Killerqueen stats service, and read events as they occur.
package kqstat

import (
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rickyninja/kqstat/event"
)

// aliveResp is the response sent back to the stats service during a keep alive event.
const aliveResp = "![k[im alive],v[]]!"

type Logger interface {
	Logf(format string, a ...interface{})
}

// Client is a connection to a stats service.
type Client struct {
	*websocket.Conn
	wmutex *sync.Mutex
	rmutex *sync.Mutex
	log    Logger
}

// NewClient connects to a stats service, and returns a *Client.
func NewClient(addr string, l Logger) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: addr}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	c := &Client{
		Conn:   ws,
		wmutex: new(sync.Mutex),
		rmutex: new(sync.Mutex),
		log:    l,
	}
	return c, nil
}

// ReadMessage wraps websocket.Conn.ReadMessage with a mutex.
func (c *Client) ReadMessage() (int, []byte, error) {
	c.rmutex.Lock()
	defer c.rmutex.Unlock()
	return c.Conn.ReadMessage()
}

// WriteMessage wraps websocket.Conn.WriteMessage with a mutex.
func (c *Client) WriteMessage(messageType int, data []byte) error {
	c.wmutex.Lock()
	defer c.wmutex.Unlock()
	return c.Conn.WriteMessage(messageType, data)
}

// GetEvent returns the next event from the stats service.
func (c *Client) GetEvent() (event.Event, error) {
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	ev, err := event.Parse(string(message))
	if err != nil {
		return ev, err
	}
	// Auto reply to keep alives as a convenience, while still allowing the caller to see the event.
	if _, ok := ev.(event.Alive); ok {
		go func() {
			err := c.WriteMessage(websocket.TextMessage, []byte(aliveResp))
			if err != nil {
				c.log.Logf("%s", err)
			}
		}()
	}
	return ev, nil
}

// Close does a graceful close of the websocket connection.
func (c *Client) Close() error {
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	return nil
}
