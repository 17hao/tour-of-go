package main

import (
	"fmt"
	"net"
)

const IntConstant = 1

func main() {
	address, _ := net.InterfaceAddrs()
	for _, add := range address {
		fmt.Println(add)
	}

	fmt.Println(*getPtr())
}

func getPtr() *int {
	ic := IntConstant
	return &ic
}
