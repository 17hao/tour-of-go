package main

import "fmt"

type structA struct {
	id   int64
	name *string
}

func (s structA) ToString() string {
	return fmt.Sprintf("id=%d\tname=%s", s.id, *s.name)
}

func main() {
	ss := make([]structA, 0)
	strs := []string{"str1", "str2", "str3"}
	for i, str := range strs {
		ss = append(ss, structA{id: int64(i), name: &str})
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s.ToString())
	}
}
