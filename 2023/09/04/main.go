package main

import (
	"fmt"
	"time"
)

func main() {
	format := "20060102150405"
	t := time.Now().Format(format)
	fmt.Printf("%s", t)
}
