package main

import (
	"fmt"
	"time"
)

func main() {
	timestampNano := time.Now().UnixNano()
	fmt.Printf("nano second: %d\n", timestampNano)

	timestampMilli := timestampNano / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Printf("mili second: %d\n", timestampMilli)
}
