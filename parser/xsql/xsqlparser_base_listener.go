// Code generated from XSQLParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package xsql // XSQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseXSQLParserListener is a complete listener for a parse tree produced by XSQLParser.
type BaseXSQLParserListener struct{}

var _ XSQLParserListener = &BaseXSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterContent is called when production content is entered.
func (s *BaseXSQLParserListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseXSQLParserListener) ExitContent(ctx *ContentContext) {}

// EnterElement is called when production element is entered.
func (s *BaseXSQLParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseXSQLParserListener) ExitElement(ctx *ElementContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseXSQLParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseXSQLParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseXSQLParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseXSQLParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterChardata is called when production chardata is entered.
func (s *BaseXSQLParserListener) EnterChardata(ctx *ChardataContext) {}

// ExitChardata is called when production chardata is exited.
func (s *BaseXSQLParserListener) ExitChardata(ctx *ChardataContext) {}
