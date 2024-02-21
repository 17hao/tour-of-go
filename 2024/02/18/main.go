package main

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

func f3(input []int) {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
}

func f2() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

func f1(minID, maxID int64) {
	batchSize := 10
	gap := (maxID - minID) / int64(batchSize)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < batchSize; i++ {
		num := i
		min := minID + gap*int64(num)
		max := minID + gap*int64(num+1)
		if max > maxID {
			max = maxID
		}
		go func(x, y int64) {
			defer wg.Done()
			logrus.Infof("min=%d, max=%d, num=%d", x, y, num)
		}(min, max)
		min = max + 1
	}
	wg.Wait()
}

func main() {
	//f1(10, 1000)
	input := []int{0, 1, 2, 3, 4, 5}
	//input := []int{0, 1, 2, 3, 4}
	f3(input)
	fmt.Println(input)
}
