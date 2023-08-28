// Code generated from SimpleMath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package simplemath // SimpleMath
import "github.com/antlr4-go/antlr/v4"

type BaseSimpleMathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleMathVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleMathVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleMathVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}
