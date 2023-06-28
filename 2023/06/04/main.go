package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 10

const nums = 1000

var sig = false

func f1() string {
	defer func() {
		wg.Wait()
		sig = true
		fmt.Print("done\n")
		fmt.Printf("sig=%v\n", sig)
	}()
	done := make(chan bool, n)
	for i := 0; i < n; i++ {
		num := nums / n * (i + 1)
		go func(max int, done chan<- bool) {
			defer func() {
				wg.Done()
			}()
			wg.Add(1)

			fmt.Print(time.Now())
			fmt.Println("== before ==")
			fmt.Printf("max=%d\n", max)
			time.Sleep(5 * time.Second)
			fmt.Print(time.Now())
			fmt.Println("== after ==")
			// done <- true
		}(num, done)
	}
	return "doing sth"
}

var wg = sync.WaitGroup{}

func f2() {
	minID := 0
	maxID := 9999
	min := minID
	for i := 0; i < n; i++ {
		var max int
		if i == n-1 {
			max = maxID
		} else {
			max = (maxID-minID)/n*(i+1) + minID
		}
		fmt.Printf("min=%d max=%d\n", min, max)

		go func(min, max int) {
			time.Sleep(2 * time.Second)
		}(min, max)

		min = max + 1
	}
}

func main() {
	// res := f1()
	// fmt.Println(res)
	// time.Sleep(20 * time.Second)
	// fmt.Print(time.Now())
	// fmt.Println("main exit")
	f2()
}
