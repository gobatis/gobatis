// Code generated from SimpleMath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package simplemath // SimpleMath
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SimpleMathParser.
type SimpleMathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleMathParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SimpleMathParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SimpleMathParser#factor.
	VisitFactor(ctx *FactorContext) interface{}
}
