// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package gpath // JsonPath
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by JsonPathParser.
type JsonPathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonPathParser#jsonpath.
	VisitJsonpath(ctx *JsonpathContext) interface{}

	// Visit a parse tree produced by JsonPathParser#dotnotation.
	VisitDotnotation(ctx *DotnotationContext) interface{}

	// Visit a parse tree produced by JsonPathParser#dotnotation_expr.
	VisitDotnotation_expr(ctx *Dotnotation_exprContext) interface{}

	// Visit a parse tree produced by JsonPathParser#identifierWithQualifier.
	VisitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) interface{}
}
