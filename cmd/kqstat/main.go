package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	port string
	host string
)

func init() {
	flag.StringVar(&port, "port", "12749", "Killerqueen stats service port")
	flag.StringVar(&host, "host", "localhost", "Killerqueen stats service host")
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		fmt.Printf("conn.Read %d bytes err: %v\n", n, err)
		if err != nil {
			if err == io.EOF {
				log.Fatalln("EOF on Read, probably sent something bogus to the server.")
			}
			log.Fatalf("Failed to Read into buf: %s\n", err)
		}
		if n == 0 {
			log.Println("Read zero bytes, sleep 100ms and keep going")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		fmt.Printf("Read %d bytes from server -> %s\n", n, string(buf))
		pairs, err := parseKV(buf[:n])
		if err != nil {
			log.Fatalf("Failed to parseKV: %s\n", err)
		}
		for _, p := range pairs {
			key := p.Key
			val := p.Value
			fmt.Printf("key: %s value: %s\n", key, val)
			if key == "alive" {
				resp := []byte("![k[im alive],v[]]!")
				fmt.Printf("got %s send %s\n", key, string(resp))
				_, err = conn.Write(resp)
				if err != nil {
					log.Fatalf("Failed to ack alive message: %s\n", err)
				}
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
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
			fmt.Printf("did not find key begin, buf %v, break\n", buf)
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
			fmt.Println("buffer has been emptied of k/v pairs, break")
			break
		}
	}
	return pairs, nil
}
