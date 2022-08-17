package main

import "fmt"

type s struct {
	id int
	m  map[string]string
}

func main() {
	ss := make([]s, 0)
	ss = append(ss, s{id: 1, m: map[string]string{"k": "v"}})

	tmp1 := ss[0]
	fmt.Printf("before tmp1: %+v\n", tmp1)
	tmp1.id = 2
	tmp1.m["k"] = "tmp1"
	fmt.Printf("after tmp1: %+v\n", tmp1)
	fmt.Printf("ss[0]: %+v\n", ss[0])

	fmt.Println(tmp1.m["k2"])
}
