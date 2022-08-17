package main

import (
	"fmt"
)

func main() {
	str := "string"
	fmt.Println(str[1 : len(str)-1])

	nums := []int{1, 2, 3, 4, 5, 6}
	//l := math.Min(float64(2), float64(len(nums)))
	//nums = nums[:int(l)]
	//for _, r := range nums {
	//	fmt.Println(r)
	//}

	fmt.Println("====")

	location := 7
	//nums = append(nums[:location-1], append([]int{1000}, nums[location-1:]...)...)
	nums = append(nums[:location], nums[location-1:]...)
	nums[location-1] = 1000
	for _, n := range nums {
		fmt.Println(n)
	}
}
