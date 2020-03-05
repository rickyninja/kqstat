//kqfakestatd simulates a Killerqueen stats service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/rickyninja/kqstat/mock/kqstatd"
)

func main() {
	var (
		host     string
		port     string
		statfile string
	)
	flag.StringVar(&host, "host", "", "host of the killerqueen stat service")
	flag.StringVar(&port, "port", "12749", "port of the killerqueen stat service")
	flag.StringVar(&statfile, "statfile", "", "stat file to simulate killerqueen stat service")
	flag.Parse()
	if statfile == "" {
		fmt.Fprintln(os.Stderr, "statfile is required")
		flag.Usage()
		os.Exit(1)
	}
	fd, err := os.Open(statfile)
	if err != nil {
		log.Fatalf("Failed to open %s: %s", statfile, err)
	}
	defer fd.Close()
	replay, err := kqstatd.NewReplay(fd)
	if err != nil {
		log.Fatal("Failed NewReplay: ", err)
	}
	http.ListenAndServe(net.JoinHostPort(host, port), replay)
}
