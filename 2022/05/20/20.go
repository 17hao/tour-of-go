package main

import "fmt"

type byteSlice []byte

func (p *byteSlice) Write(data []byte) (int, error) {
	slice := *p
	*p = slice
	return len(data), nil
}

func main() {
	var b byteSlice
	fmt.Fprintf(&b, "hello, world")
	fmt.Println(string(b))
}
