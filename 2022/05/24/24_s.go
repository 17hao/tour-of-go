package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server()
}

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("server started.")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			log.Println(conn.RemoteAddr().String())
			content, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			log.Print(content)
			fmt.Fprint(conn, content)
		}(conn)
	}
}
