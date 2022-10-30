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
	wg.Add(2)
	go accumulate()
	go accumulate()
	wg.Wait()
	wg.Add(1)
	accumulate()
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
