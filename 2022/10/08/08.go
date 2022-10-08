package main

import (
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	println(math.MaxFloat64)
	println(9.999999999999831)
	println(9.99999)

	num1, _ := decimal.NewFromString("9.999999999999871")
	println(num1.String())

	num2 := num1.Round(13)
	println(num2.String())

	num3 := num2.Round(5)
	println(num3.String())

	x := float64(1 / 3)
	var y float64 = 1 / 3
	z := float64(1) - x - y
	println(z)

	xx := decimal.NewFromFloat(1 / 3)
	yy := decimal.NewFromFloat(1 / 3)
	zz := decimal.NewFromInt(1)
	println(zz.Sub(xx).Sub(yy).String())

	d := math.MaxInt64
	println(float64(d))

	println(float64(9007199254740992) == float64(9007199254740991)) // false
	println(float64(9007199254740993) == float64(9007199254740992)) // true

	println(math.Pow(2, 53) == 9007199254740992)

	println(precision(3.14))             // 2
	println(precision(3.141592653589))   // 12
	println(precision(3.1415926535897))  // x
	println(precision(3.14159265358979)) // x

	println("===")

	println(digits(3.14))                // 3
	println(digits(3.141592653589))      // 13
	println(digits(3.1415926535897))     // x
	println(digits(3.14159265358979))    // x
	println(digits(13.14))               // 4
	println(digits(12233.141592653589))  // 17
	println(digits(112233.141592653589)) // x
}

func precision(f float64) int {
	res := 0
	i := float64(1)
	for math.Round(f*i)/i != f {
		f *= 10
		res++
	}
	return res
}

func digits(f float64) int {
	res := 0
	i := float64(1)
	for math.Round(f*i)/i != f {
		f *= 10
	}
	for f > 1 {
		f /= 10
		res++
	}
	return res
}
