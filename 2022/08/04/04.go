package main

import "fmt"

func main() {
	//str := "string"
	type s struct {
		id string
	}
	i := s{
		id: "string",
	}
	fmt.Printf("%s", &i)
	fmt.Printf("%+v", &i)
}
