package main

import (
	"fmt"
	"regexp"
)

func main() {
	boolean := regexp.MustCompile(`bool\([tf]\)`)
	fmt.Println(boolean.MatchString("bool(t)"))

	num := regexp.MustCompile(`number\(\d+\.?\d+\)`)
	fmt.Println(num.MatchString("number(12345.6)"))
	fmt.Println(num.MatchString("number(12345)"))

	str := regexp.MustCompile(`string\(.+\)`)
	fmt.Println(str.MatchString("string(12345)"))
	fmt.Println(str.MatchString("string(12345.6)"))

	enum := regexp.MustCompile(`enum\((\w{1,64},?)+\)`)
	fmt.Println(enum.MatchString("enum(e1,e2,e3,e4,e5,e6)"))

	date := regexp.MustCompile(`date\(\d{4}-\d{2}-\d{2}\)`)
	fmt.Println(date.MatchString("date(2021-11-24)"))

	datetime := regexp.MustCompile(`datetime\(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\)`)
	fmt.Println(datetime.MatchString("datetime(2021-11-24 16:00:00)"))

	strings := regexp.MustCompile(`strings\((.{1,64},?)+\)`)
	fmt.Println(strings.MatchString("strings(str-1,str-2,str-3)"))

	empty := regexp.MustCompile(`empty\(\)`)
	fmt.Println(empty.MatchString("empty()"))
}
