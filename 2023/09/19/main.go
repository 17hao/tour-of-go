package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

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
		go func(int64, int64) {
			defer wg.Done()
			logrus.Infof("min=%d, max=%d, num=%d", min, max, num+1)
		}(min, max)
	}
	wg.Wait()
}

func main() {
	f1(1, 1000)
}
