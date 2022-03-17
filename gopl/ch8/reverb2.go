package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8888")
	for {
		conn, _ := listener.Accept()
		handleConn4(conn)
	}
}

func echo2(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn4(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo2(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
