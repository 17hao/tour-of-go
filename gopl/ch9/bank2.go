package main

import "fmt"

var (
	sema = make(chan struct{}, 1)
	balance int
)

func Deposit2(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func main() {
	balance = 100
	Deposit2(10)
	fmt.Println(Balance2())
}