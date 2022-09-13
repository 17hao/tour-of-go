package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	// go 解析时间格式 年/月/日/时/分/秒 有固定的数值，参见 go 源码
	start, _ := time.Parse(`2006-01-02`, "2019-07-27")
	end, _ := time.Parse("2006-01-02 15:04:05", "2022-08-26 00:00:00")
	println(start.YearDay())
	println(end.YearDay())

	i := 100
	n := strconv.FormatInt(int64(i), 16)
	println(n)

	_, err := time.Parse("2006-01-02", "2022-02-29")
	if err != nil {
		log.Fatal(err)
	}
}
