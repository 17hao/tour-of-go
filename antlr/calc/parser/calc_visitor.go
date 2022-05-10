// Code generated from Calc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#calc.
	VisitCalc(ctx *CalcContext) interface{}

	// Visit a parse tree produced by CalcParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by CalcParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by CalcParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by CalcParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by CalcParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalcParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalcParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by CalcParser#int.
	VisitInt(ctx *IntContext) interface{}
}
