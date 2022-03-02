package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"testing"
)

func Test_handleConnection(t *testing.T) {
	srv, cl := net.Pipe()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		handleConnection(srv)
		srv.Close()
		wg.Done()
	}()

	reader := bufio.NewReader(cl)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	if len(b) < 5 {
		t.Error("no message received")
	}
	wg.Wait()
	cl.Close()
}
