package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	m := map[string]string{
		"employee":   "a",
		"joinTime":   strconv.FormatInt(time.Now().Unix(), 10),
		"isResigned": "true",
	}
	s := make([]string, 0)
	for k, v := range m {
		s = append(s, k+v)
	}
	fmt.Printf("%+v", s)
}
