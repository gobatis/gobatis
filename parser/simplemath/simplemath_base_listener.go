// Code generated from SimpleMath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package simplemath // SimpleMath
import "github.com/antlr4-go/antlr/v4"

// BaseSimpleMathListener is a complete listener for a parse tree produced by SimpleMathParser.
type BaseSimpleMathListener struct{}

var _ SimpleMathListener = &BaseSimpleMathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleMathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleMathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleMathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleMathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSimpleMathListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSimpleMathListener) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSimpleMathListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSimpleMathListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseSimpleMathListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseSimpleMathListener) ExitFactor(ctx *FactorContext) {}
