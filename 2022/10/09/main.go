package main

import (
	"fmt"
)

func main() {
	//calculate()
	ch := make(chan int)

	go put(ch)
	print(ch)
	//for {
	//	fmt.Println(<-ch)
	//}
}

func put(in chan<- int) {
	in <- 1
	in <- 2
	close(in)
}

func print(out <-chan int) {
	for {
		if i, ok := <-out; ok {
			fmt.Println(i)
		} else {
			return
		}
	}
}

func calculate() {
	// c and quit are bidirectional.
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
