package main

type myInterface interface {
	doSomething()
}

type myStruct struct {
	i myInterface
}

type myImplementation struct{}

func (s myImplementation) doSomething() {
	println("hello, world")
}

func main() {
	s := myStruct{}
	if s.i != nil {
		s.i.doSomething()
	} else {
		println("s.i is nil")
	}
}
