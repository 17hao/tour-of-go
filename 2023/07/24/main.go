package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(i int64) {
	fmt.Printf("%s i=%d\n", time.Now().String(), i)
}

func main() {
	nums := make([]*int64, 0)
	for i := 0; i < 10; i++ {
		i64 := int64(i)
		nums = append(nums, &i64)
	}

	concurrentSize := 5
	threshold := make(chan struct{}, concurrentSize)
	wg := sync.WaitGroup{}
	wg.Add(10)

	for _, n := range nums {
		threshold <- struct{}{}
		go func(i *int64) {
			defer wg.Done()
			f1(*i)
			// time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			time.Sleep(2 * time.Second)
			<-threshold
		}(n)
	}
	wg.Wait()
}
