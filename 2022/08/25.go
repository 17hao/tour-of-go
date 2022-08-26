package main

import (
	"strconv"
	"time"
)

func main() {
	start, _ := time.Parse(`2006-01-02`, "2019-07-27")
	end, _ := time.Parse("2006-01-02 15:04:05", "2022-08-26 00:00:00")
	println(start.YearDay())
	println(end.YearDay())

	i := 100
	n := strconv.FormatInt(int64(i), 16)
	println(n)
}
