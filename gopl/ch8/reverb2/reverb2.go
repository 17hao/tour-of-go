package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// The echo server in reverb1 doesn't deal with second shout until last shout was petered out.
// A real echo would consist of the composition of 3 independent shouts.
func main() {
	listener, _ := net.Listen("tcp", "localhost:8888")
	for {
		conn, _ := listener.Accept()
		handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
