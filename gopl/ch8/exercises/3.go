package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 8.4.1
//
// The client program in netcat1 and netcat2 copies the input and sends to the server in main goroutine, the program
// terminates as soon as the input stream closes, even if the background goroutine is still running.
//
// To make the program wait for the background goroutine to complete before existing, we use a channel to synchronize
// two goroutines.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(os.Stdin, conn)
	conn.(*net.TCPConn).CloseWrite() //
	//conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}
