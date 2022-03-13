package main

import "fmt"

func main() {
	join()
}

func join() {
	sliceA := []int{1, 2}
	sliceB := []int{2, 3}
	for _, i := range sliceB {
		contains := false
		for _, j := range sliceA {
			if i == j {
				contains = true
			}
		}
		if !contains {
			sliceA = append(sliceA, i)
		}
	}

	for _, i := range sliceA {
		fmt.Println(i)
	}
}

func mapTest() {
	employees := map[string][]int{"a": []int{3, 4}}
	for _, i := range []int{4, 5} {
		contains := false
		for _, j := range employees["a"] {
			if i == j {
				contains = true
			}
		}
		if !contains {
			employees["a"] = append(employees["a"], i)
		}
	}
	employees["b"] = []int{1, 2}
	for k, v := range employees {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("c", ": ", employees["c"])
}
