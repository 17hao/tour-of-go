package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("1", "0.1"))
	fmt.Println(strings.Compare("1", "1.1"))
	fmt.Println(strings.Compare("1", "3"))
	fmt.Println(strings.Compare("-1", "0.11"))
	fmt.Println("===")
	nums := []string{"0.11", "-1", "1"}
	//asc := false
	asc := true
	sort.SliceStable(nums, func(i, j int) bool {
		// nums[i] < nums[j]  asc          nums[i] < nums[j] != asc
		// true				  true		   false
		// false			  true		   true
		// true				  false	       true
		// false			  false 	   false
		return !(strings.Compare(nums[i], nums[j]) < 0 != asc)
	})
	for _, n := range nums {
		fmt.Println(n)
	}

	fmt.Println("compare: ", strings.Compare("10", "6"))
}
