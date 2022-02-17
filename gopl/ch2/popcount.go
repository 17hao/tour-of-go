// PopCount return the number of set bits, that is, bits whose value
// is 1, in an uint64 value.
package main

import "fmt"

var pc [256]byte

func init() {
	//for i, _ := range pc {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(PopCount(1))
	fmt.Println(PopCount(2))
	fmt.Println(PopCount(4))
	fmt.Println(PopCount(8))
}
