package main

import (
	"encoding/json"
	"fmt"
)

type innerModel struct {
	Field1 string
}

type model struct {
	ID int64
	M1 *innerModel
}

func main() {
	m := &model{
		ID: int64(1),
		M1: &innerModel{
			Field1: "f1",
		},
	}
	bytes, _ := json.MarshalIndent(m, "", "\t")
	fmt.Println(string(bytes))

	fmt.Printf("%+v", m)
}
