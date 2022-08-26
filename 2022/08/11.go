package main

import "github.com/nyaruka/phonenumbers"

func main() {
	phoneNum, _ := phonenumbers.Parse("+639694878122", "PH")
	res := phonenumbers.IsValidNumberForRegion(phoneNum, "PH")
	println(res)
}
