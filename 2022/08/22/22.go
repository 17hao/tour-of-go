package main

func main() {
	nums := []int{1, 2, 3, 4}

	println(`nums[:0]`)
	v1 := nums[:0]
	for _, n := range v1 {
		println(n)
	}

	println(`nums[1:]`)
	v2 := nums[1:]
	for _, n := range v2 {
		println(n)
	}
}
