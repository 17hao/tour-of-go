package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	name := "world"
	c.Write([]byte(name))
	fmt.Println(c) // 5
	c = 0
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, world")
}
