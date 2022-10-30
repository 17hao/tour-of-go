package main

import (
	"io"
	"log"
	"net"
	"os"
)

// netcat1 is a TCP client that transfer data from connection to stdout.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(conn, os.Stdout)
}

func mustCopy(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
