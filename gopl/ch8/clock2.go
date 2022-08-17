package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn)
	}
}

func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.RFC3339+"\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
