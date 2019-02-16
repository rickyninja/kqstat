package kqstat

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func NewClient(host, port string) (*Client, error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:      conn,
		EventChan: make(chan Pair),
	}
	go c.statLoop()
	return c, nil
}

func (c *Client) statLoop() {
	log.SetOutput(os.Stderr)
	for {
		buf := make([]byte, 4096)
		n, err := c.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("Failed to Read into buf: %s\n", err)
			continue
		}
		pairs, err := parseKV(buf[:n])
		if err != nil {
			log.Printf("Failed to parseKV: %s\n", err)
			continue
		}
		for _, p := range pairs {
			key := p.Key
			if key == "alive" {
				resp := []byte("![k[im alive],v[]]!")
				_, err = c.conn.Write(resp)
				if err != nil {
					log.Printf("Failed to ack alive message: %s\n", err)
				}
			} else {
				c.EventChan <- p
			}
		}
	}
}

type Client struct {
	conn      net.Conn
	EventChan chan Pair
}

type Pair struct {
	Key   string
	Value string
}

func parseKV(buf []byte) ([]Pair, error) {
	pairs := make([]Pair, 0)
	for {
		kb := bytes.Index(buf, []byte("![k["))
		if kb == -1 {
			log.Printf("did not find key begin, buf %s\n", string(buf))
			break
		}
		kb += 4 // skip past ![k[
		ke := bytes.Index(buf, []byte("],v["))
		if ke == -1 {
			return nil, fmt.Errorf("Failed to find end of key: %s\n", string(buf))
		}
		vb := ke + 4 // skip past ],v[
		ve := bytes.Index(buf, []byte("]]!"))
		if ve == -1 {
			return nil, fmt.Errorf("Failed to find end of value: %s\n", string(buf))
		}
		key := string(buf[kb:ke])
		value := string(buf[vb:ve])
		buf = buf[ve+3:]
		pairs = append(pairs, Pair{key, value})
		if len(buf) == 0 {
			break
		}
	}
	return pairs, nil
}
