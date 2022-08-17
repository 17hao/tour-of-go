package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"shiqihao.xyz/tour-of-go/antlr/expr/parser"
)

type TreeShapeListener struct {
	*parser.BaseExprListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	lexer := parser.NewExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExprParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	fmt.Println(tree.ToStringTree([]string{}, p))
}
