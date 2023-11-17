package main

import (
	"fmt"
	"sort"
)

type structA struct {
	id int64
	name string
}

func main() {
	s := []structA{
		{id: 1, name: "1"},
		{id: 3, name: "3"},
		{id: 2, name: "2"},
	}
	sort.Slice(s, func(i, j int) bool {
		// return s[i].id < s[j].id
		return s[i].id > s[j].id
	})
	fmt.Printf("%+v\n", s)
}
