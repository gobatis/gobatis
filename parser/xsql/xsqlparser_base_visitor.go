// Code generated from XSQLParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package xsql // XSQLParser
import "github.com/antlr4-go/antlr/v4"

type BaseXSQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseXSQLParserVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXSQLParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXSQLParserVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXSQLParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXSQLParserVisitor) VisitChardata(ctx *ChardataContext) interface{} {
	return v.VisitChildren(ctx)
}
