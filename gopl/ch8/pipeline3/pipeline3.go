package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 3; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}

// chan<- send-only
// <-chan receive-only
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	// close(in) compile error
	close(out)
}

func printer(in <-chan int) {
	//for x := range in {
	//	fmt.Println(x)
	//}

	//for {
	//	x, ok := <-in
	//	if ok {
	//		fmt.Println(x)
	//	} else {
	//		return
	//	}
	//}

	for {
		fmt.Println(<-in)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
