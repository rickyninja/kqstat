package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rickyninja/kqstat"
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
	kq, err := kqstat.NewClient(host, port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		event := <-kq.EventChan
		fmt.Printf("event: %#v\n", event)
	}
}
