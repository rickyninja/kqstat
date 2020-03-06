package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

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
	logger := newMylog(log.New(os.Stderr, "", 0))
	flag.Parse()
	cl, err := kqstat.NewClient(net.JoinHostPort(host, port), logger)
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

type mylog struct {
	*log.Logger
}

func newMylog(l *log.Logger) *mylog {
	return &mylog{l}
}

func (l *mylog) Logf(format string, a ...interface{}) {
	l.Printf(format, a...)
}
