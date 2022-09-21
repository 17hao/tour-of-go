package main

import (
	"fmt"
	"testing"
)

func TestFoo26(t *testing.T) {
	fmt.Println("foo")
	var err error
	if &err != nil {
		panic(err)
	}
	fmt.Println("bar")
}
