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
	//r := bufio.NewReader(conn)
	for {
		buf := make([]byte, 4096)
		//n, err := r.Read(buf)
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
		pairs, err := parseKV(buf)
		if err != nil {
			log.Fatalf("Failed to parseKV: %s\n", err)
		}
		//r.Reset(conn)
		for _, p := range pairs {
			key := p.Key
			val := p.Value
			fmt.Printf("key: %s value: %s\n", key, val)
			if key == "alive" {
				//resp := []byte("im alive")
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
		//fmt.Printf("buf before key begin: %s\n", string(buf))
		kb := bytes.Index(buf, []byte("![k["))
		if kb == -1 {
			fmt.Println("did not find key begin, break")
			break
			//return nil, fmt.Errorf("Failed to find start of key: %s\n", string(buf))
		}
		kb += 4 // skip past ![k[
		//fmt.Printf("buf before key end: %s\n", string(buf))
		ke := bytes.Index(buf, []byte("],v["))
		if ke == -1 {
			return nil, fmt.Errorf("Failed to find end of key: %s\n", string(buf))
		}
		vb := ke + 4 // skip past ],v[
		//fmt.Printf("buf before value end: %s\n", string(buf))
		ve := bytes.Index(buf, []byte("]]!"))
		if ve == -1 {
			return nil, fmt.Errorf("Failed to find end of value: %s\n", string(buf))
		}
		// Read 27 bytes from server -> ![k[alive],v[10:24:42 PM]]!
		// 2019/01/31 22:24:42 Failed to parseKV: Failed to find end of value: ![k[alive],v[10:24:42 PM
		key := string(buf[kb:ke])
		value := string(buf[vb:ve])
		//fmt.Printf("buf before reslice: %s\n", string(buf))
		buf = buf[ve+3:]
		//fmt.Printf("buf after reslice: %s\n", string(buf))
		//fmt.Printf("buf after reslice to value end len %d >%s<\n", len(buf), string(buf))
		pairs = append(pairs, Pair{key, value})
	}
	return pairs, nil
}
