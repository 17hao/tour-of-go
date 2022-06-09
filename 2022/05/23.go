package main

import "reflect"

func main() {
	a := []int{1, 2, 3}
	b := []string{"1", "2"}
	c := []float64{1.0, 2.0, 3.0}
	d := []bool{true, false}
	e := 1
	pe := &e
	typeSwitch(a)
	typeSwitch(b)
	typeSwitch(c)
	typeSwitch(d)
	typeSwitch(e)
	typeSwitch(pe)

	println(reflect.TypeOf(a).String())
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case []int:
		println("[]int")
	case []string:
		println("[]string")
	case []float64:
		println("[]float64")
	case []bool:
		println("[]bool")
	case int:
		println("int")
	case *int:
		println("*int")
	}
}
