// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package gpath // JsonPath
import "github.com/antlr4-go/antlr/v4"

type BaseJsonPathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonPathVisitor) VisitJsonpath(ctx *JsonpathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonPathVisitor) VisitDotnotation(ctx *DotnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonPathVisitor) VisitDotnotation_expr(ctx *Dotnotation_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonPathVisitor) VisitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}
