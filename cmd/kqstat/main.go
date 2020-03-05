package main

import (
	"flag"
	"fmt"
	"log"
	"net"

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
	cl, err := kqstat.NewClient(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		ev, err := cl.GetEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", ev)
	}
}
