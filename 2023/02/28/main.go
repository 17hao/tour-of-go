package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	m := map[string]string{"key": "v1"}
	doSomething(m)
	fmt.Printf("%+v\n", m)

	for i := 0; i < 5; i++ {
		res := queryAll()
		fmt.Println(res)
		fmt.Printf("number of goroutines: %d\n", runtime.NumGoroutine())
	}
}

func queryAll() string {
	resp := make(chan string, 3)
	//resp := make(chan string) // goroutine leak

	go func() {
		str := query("goroutine-1", 2)
		resp <- str
	}()

	go func() {
		str := query("goroutine-2", 1)
		resp <- str
	}()

	go func() {
		str := query("goroutine-3", 3)
		resp <- str
	}()

	//fmt.Printf(<-resp)
	return <-resp
}

func query(input string, n int) string {
	time.Sleep(time.Duration(n) * time.Second)
	return input
}

func doSomething(m map[string]string) {
	m["key"] = "do something"
}
