package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// go run string.go
// go tool trace trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("hello, world!")
}
