package main

import "fmt"

func main() {
	var strings []string
	strings = append(strings, "hello")
	for str := range strings {
		fmt.Println(str)
	}
}
