package main

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
)

type Token struct {
	Type    Type
	Literal string
	Start   int
	End     int
}
