package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	chunkSize := 3
	var divided [][]int
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		divided = append(divided, nums[i:end])
	}
	fmt.Println(divided)
}
