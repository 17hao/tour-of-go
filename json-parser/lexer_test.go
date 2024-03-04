package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	input := `{"str": "a neat string", "num": 3.14159, "obj": {"num": -1}, "arr": [1,2,3]}`
	l := NewLexer(input)
	for {
		token, err := l.Scan()
		if err != nil {
			panic(err)
		}
		if token.Type == EOF {
			break
		}
		fmt.Println(token.String())
	}
}
