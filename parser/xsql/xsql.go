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

type tag struct {
	*Fragment
	ctx *StartContext
}

type Fragment struct {
	statement strings.Builder
	dynamic   bool
	vars      []any
	ws        bool
}

func (x *Fragment) Statement() string {
	return x.statement.String()
}

func (x *Fragment) Dynamic() bool {
	return x.dynamic
}

func (x *Fragment) Vars() []any {
	return x.vars
}

func (x *Fragment) writeWS() {
	if x.ws {
		return
	}
	x.ws = true
	x.statement.WriteString(" ")
}

func (x *Fragment) writeString(v string) {
	x.ws = false
	x.statement.WriteString(v)
}

func (x *Fragment) addVar(v ...any) {
	x.vars = append(x.vars, v...)
}

func Parse(source string, vars map[string]any) (*Fragment, error) {
	return parse(source, nil, vars)
}

type Formatter func(rv reflect.Value, escaper string) (s string, err error)

func Explain(formatter Formatter, source string, vars map[string]any) (string, error) {
	r, err := parse(source, formatter, vars)
	if err != nil {
		return "", err
	}
	return r.Statement(), nil
}

func parse(source string, formatter Formatter, vars map[string]any) (*Fragment, error) {

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

	v := &visitor{
		ErrorListener: errs,
		vars:          vars,
		formatter:     formatter,
		stack:         commons.NewStack[*tag](),
	}
	v.stack.Push(newTag(nil))
	v.VisitContent(tree.(*ContentContext))

	if v.stack.Len() != 1 {
		return nil, fmt.Errorf("expact 1 tag in stack, got: %d", v.stack.Len())
	}
	return v.stack.Peek().Fragment, errs.Error()
}

type visitor struct {
	*commons.ErrorListener
	count     int
	formatter Formatter
	stack     *commons.Stack[*tag]
	choose    []*tag
	vars      map[string]any
}

func (v *visitor) writeWS() {
	v.stack.Peek().writeWS()
}

func (v *visitor) writeString(vv string) {
	v.stack.Peek().writeString(vv)
}

func (v *visitor) addVar(vv ...any) {
	v.stack.Peek().addVar(vv...)
	v.count += len(vv)
}

func (v *visitor) VisitContent(ctx *ContentContext) {

	for _, c := range ctx.GetChildren() {
		if v.Error() != nil {
			return
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
}

func newTag(ctx *StartContext) *tag {
	return &tag{
		ctx:      ctx,
		Fragment: &Fragment{},
	}
}

func (v *visitor) visitStart(ctx *StartContext) {
	if v.Error() != nil {
		return
	}
	v.stack.Push(newTag(ctx))
}

func (v *visitor) visitIf() {
	t := v.stack.Peek()
	if l := len(t.ctx.AllAttribute()); l != 1 {
		v.AddError(fmt.Errorf("<if> expect 1 attribte got: %d", l))
		return
	}
	if t.ctx.Attribute(0).NAME().GetText() != "test" {
		v.AddError(fmt.Errorf("<if> only accept test attribte got: %s", t.ctx.Attribute(0).NAME().GetText()))
		return
	}

	v.stack.Pop()
	if v.test(t.ctx.Attribute(0).STRING().GetText()) {
		fmt.Println("if t", t.Statement())
		v.merge(t)
	}
}

func (v *visitor) test(s string) bool {
	r, err := expr.Parse(s[1:len(s)-1], v.vars)
	if err != nil {
		v.AddError(fmt.Errorf("parse test expression: %s error: %w", s, err))
		return false
	}

	if r.Kind() != reflect.Bool {
		v.AddError(fmt.Errorf("if test expression result: %s expect bool, got: %s", s, r.Kind()))
		return false
	}
	return r.Bool()
}

func (v *visitor) visitChoose() {

	var otherwise *tag
	var whens []*tag
	for i := 0; i < len(v.choose); i++ {
		if v.choose[i].ctx.NAME().GetText() == "otherwise" {
			if otherwise != nil {
				v.AddError(fmt.Errorf("in the choose block, otherwise is allowed to appear only once"))
				return
			}
			otherwise = v.choose[i]
		} else {
			whens = append(whens, v.choose[i])
		}
	}

	v.stack.Pop()

	for _, vv := range whens {
		if v.visitWhen(vv) || v.Error() != nil {
			return
		}
	}

	if otherwise != nil {
		v.merge(otherwise)
	}
}

func (v *visitor) merge(t *tag) {
	v.writeString(t.Statement())
	v.addVar(t.vars...)
}

func (v *visitor) visitWhen(t *tag) bool {
	if l := len(t.ctx.AllAttribute()); l != 1 {
		v.AddError(fmt.Errorf("<when> expect 1 attribte got: %d", l))
		return false
	}
	if t.ctx.Attribute(0).NAME().GetText() != "test" {
		v.AddError(fmt.Errorf("<when> only accept test attribte got: %s", t.ctx.Attribute(0).NAME().GetText()))
		return false
	}
	if v.test(t.ctx.Attribute(0).STRING().GetText()) {
		v.merge(t)
		return true
	}
	return false
}

func (v *visitor) visitEnd(ctx *EndContext) {
	if v.Error() != nil {
		return
	}

	t := v.stack.Peek()

	//fmt.Println("end", ctx.NAME().GetText(), v.stack.Len())
	if ctx.NAME().GetText() == t.ctx.NAME().GetText() {
		switch ctx.NAME().GetText() {
		case "if":
			v.visitIf()
			return
		case "choose":
			v.visitChoose()
			return
		case "when", "otherwise":
			v.visitWhenOtherwise()
			return
		case "trim":

		case "where":

		case "set":

		case "foreach":

		}
	}

	v.writeString(t.ctx.GetText())
	v.addVar(t.vars...)
	v.stack.Pop()
}

func (v *visitor) visitWhenOtherwise() {
	v.choose = append(v.choose, v.stack.Pop())
}

func (v *visitor) visitExpr(ctx *ExprContext) {
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

func (v *visitor) bindExpr(rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			v.addVar(rv.Index(i).Interface())
			// TODO handle count
			s = append(s, fmt.Sprintf("$%d", v.count))
		}
		v.writeString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		v.addVar(rv.Interface())
		v.writeString(fmt.Sprintf("$%d", v.count))
	}
}

func (v *visitor) explainExpr(rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			s = append(s, fmt.Sprintf("%s", v.explainVar(rv.Index(i))))
		}
		v.writeString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		v.writeString(fmt.Sprintf("%s", v.explainVar(rv)))
	}
}

func (v *visitor) explainVar(rv reflect.Value) (r string) {
	r, err := v.formatter(rv, "'")
	if err != nil {
		v.AddError(err)
		return
	}
	return
}

func (v *visitor) visitAttribute(ctx *AttributeContext) {
	if v.Error() != nil {
		return
	}
	spew.Json(ctx.GetText())
}

func (v *visitor) visitCharData(ctx *ChardataContext) {
	if v.Error() != nil {
		return
	}

	if ctx.WS() != nil {
		v.writeWS()
	} else {
		v.writeString(ctx.GetText())
	}
}

func (v *visitor) visitReference(ctx *ReferenceContext) {
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
		v.writeString(c)
	}
}
