package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func main() {
	id := 1
	res := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(id), 10)))
	fmt.Printf("%s\n", res)

	f := float64(41) / float64(42)
	fmt.Printf("%f\n", f)

	f1 := 1.236
	f2 := round(f1 * 100)
	fmt.Printf("%d\n", f2)
}
