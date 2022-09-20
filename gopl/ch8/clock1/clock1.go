package main

import (
	"io"
	"log"
	"net"
	"time"
)

// Clock1 is a TCP server that periodically writes the time to client.
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Print("Accept remote: ", c.RemoteAddr().String())
		handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02T15:04:05Z07:00\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
