package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello, 世界"
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("%q ", str[i])
	}

	m := map[interface{}]int{
		'*': 1,
	}
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	fmt.Println(reflect.TypeOf(f1()))
	fmt.Println(reflect.TypeOf(f2()))
	fmt.Println(reflect.TypeOf(f3()))
	fmt.Println(reflect.TypeOf(f4()))

	fmt.Println(m[f1()])
	fmt.Println(m[f2()])
	fmt.Println(m[f3()])
	fmt.Println(m[f4()])
}

func f1() interface{} {
	return "*"
}

func f2() interface{} {
	var str = "*"
	return str[0]
}

func f3() interface{} {
	return '*'
}

func f4() interface{} {
	str := "*"
	return rune(str[0])
}
