package main

import "fmt"

var Conf = FooConf

func loadConf() {
	FooConf.Key1 = "value1"
}

func main() {
	loadConf()
	fmt.Printf("Conf=%+v\n", Conf)
	fmt.Printf("FooConf=%+v\n", FooConf)
	fmt.Printf("%s\n", Conf.Key1)
}
