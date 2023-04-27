package main

import (
	"fmt"
)

var v = "v0"

func f1() {
	v = "v1"
	fmt.Println(v)
}

func f2() {
	v = "v2"
	fmt.Println(v)
}

func main() {
	f1()
	f2()
}
