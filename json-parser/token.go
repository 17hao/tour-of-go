package main

import (
	"fmt"
)

type Type string

const (
	LeftBrace    Type = "{"
	RightBrace   Type = "{"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"

	String Type = "string"
	Number Type = "number"

	True  Type = "true"
	False Type = "false"
	Null  Type = "null"

	EOF     Type = "eof"
	Invalid Type = "invalid"
)

type Token struct {
	Type    Type
	Literal string
	Start   int
	End     int
}

func (t Token) String() string {
	return fmt.Sprintf("Start=%-3d\t\tEnd=%-3d\t\tType=%-8s\t\tLiteral=%s\t\t", t.Start, t.End, t.Type, t.Literal)
}
