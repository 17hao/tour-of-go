package main

import (
	"fmt"
	"os"

	"github.com/17hao/tour-of-go/antlr/expr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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

//func main() {
//	input, err := antlr.NewFileStream(os.Args[1])
//	if err != nil {
//		panic(err)
//	}
//	lexer := parser.NewCalcLexer(input)
//	tokens := antlr.NewCommonTokenStream(lexer, 0)
//	p := parser.NewCalcParser(tokens)
//	tree := p.Prog()
//	println(tree.GetText())
//	eval := NewMyCalcVisitor()
//	eval.Visit(tree)
//}

//func main() {
//	input := "one two three"
//	words := strings.Fields(input)
//	for _, word := range words {
//		println(word)
//	}
//}
