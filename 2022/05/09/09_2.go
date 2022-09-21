package main

import (
	"math/rand"
	"time"
)

func main() {
	ch := await()
	println(<-ch)
}

func await() <-chan int32 {
	ch := make(chan int32)
	go func() {
		defer close(ch)
		time.Sleep(3 * time.Second)
		rand.Seed(time.Now().UnixNano())
		ch <- rand.Int31n(100)
	}()
	return ch
}
