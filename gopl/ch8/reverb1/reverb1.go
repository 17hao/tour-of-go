package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Echo server simulates the reverberation of a real echo.
func main() {
	listener, _ := net.Listen("tcp", "localhost:8000")
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

	// naive implementation
	//for {
	//	buf := make([]byte, 100)
	//	c.Read(buf)
	//	echo(c, string(buf), 1*time.Second)
	//}

	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
