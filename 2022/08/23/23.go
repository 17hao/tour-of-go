package main

import "strconv"

// mnd: magic number detect
func main() {
	i := int64(100)
	n := strconv.FormatInt(i, 10)
	println(n)
}
