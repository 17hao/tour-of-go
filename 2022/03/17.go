package main

import "fmt"

func main() {
	nums := []int{1, 2}
	nums = append(nums, nums...)
	for _, i := range nums {
		fmt.Println(i)
	}
}