package main

import "time"

func main() {
	n := time.Now()
	println(n.Unix())
	println(n.UnixNano())
	println(n.Unix() * 1e3)
	println(n.UnixNano() / 1e6)

	m := time.Date(2024, 2, 10, 0, 0, 0, 0, time.Now().Location())
	println(lastOfMonth(m).String())
}

func lastOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 1, -1)
}
