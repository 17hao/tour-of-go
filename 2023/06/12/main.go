package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// 7108291135797362689,7241100973208027138,7241103841754611715,7241105661805756420,7241108703970181124
func f1() {
	data := "hello, world!"
	h := sha1.New()
	h.Write([]byte(data))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
	fmt.Printf("%d\n", len(hex.EncodeToString(h.Sum(nil))))
}

func main() {
	// f1()
	minID := 7108291135797362689
	maxID := 7241108703970181124
	count := 10
	min := minID
	gap := (maxID - minID) / count
	fmt.Printf("gap=%d\n", gap)

	for i := 1; i <= count; i++ {
		var max = 0
		if i == count {
			max = maxID
		} else {
			max = gap*i + minID
		}
		fmt.Printf("%d %d\n", min, max)
		min = max + 1
	}
}
