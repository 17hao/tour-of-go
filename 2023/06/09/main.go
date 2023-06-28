package main

import (
	"fmt"
	"reflect"
)

type s struct {
	id   int64
	name string
}

func main() {
	x := s{
		id:   1,
		name: "name",
	}
	y := s{
		id:   1,
		name: "name",
	}
	fmt.Printf("%v", reflect.DeepEqual(x, y))
}
