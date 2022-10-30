package main

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type struct15 struct {
	Id      int
	Name    string
	Decimal decimal.Decimal
	HashMap map[string]string
	Array   []float64
}

func main() {
	s := struct15{
		Id:      1,
		Name:    "name",
		Decimal: decimal.NewFromFloat(3.14),
		HashMap: map[string]string{"k1": "v1", "k2": "v2"},
		Array:   []float64{2.14},
	}

	ps := &s

	fmt.Printf("s: %+v\n", s)
	fmt.Printf("ps: %+v\n", ps)

	b, _ := json.MarshalIndent(s, "", "\t")
	fmt.Printf("indented json s: %s\n", string(b))
}
