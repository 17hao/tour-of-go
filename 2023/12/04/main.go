package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <- ticker.C:
			fmt.Printf("%s ticker.C\n", time.Now().String())
		default:
			fmt.Printf("%s default\n", time.Now().String())
			time.Sleep(500 * time.Millisecond)
		}
	}
}