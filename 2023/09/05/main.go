package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	s := "-2231605685775754701"
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("%d\n", i)

	year, month, day := time.Now().Date()
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		logrus.Fatal(err)
	}
	t := time.Date(year, month, day, 10, 0, 0, 0, location)
	fmt.Printf("%v\n", t)

	d1, err := time.Parse(time.RFC3339, "2022-08-26T10:40:00+08:00")
	d2, err := time.Parse("2006-01-02 15:04:05.999999999", "2022-08-26 10:39:59.950345")
	fmt.Printf("d1=%v\nd2=%v\n", d1, d2)
	// fmt.Printf("d1-d3=%v\n", d1.Unix()-d3.Unix())

	t1 := time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), 0, 0, time.Now().Location())
	t2 := time.Date(d2.Year(), d2.Month(), d2.Day(), d2.Hour(), d2.Minute(), 0, 0, time.Now().Location())
	fmt.Printf("t1=%v\nt2=%v\n", t1, t2)
	fmt.Printf("t1-t2=%v\n", t1.Unix()-t2.Unix())
}
