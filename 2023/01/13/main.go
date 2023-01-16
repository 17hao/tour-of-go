package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	d := "2019-10-22T07:54:58.000000Z"

	t, err := time.Parse("2006-01-02T15:04:05.999999Z", d)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%+v", t)
}
