package main

import (
	"io"
	"log"
	"net"
	"os"
)

// netcat2 has 2 goroutines, the main goroutine reads the standard input and sends it to the server,
// a second goroutine reads and prints the server's response.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(conn, os.Stdout)
	mustCopy(os.Stdin, conn)
}

func mustCopy(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
