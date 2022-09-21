package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("k1", "v1")
	fmt.Println(m.Load("k1"))
}
