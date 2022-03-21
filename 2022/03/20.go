package main

import (
	"log"
	"sync"
)

var (
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	wg.Add(3)
	go accumulate()
	go accumulate()
	accumulate()
	wg.Wait()
	log.Println(count)
}

func accumulate() {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	for i := 0; i < 100000000; i++ {
		count++
	}
}
