package main

import (
	"fmt"
	"time"
)

func main() {
	maps := make(map[int]map[string]time.Time, 0)
	maps[0]["k"] = time.Now()
	fmt.Printf("%+v", maps)
}
