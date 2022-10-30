package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	client()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Dial failed: ", err.Error())
	}
	fmt.Fprint(conn, "test\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)
}
