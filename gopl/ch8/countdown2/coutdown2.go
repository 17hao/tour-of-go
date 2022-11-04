package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte from standard input
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.    Press return to abort.")
	select {
	case <-time.After(5 * time.Second):
		println("finish")
	case <-abort:
		fmt.Println("abort")
		return
	}
}
