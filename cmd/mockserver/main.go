package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rickyninja/kqstat"
	"github.com/rickyninja/kqstat/mock/server"
)

func main() {
	fd, err := os.Open("./mock/server/output.txt")
	must(err)
	defer fd.Close()
	server, err := server.NewServer(fd)
	must(err)
	go server.ListenAndServe()
	client, err := kqstat.NewClient(server.Addr().String())
	for {
		event := <-client.EventChan
		fmt.Printf("event: %#v\n", event)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
