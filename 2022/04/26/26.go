package main

import (
	"fmt"

	"github.com/17hao/tour-of-go/antlr/calc/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type MyListener struct {
	*parser.BaseCalcListener
}

func NewMyListener() *MyListener {
	return new(MyListener)
}

func (l *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	lexer := parser.NewCalcLexer(antlr.NewInputStream("(1+3)*5"))
	p := parser.NewCalcParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := p.Prog()
	listener := NewMyListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}
