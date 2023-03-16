package main

import (
	"fmt"
)

func main() {
	str1 := "hello, 世界"
	fmt.Printf("%q\n", str1[7])
	fmt.Printf("%+q\n", str1[7])
	fmt.Printf("%q\n", []rune(str1)[7])
	fmt.Printf("%+q\n", []rune(str1)[7])
	fmt.Printf("%s\n", string([]rune(str1)[7]))
	fmt.Printf("% x\n", str1)

	fmt.Println()

	str2 := "\x68\x65\x6c\x6c\x6f\x2c\x20\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Printf("%s\n", str2)  // hello, 世界
	fmt.Printf("%q\n", str2)  // "hello, 世界"
	fmt.Printf("%+q\n", str2) // "hello, \u4e16\u754c"
}
