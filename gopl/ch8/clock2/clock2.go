package main

import (
	"io"
	"log"
	"net"
	"time"
)

// Clock2 is a concurrent TCP server, it deals with two or more clients at a time.
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
		go handleConn(c)
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
