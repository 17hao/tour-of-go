package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("ENV"))
}

func init() {
	_ = os.Setenv("ENV", "staging")
}