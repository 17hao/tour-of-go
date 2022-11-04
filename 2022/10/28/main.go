package main

type MyInterface interface {
	DoSomething()
}

type MyImplementation struct {
}

func (s MyImplementation) DoSomething() {
	println("xxxx")
}

type MyStruct struct {
	i MyInterface
}

func Foo(i MyInterface) {
	i.DoSomething()
}

func main() {
	s := MyStruct{
		i: MyImplementation{},
	}
	if s.i != nil {
		Foo(&MyImplementation{})
	} else {
		print("MyStruct.i is nil")
	}
}
