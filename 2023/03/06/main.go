package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		demo()
		fmt.Printf("number of goroutines=%d\n", runtime.NumGoroutine())
	}
}

func demo() {
	//ch := make(chan bool)
	ch := make(chan bool, 0)
	fmt.Printf("cap(ch)=%d len(ch)=%d\n", cap(ch), len(ch))

	go func(ch chan<- bool) {
		time.Sleep(1 * time.Second)
		ch <- true
		fmt.Printf("cap(ch)=%d len(ch)=%d\n", cap(ch), len(ch))
	}(ch)

	go func(ch chan<- bool) {
		time.Sleep(1 * time.Second)
		ch <- true
		fmt.Printf("cap(ch)=%d len(ch)=%d\n", cap(ch), len(ch))
	}(ch)

	go func(ch chan<- bool) {
		time.Sleep(1 * time.Second)
		ch <- true
		fmt.Printf("cap(ch)=%d len(ch)=%d\n", cap(ch), len(ch))
	}(ch)

	<-ch
}
