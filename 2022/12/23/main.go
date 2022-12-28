package main

import (
	"encoding/json"
	"fmt"
)

type enum string

var (
	text enum = "enum"
)

type s struct {
	Enum     enum        `json:"type"`
	Data     interface{} `json:"data"`
	Property interface{} `json:"property,omitempty"`
}

func main() {
	m := map[string]string{"data": "v1", "type": "text"}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
	var i s
	_ = json.Unmarshal(b, &i)
	fmt.Printf("%+v", i)
}
