package main

import (
	"fmt"
	"regexp"
)

func main() {
	// match, err := regexp.Match("[0-9]+", []byte("123456"))
	// fmt.Printf("match=%v err=%v", match, err)

	regex := regexp.MustCompile("[0-9]+")

	match := regex.MatchString("6108ae92-f3af-4c55-8590-8836d95c9368")
	fmt.Printf("match=%v", match)
}
