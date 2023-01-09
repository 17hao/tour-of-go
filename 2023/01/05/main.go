package main

import "fmt"

type sub struct {
	i int
}

type parent struct {
	s *sub
}

func main() {
	p := parent{}
	p.s.i = 1
	fmt.Printf("%v\n", p)
}
