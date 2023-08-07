package main

import (
	"errors"
	"fmt"
)

func f1() {
	err := errors.New("mock")
	defer func(error) {
		if err != nil {
			fmt.Printf("defer: %v\n", err)
		}
	}(err)
}

func f2() {
	err := errors.New("mock")
	defer func(err error) {
		if err != nil {
			fmt.Printf("defer: %v\n", err)
		}
	}(err)
}

func f3() {
	x := 10
	defer func(i int) {
		fmt.Printf("defer: i=%v\n", i)
		fmt.Printf("defer: x=%v\n", x)
	}(x)
}

func f4() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return 9
}

func f5() {
	x := 9
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return
}

func main() {
	// const i = iota
	//
	// fmt.Printf("%+v", i == 0)
	f1()
	f2()
	f3()
	fmt.Printf("x=%v\n", f4())
	f5()
}
