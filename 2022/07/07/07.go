package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	s := []string{"v", "x", "y", "u", "v"}
	res := reverse(s)
	for _, r := range res {
		println(r)
	}
}

func reverse(s []string) []string {
	fmt.Println(string(debug.Stack()))
	l := len(s)
	for i := 0; i < l/2; i++ {
		tmp := s[i]
		s[i] = s[l-1-i]
		s[l-1-i] = tmp
	}
	return s
}
