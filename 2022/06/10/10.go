package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "bytedance.com:80")
	if err != nil {
		log.Fatal(err)
	}
	buffer := bufio.NewReader(conn)
	status, err := buffer.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)
}
