package main

import (
	"fmt"
)

func main() {
	m := make(map[string]byte, 0)
	m["k1"] = 1
	m["k2"] = 1
	fmt.Printf("%+v", m)
}
