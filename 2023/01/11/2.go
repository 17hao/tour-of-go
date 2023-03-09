package main

import (
	"errors"
	"fmt"
	"sync"
)

func f4() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()

		errChan <- errors.New("err A")
	}()

	go func() {
		defer wg.Done()

		errChan <- errors.New("err B")
	}()

	wg.Wait()
	close(errChan)

	for e := range errChan {
		fmt.Println(e.Error())
	}
}

func f5() {
	//errChan := make(chan string)
	errChan := make(chan string, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		println("before")
		errChan <- "err"
		println("after")
	}()

	wg.Wait()
	close(errChan)

	select {
	case err := <-errChan:
		println(err)
	default:
		println("success")
		return
	}
}
