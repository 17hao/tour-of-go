package main

import "fmt"

func main() {
	f := float64(1)
	const i = 1
	if f == i {
		println("f == i")
	} else {
		println("f != i")
	}

	//fmt.Printf("f=%d", f)
	fmt.Printf("f=%f\n", f)
	fmt.Printf("int(f)=%d\n", int64(f))
}
