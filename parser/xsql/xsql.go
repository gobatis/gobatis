package xsql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gozelle/spew"
)

type XSQL struct {
	raw     strings.Builder
	dynamic bool
	vars    []any
	sql     string
	ws      bool
}

func (x *XSQL) Raw() string {
	return x.raw.String()
}

func (x *XSQL) SQL() string {
	return x.sql
}

func (x *XSQL) Dynamic() bool {
	return x.dynamic
}

func (x *XSQL) Vars() []any {
	return x.vars
}

func (x *XSQL) WriteWS() {
	if x.ws {
		return
	}
	x.ws = true
	x.raw.WriteString(" ")
}

func (x *XSQL) WriteString(v string) {
	x.ws = false
	x.raw.WriteString(v)
}

func Parse(source string, vars map[string]any) (*XSQL, error) {

	errs := &commons.ErrorListener{}

	//source = replaceIsolatedLessThanWithEntity(source)
	lexer := NewXSQLLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewXSQLParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.Content()
	if errs.GetError() != nil {
		return nil, errs.GetError()
	}

	v := &Visitor{
		errs: errs,
		xsql: &XSQL{},
	}
	_ = v.VisitContent(tree.(*ContentContext))

	return v.xsql, errs.GetError()
}

type Visitor struct {
	errs *commons.ErrorListener
	xsql *XSQL
}

func (v Visitor) VisitContent(ctx *ContentContext) interface{} {
	fmt.Println("content:", ctx.GetText())
	for _, c := range ctx.GetChildren() {
		if v.errs.GetError() != nil {
			return nil
		}
		switch t := c.(type) {
		case *StartContext:
			v.visitStart(t)
		case *EndContext:
			v.visitEnd(t)
		case *ContentContext:
			v.VisitContent(t)
		case *ExprContext:
			v.visitExpr(t)
		case *ReferenceContext:
			v.visitReference(t)
		case *ChardataContext:
			v.visitCharData(t)
		default:
			v.errs.AddError(fmt.Errorf("unsupport rule: %v", c.GetPayload()))
		}
	}

	return "a"
}

func (v Visitor) visitStart(ctx *StartContext) {
	v.xsql.WriteString(ctx.GetText())
}

func (v Visitor) visitEnd(ctx *EndContext) {
	v.xsql.WriteString(ctx.GetText())
}

func (v Visitor) visitExpr(ctx *ExprContext) {
	if v.errs.GetError() != nil {
		return
	}
	if ctx.HASH() != nil {
		v.xsql.raw.WriteString(fmt.Sprintf("##{%s}", ctx.GetVal().GetText()))
	} else {
		v.xsql.raw.WriteString(fmt.Sprintf("$${%s}", ctx.GetVal().GetText()))
	}
}

func (v Visitor) visitAttribute(ctx *AttributeContext) {
	if v.errs.GetError() != nil {
		return
	}
	spew.Json(ctx.GetText())
}

func (v Visitor) visitCharData(ctx *ChardataContext) {
	if v.errs.GetError() != nil {
		return
	}
	if ctx.WS() != nil {
		v.xsql.WriteWS()
	} else {
		v.xsql.WriteString(ctx.GetText())
	}
}

func (v Visitor) visitReference(ctx *ReferenceContext) {
	if ctx.EntityRef() != nil {
		c := ""
		switch ctx.EntityRef().GetText() {
		case "&lt;":
			c = "<"
		case "&gt":
			c = ">"
		case "&amp;":
			c = "&"
		case "&apos;":
			c = "'"
		case "&quot;":
			c = "\""
		default:
			v.errs.AddError(fmt.Errorf("unkonwn reference: %s", ctx.EntityRef().GetText()))
			return
		}
		v.xsql.WriteString(c)
	}
}
