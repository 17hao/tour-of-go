package main

import "fmt"

func newNums() []int64 {
	return nil
}

func main() {
	nums := make([]int64, 0)
	nums = append(nums, int64(0))
	nums = append(nums, newNums()...)
	fmt.Printf("%+v", nums)
}
