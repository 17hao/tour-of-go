package main

type MyInterface interface {
	DoSomething() string
}

func GetMyInterface() MyInterface {
	return &MyInterfaceImpl{}
}

type MyInterfaceImpl struct{}

func (i *MyInterfaceImpl) DoSomething() string {
	return "impl"
}
