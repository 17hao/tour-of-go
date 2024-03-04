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
	return fmt.Sprintf("Type=%8s\t\tLiteral=%15s\t\tStart=%3d\t\tEnd=%3d", t.Type, t.Literal, t.Start, t.End)
}
