package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := MessageA{
		Type: "type-a",
		A:    "a",
	}

	b := MessageB{
		Type: "type-b",
		B:    []Message{a},
	}

	bytes, _ := json.Marshal(b)
	fmt.Printf("%s\n", string(bytes))

	var newB MessageB
	_ = json.Unmarshal(bytes, &newB)
	fmt.Printf("%+v\n", newB)
}

type Message interface {
	GetType() string
}

type MessageA struct {
	Type string

	A string
}

func (m MessageA) GetType() string {
	return m.Type
}

type MessageB struct {
	Type string

	B []Message
}

func (m MessageB) GetType() string {
	return m.Type
}
