package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"unsafe"
)

func f1(s1 []int, s2 *[][]int, idx int) {
	if len(*s2) == 2 {
		return
	}
	if len(s1) == 9 {
		*s2 = append(*s2, s1)
		return
	}
	for i := idx; i < 9; i++ {
		s1 = append(s1, i)
		f1(s1, s2, i+1)
	}
}

func f2(s1 []int, s2 [][]int, idx int) {
	if len(s2) == 2 {
		return
	}
	if len(s1) == 9 {
		//s := s1[0:]
		s2 = append(s2, s1)
		return
	}
	for i := idx; i < 9; i++ {
		s1 = append(s1, i)
		f2(s1, s2, i+1)
	}
}

func f3(s []int, ss *[][]int, count int) {
	if count == 0 {
		(*ss)[2][2] = -100
		return
	}
	s = append(s, rand.Intn(10))
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("len=%d cap=%d hdr=%+v %+v\n", len(s), cap(s), hdr, ss)
	*ss = append(*ss, s)
	f3(s, ss, count-1)
}

func contains(s []int, n int) bool {
	for _, each := range s {
		if each == n {
			return true
		}
	}
	return false
}

func dfs(nums []int, s *[]int, res *[][]int) {
	if len(*s) == len(nums) {
		news := make([]int, 0)
		for _, n := range *s {
			news = append(news, n)
		}
		//news := (*s)[0:]
		*res = append(*res, news)
		return
	}
	for _, n := range nums {
		if contains(*s, n) {
			continue
		}
		*s = append(*s, n)
		dfs(nums, s, res)
		*s = (*s)[0 : len(*s)-1]
	}
}

func dfs2(nums []int, s []int, res *[][]int) {
	if len(s) == len(nums) {
		*res = append(*res, s)
		return
	}
	for _, n := range nums {
		if contains(s, n) {
			continue
		}
		//s = append(s, n)
		//dfs2(nums, s, res)
		//s = s[0 : len(s)-1]
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	dfs2(nums, []int{}, &res)
	return res
}

func main() {
	//s1 := []int{}
	//ps2 := &[][]int{}
	//f1(s1, ps2, 0)
	//for _, s := range *ps2 {
	//	fmt.Printf("%v\n", s)
	//}

	//s1 := []int{}
	//s2 := make([][]int, 2)
	////s2[0] = make([]int, 9)
	////s2[1] = make([]int, 9)
	//f2(s1, s2, 0)
	//for _, s := range s2 {
	//	fmt.Printf("%v\n", s)
	//}

	ss := [][]int{}
	f3([]int{}, &ss, 5)
	for _, s := range ss {
		fmt.Printf("%v\n", s)
	}

	//res := permute([]int{1, 2, 3, 4})
	//for _, p := range res {
	//	fmt.Printf("%+v\n", p)
	//}
}
