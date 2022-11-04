package main

import (
	"os"
	"time"
)

func main() {
	println("Commencing countdown.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			println("abort.")
			return
		}
	}
	println("finish.")
}
