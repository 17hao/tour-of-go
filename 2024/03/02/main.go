package main

import (
	"fmt"
)

type testStruct struct {
	m map[string]string
}

func (s *testStruct) set(k string, v string) {
	if s.m == nil {
		s.m = map[string]string{}
	}
	s.m[k] = v
}

func main() {
	i := &testStruct{}
	i.set("k", "v")
	fmt.Printf("%+v", i)
}