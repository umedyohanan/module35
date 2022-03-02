package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func getRandomProverb() string {
	rand.Seed(time.Now().Unix())
	proverbs := [19]string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}

	return proverbs[rand.Intn(len(proverbs))]
}

func handleConnection(conn net.Conn) {
	//defer conn.Close()
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for i := 0; i < 19; i++ {
		conn.Write([]byte(getRandomProverb() + "\n"))
		time.Sleep(3 * time.Second)
	}
	conn.Close()
}

func main() {
	tcpNetwork := "tcp4"
	addrPort := "0.0.0.0:12345"
	listener, err := net.Listen(tcpNetwork, addrPort)

	if err != nil {
		log.Panic("error: ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic("error: ", err)
		}

		go handleConnection(conn)
	}

}
