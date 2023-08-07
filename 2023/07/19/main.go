package main

import (
	"fmt"
)

func f1() bool {
	i := 1
	switch i {
	case 1:
		fmt.Printf("%d", i)
		return true
	default:
		fmt.Printf("default")
		return false
	}
}

func main() {
	f1()
}
