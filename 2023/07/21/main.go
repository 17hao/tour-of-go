package main

import (
	"fmt"
	"time"
)

func f1(i int64) {
	fmt.Printf("%s i=%d\n", time.Now().String(), i)
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for i := 0; i < 10; i++ {
		select {
		// case <-done:
		// 	fmt.Println("Done!")
		// 	return
		case <-ticker.C:
			f1(int64(i))
		}
	}
}
