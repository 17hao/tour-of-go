package main

import "strings"

func main() {
	input := "one two three"
	words := strings.Fields(input)
	for _, word := range words {
		println(word)
	}
}
