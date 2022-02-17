// greatest common divisor
package main

import "fmt"

func main() {
	fmt.Print(gcd(9, 15))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
