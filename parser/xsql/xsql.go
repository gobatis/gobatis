package xsql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gobatis/gobatis/parser/expr"
	"github.com/gozelle/spew"
)

type XSQL struct {
	sql     strings.Builder
	dynamic bool
	vars    []any
	ws      bool
	count   int
}

func (x *XSQL) SQL() string {
	return x.sql.String()
}

func (x *XSQL) Dynamic() bool {
	return x.dynamic
}

func (x *XSQL) Vars() []any {
	return x.vars
}

func (x *XSQL) Count() int {
	return x.count
}

func (x *XSQL) WriteWS() {
	if x.ws {
		return
	}
	x.ws = true
	x.sql.WriteString(" ")
}

func (x *XSQL) WriteString(v string) {
	x.ws = false
	x.sql.WriteString(v)
}

func (x *XSQL) AddVar(v any) {
	x.vars = append(x.vars, v)
	x.count++
}

type Tag struct {
	start *StartContext
	xsql  *XSQL
}

func Parse(source string, vars map[string]any) (*XSQL, error) {
	return parse(source, nil, vars)
}

type Formatter func(rv reflect.Value, escaper string) (s string, err error)

func Explain(formatter Formatter, source string, vars map[string]any) (string, error) {
	r, err := parse(source, formatter, vars)
	if err != nil {
		return "", err
	}
	return r.SQL(), nil
}

func parse(source string, formatter Formatter, vars map[string]any) (*XSQL, error) {

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
	if errs.Error() != nil {
		return nil, errs.Error()
	}

	v := &Visitor{
		ErrorListener: errs,
		vars:          vars,
		xsql:          &XSQL{},
		formatter:     formatter,
	}
	_ = v.VisitContent(tree.(*ContentContext))

	return v.xsql, errs.Error()
}

type Visitor struct {
	*commons.ErrorListener
	xsql      *XSQL
	vars      map[string]any
	formatter Formatter
	stack     *commons.Stack[Tag]
}

func (v Visitor) VisitContent(ctx *ContentContext) interface{} {
	//fmt.Println("content:", ctx.GetText())
	for _, c := range ctx.GetChildren() {
		if v.Error() != nil {
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
			v.AddError(fmt.Errorf("unsupport rule: %v", c.GetPayload()))
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
	if v.Error() != nil {
		return
	}
	rv, err := expr.Parse(ctx.GetVal().GetText(), v.vars)
	if err != nil {
		v.AddError(fmt.Errorf("parse expression: %s error: %w", ctx.GetVal().GetText(), err))
		return
	}
	if ctx.HASH() != nil && v.formatter == nil {
		v.bindExpr(rv)
	} else {
		v.explainExpr(rv)
	}
}

func (v Visitor) bindExpr(rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			v.xsql.AddVar(rv.Index(i).Interface())
			s = append(s, fmt.Sprintf("$%d", v.xsql.Count()))
		}
		v.xsql.WriteString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		v.xsql.AddVar(rv.Interface())
		v.xsql.WriteString(fmt.Sprintf("$%d", v.xsql.Count()))
	}
}

func (v Visitor) explainExpr(rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			s = append(s, fmt.Sprintf("%s", v.explainVar(rv.Index(i))))
		}
		v.xsql.WriteString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		v.xsql.WriteString(fmt.Sprintf("%s", v.explainVar(rv)))
	}
}

func (v Visitor) explainVar(rv reflect.Value) (r string) {
	r, err := v.formatter(rv, "'")
	if err != nil {
		v.AddError(err)
		return
	}
	return
}

func (v Visitor) visitAttribute(ctx *AttributeContext) {
	if v.Error() != nil {
		return
	}
	spew.Json(ctx.GetText())
}

func (v Visitor) visitCharData(ctx *ChardataContext) {
	if v.Error() != nil {
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
			v.AddError(fmt.Errorf("unkonwn reference: %s", ctx.EntityRef().GetText()))
			return
		}
		v.xsql.WriteString(c)
	}
}
