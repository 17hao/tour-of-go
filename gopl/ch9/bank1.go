package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit1(amount int) {
	deposits <- amount
}

func Balance1() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select{
		case amount := <- deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit1(100)
	fmt.Println(Balance1())
}
