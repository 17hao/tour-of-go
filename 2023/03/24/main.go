package main

import (
	"reflect"
)

func main() {
	s1 := make([]string, 0)
	s2 := []string{}
	var s3 []string

	rv1 := reflect.ValueOf(s1)
	println(rv1.Len())

	rv2 := reflect.ValueOf(s2)
	println(rv2.Len())

	rv3 := reflect.ValueOf(s3)
	println(rv3.Len())

	println(reflect.ValueOf(nil).Len())
}
