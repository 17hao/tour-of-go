// Code generated from Calc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCalc is called when production calc is entered.
func (s *BaseCalcListener) EnterCalc(ctx *CalcContext) {}

// ExitCalc is called when production calc is exited.
func (s *BaseCalcListener) ExitCalc(ctx *CalcContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseCalcListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseCalcListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseCalcListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseCalcListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseCalcListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseCalcListener) ExitBlank(ctx *BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseCalcListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseCalcListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseCalcListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseCalcListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *BaseCalcListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseCalcListener) ExitInt(ctx *IntContext) {}
