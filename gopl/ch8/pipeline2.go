package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
