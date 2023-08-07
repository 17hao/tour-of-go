package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var v1 atomic.Value
	fmt.Printf("%v", v1.Load())
}
