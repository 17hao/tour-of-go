package main

import "github.com/nyaruka/phonenumbers"

func main() {
	phoneNum, _ := phonenumbers.Parse("15868170425", "CN")
	res := phonenumbers.IsValidNumber(phoneNum)
	println(res)
}
