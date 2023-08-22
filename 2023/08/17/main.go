package main

import (
	"encoding/json"
	"fmt"
)

type s1 struct {
	F1 string
	F2 int64
	F3 int32
}

type s2 s1

func main() {
	m := s1{
		F1: "a",
		F2: int64(1),
	}
	var n s2
	b, _ := json.Marshal(m)
	_ = json.Unmarshal(b, &n)
	fmt.Printf("m=%+v\n", string(b))
	fmt.Printf("n=%+v\n", n)
}
