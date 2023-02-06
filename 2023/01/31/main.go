package main

import "fmt"

func main() {
	m := map[int64]int8{}

	i := m[1]
	fmt.Printf("%d", i) // output: 0
}
