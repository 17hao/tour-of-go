package main

import (
	"encoding/json"
	"fmt"
)

type s struct {
	A int64
	B *string
}

func f1() {
	str := "{\"A\":1}\n"
	var s1 s
	_ = json.Unmarshal([]byte(str), &s1)
	fmt.Printf("f1:%+v\n", s1)
}

func main() {
	f1()

	str := "test"
	s1 := &s{A: 1, B: &str}
	b, _ := json.Marshal(&s1)
	fmt.Printf("%s\n", string(b))

	s2 := []*s{s1}

	fmt.Printf("s1=%+v s2=%+v", s1, s2)
}
