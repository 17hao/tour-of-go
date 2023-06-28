package main

import (
	"encoding/json"
	"fmt"
)

type s1 struct {
	Key1 string
	Key2 int64
}

type s2 struct {
	Key1 string
	Key2 int64
}

func main() {
	a := s1{
		Key1: "key1",
		Key2: int64(1),
	}
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))

	b := s2{}
	_ = json.Unmarshal(bytes, &b)
	fmt.Printf("%+v", b)
}
