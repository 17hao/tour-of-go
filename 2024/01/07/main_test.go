package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_combinations(t *testing.T) {
	decimals := []decimal.Decimal{
		decimal.NewFromInt(1),
		decimal.NewFromInt(2),
		decimal.NewFromInt(3),
		decimal.NewFromInt(4),
	}
	combinations(decimals, 3, []decimal.Decimal{}, 0)
	for _, list := range combinationLists {
		fmt.Println(list)
	}
	//assert.Equal(t, 4, len(combinationLists))
	//assert.Equal(t, "1", combinationLists[3][2].String())
}

func Test_f1(t *testing.T) {
	//for i := 0; i < b.N; i++ {
	builder := strings.Builder{}
	for j := 0; j < 3000; j++ {
		f := math.Trunc(rand.Float64()*100) / 100
		builder.WriteString(decimal.NewFromFloat(f).String())
		builder.WriteString(" ")
	}
	fmt.Println(builder.String())

	//f1(decimals, 20, decimal.NewFromFloat(100))
	//}
}
