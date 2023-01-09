package main

import "runtime"

func main() {
	println("NumCPU()=", runtime.NumCPU())
}
