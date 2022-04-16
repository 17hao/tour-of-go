package main

import "sort"

func main() {
	nums := []int{1, 4, 10, 1, 0, 3}
	sort.Ints(nums)
	for _, n := range nums {
		print(n, " ")
	}
}
