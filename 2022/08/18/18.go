package main

import "time"

func main() {
	start, _ := time.Parse(`2006-01-02 15:04:05`, "2022-03-01 02:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2024-03-03 01:00:00")
	day, _ := time.ParseDuration("24h")
	println(start.Sub(end).Round(day).String())
}
