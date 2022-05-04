package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("main")
}
