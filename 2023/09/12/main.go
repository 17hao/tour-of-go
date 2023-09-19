package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func encrypt(id int64) string {
	const s = "const"
	str := strconv.FormatInt(id, 10)
	return base64.RawURLEncoding.EncodeToString([]byte(str + s))
}

func main() {
	fmt.Printf("%s\n", encrypt(100))

	nums := []int64{0, 1, 2, 3}
	pageSize := 3
	fmt.Printf("%+v\n", nums[:pageSize])

	// ticker := time.NewTicker(1 * time.Second)
	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case <-ticker.C:
	// 			fmt.Printf("%s\n", time.Now().String())
	// 		}
	// 	}
	// }()
	// time.Sleep(3 * time.Second)
	// ticker.Stop()
	// done <- true
}
