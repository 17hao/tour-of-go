package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	numbers = append(numbers[0:1], numbers[2:]...)
	fmt.Println(numbers)
}
