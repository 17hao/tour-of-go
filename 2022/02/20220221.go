package main

import (
	"fmt"
	"strings"
)

type MyInterface interface {
	toUpperCase() string

	toLowerCase() string

	original() string
}

type UpperStruct struct {
	s string
}

type LowerStruct struct {
	s string
}

type OriginStruct struct {
	s string
}

func (u UpperStruct) toUpperCase() string {
	return strings.ToUpper(u.s)
}

func (u UpperStruct) original() string {
	return u.s
}

func (l LowerStruct) toLowerCase() string {
	return strings.ToLower(l.s)
}

//func (l LowerStruct) original() string {
//	return l.s
//}

type MyStruct struct {
	*UpperStruct
	*LowerStruct
	*OriginStruct
}

func main() {
	//fmt.Println("hello, world!")
	str := "Hello, World!"
	var s = MyStruct{
		UpperStruct: &UpperStruct{str},
		LowerStruct: &LowerStruct{str},
	}
	fmt.Println(s.toUpperCase())
	fmt.Println(s.toLowerCase())
}
