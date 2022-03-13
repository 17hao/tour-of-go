package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"1": 1}
	strings := []string{"1", "2"}

	for _, str := range strings {
		if key, ok := m[str]; ok {
			fmt.Println(key)
		} else {
			fmt.Println(ok)
		}
	}

	sliceDemo()
}

func sliceDemo() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = 2022 - i
	}
	printSlice(s)
	//modifySlice(s)
	s = append(s, 1)
	printSlice(s)
}

func printSlice(s []int) {
	for _, i := range s {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func modifySlice(s []int) []int {
	return append(s, 1)

	//s = append(s, 1)
	//s[0] = 1
	//return s

	//s[3] = 1
	//return s
}
