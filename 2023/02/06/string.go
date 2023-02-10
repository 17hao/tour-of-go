package main

import (
	"fmt"
	"strconv"
)

func FmtString() string {
	i := 1
	return fmt.Sprintf("this is a string %d", i)
}

func StrconvString() string {
	i := 1
	return "this is a string " + strconv.Itoa(i)
}
