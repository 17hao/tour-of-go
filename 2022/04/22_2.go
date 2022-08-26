package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"shiqihao.xyz/tour-of-go/antlr/calc/parser"
)

type MyCalcVisitor struct {
	memory map[string]int
	*parser.BaseCalcVisitor
}

func NewMyCalcVisitor() *MyCalcVisitor {
	return &MyCalcVisitor{
		memory:          make(map[string]int),
		BaseCalcVisitor: &parser.BaseCalcVisitor{},
	}
}

func (visitor *MyCalcVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(visitor)
}

func (visitor *MyCalcVisitor) VisitCalc(ctx *parser.CalcContext) interface{} {
	return ctx.Stat(0).Accept(visitor)
}

// VisitAssign ID '=' expr NEWLINE
func (visitor *MyCalcVisitor) VisitAssign(ctx *parser.AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := visitor.Visit(ctx.Expr()).(int)
	visitor.memory[id] = value
	return value
}

// VisitPrintExpr expr NEWLINE
func (visitor *MyCalcVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	value := visitor.Visit(ctx.Expr()).(int)
	fmt.Println(value)
	return 0
}

// VisitInt INT
func (visitor *MyCalcVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	res, _ := strconv.Atoi(ctx.INT().GetText())
	return res
}

// VisitId ID
func (visitor *MyCalcVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	if _, ok := visitor.memory[id]; ok {
		return visitor.memory[id]
	}
	return 0
}

// VisitAddSub expr op=('+'|'-') expr
func (visitor *MyCalcVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := visitor.Visit(ctx.Expr(0)).(int)
	right := visitor.Visit(ctx.Expr(1)).(int)
	if ctx.GetOp().GetTokenType() == parser.CalcParserADD {
		return left + right
	} else {
		return left - right
	}
}

// VisitMulDiv expr op=('*'|'/') expr
func (visitor *MyCalcVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := visitor.Visit(ctx.Expr(0)).(int)  // get value of left subexpression
	right := visitor.Visit(ctx.Expr(1)).(int) // get right value of right subexpression
	if ctx.GetOp().GetTokenType() == parser.CalcParserMUL {
		return left * right
	} else {
		return left / right
	}
}

// VisitParens '(' expr ')'
func (visitor *MyCalcVisitor) VisitParens(ctx *parser.ParensContext) interface{} {
	return visitor.Visit(ctx.Expr()) // return child expr;s value
}

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	lexer := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcParser(tokens)
	tree := p.Prog()
	println(tree.GetText())
	eval := NewMyCalcVisitor()
	eval.Visit(tree)
}
