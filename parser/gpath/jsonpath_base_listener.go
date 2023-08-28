// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package gpath // JsonPath
import "github.com/antlr4-go/antlr/v4"

// BaseJsonPathListener is a complete listener for a parse tree produced by JsonPathParser.
type BaseJsonPathListener struct{}

var _ JsonPathListener = &BaseJsonPathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonPathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonPathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonPathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonPathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJsonpath is called when production jsonpath is entered.
func (s *BaseJsonPathListener) EnterJsonpath(ctx *JsonpathContext) {}

// ExitJsonpath is called when production jsonpath is exited.
func (s *BaseJsonPathListener) ExitJsonpath(ctx *JsonpathContext) {}

// EnterDotnotation is called when production dotnotation is entered.
func (s *BaseJsonPathListener) EnterDotnotation(ctx *DotnotationContext) {}

// ExitDotnotation is called when production dotnotation is exited.
func (s *BaseJsonPathListener) ExitDotnotation(ctx *DotnotationContext) {}

// EnterDotnotation_expr is called when production dotnotation_expr is entered.
func (s *BaseJsonPathListener) EnterDotnotation_expr(ctx *Dotnotation_exprContext) {}

// ExitDotnotation_expr is called when production dotnotation_expr is exited.
func (s *BaseJsonPathListener) ExitDotnotation_expr(ctx *Dotnotation_exprContext) {}

// EnterIdentifierWithQualifier is called when production identifierWithQualifier is entered.
func (s *BaseJsonPathListener) EnterIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// ExitIdentifierWithQualifier is called when production identifierWithQualifier is exited.
func (s *BaseJsonPathListener) ExitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}
