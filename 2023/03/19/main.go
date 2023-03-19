package main

import (
	"fmt"
)

func main() {
	strs := make([]string, 0)

	f1(strs)
	for _, str := range strs {
		fmt.Printf("%s ", str)
	}

	f2(&strs)
	for _, str := range strs {
		fmt.Printf("%s ", str)
	}
	fmt.Println()

	// ====

	nums1 := make([][]int, 0)

	f3(&nums1)
	for _, row := range nums1 {
		for _, n := range row {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}

	nums2 := make([][]int, 0)
	f4(&nums2)
	for _, row := range nums2 {
		for _, n := range row {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}

func f1(strs []string) {
	strs = append(strs, "f1")
}

func f2(pstrs *[]string) {
	*pstrs = append(*pstrs, "f2")
}

func f3(pnums *[][]int) {
	tmp1 := &[]int{1, 2}
	*pnums = append(*pnums, *tmp1)
	*tmp1 = append(*tmp1, 3)
	*pnums = append(*pnums, *tmp1)
}

func f4(pnums *[][]int) {
	tmp1 := &[]int{1, 2}
	*pnums = append(*pnums, *tmp1)
	*tmp1 = append(*tmp1, 3)
	*pnums = append(*pnums, *tmp1)
}
