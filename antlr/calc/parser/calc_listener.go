// Code generated from Calc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterCalc is called when entering the calc production.
	EnterCalc(c *CalcContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterBlank is called when entering the blank production.
	EnterBlank(c *BlankContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// ExitCalc is called when exiting the calc production.
	ExitCalc(c *CalcContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitBlank is called when exiting the blank production.
	ExitBlank(c *BlankContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)
}
