package main

import "fmt"

func g(i int) {
	if i > 3 {
		fmt.Printf("Panicking\n")
		panic(fmt.Sprintf("panic: i=%d", i))
	}
	fmt.Printf("g() i=%d\n", i)
	defer fmt.Printf("defer i=%d\n", i)
	g(i + 1)
}

func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v", err)
		}
	}()
	fmt.Printf("f()\n")
	g(0)
}

func main() {
	f()
}
