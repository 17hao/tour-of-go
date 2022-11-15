package main

import (
	"fmt"
	"sync"
)

func main() {
	ids := []int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	for i, id := range ids {
		wg.Add(1)
		go func(i int, id int) {
			defer wg.Done()
			fmt.Printf("i=%d, &id=%d\n", i, &id)
		}(i, id)

		wg.Add(1)
		go func(i int, id int) {
			defer wg.Done()
			fmt.Printf("i=%d, id=%d\n", i, id)
		}(i, id)
	}
	wg.Wait()
}
