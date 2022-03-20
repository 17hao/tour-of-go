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
	go accumulate()
	go accumulate()
	accumulate()
	wg.Wait()
	log.Println(count)
}

func accumulate() {
	wg.Add(1)
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		count++
	}
}
