// Code generated from XSQLParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package xsql // XSQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by XSQLParser.
type XSQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by XSQLParser#content.
	VisitContent(ctx *ContentContext) interface{}

	// Visit a parse tree produced by XSQLParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by XSQLParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by XSQLParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by XSQLParser#chardata.
	VisitChardata(ctx *ChardataContext) interface{}
}
