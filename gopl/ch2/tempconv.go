package main

import (
	"fmt"
	"shiqihao.xyz/tour-of-go/gopl/ch2/tempconv"
)

func main() {
	fmt.Printf("%g°C\n", tempconv.BoilingC-tempconv.FreezingC)
	fmt.Println(tempconv.FToC(212.0))
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
