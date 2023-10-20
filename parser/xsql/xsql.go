package xsql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gobatis/gobatis/parser/expr"
)

const (
	tagIf        = "if"
	tagChoose    = "choose"
	tagWhen      = "when"
	tagOtherwise = "otherwise"
	tagTrim      = "trim"
	tagWhere     = "where"
	tagSet       = "set"
	tagForeach   = "foreach"
)

const (
	accept = iota
	reject
	lazy
)

type tag struct {
	*Fragment
	ctx      *StartContext
	vars     map[string]any
	children []antlr.Tree
	test     bool
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

func Parse(formatter Formatter, source string, vars map[string]any) (*Fragment, error) {
	return parse(source, formatter, vars)
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
		formatter:     formatter,
		tags:          commons.NewStack[*tag](),
	}
	v.tags.Push(newTag(nil, vars))
	v.VisitContent(tree.(*ContentContext).GetChildren())
	if v.Error() != nil {
		return nil, v.Error()
	}

	if v.tags.Len() != 1 {
		return nil, fmt.Errorf("expact 1 tag in stack, got: %d", v.tags.Len())
	}

	return v.tags.Peek().Fragment, errs.Error()
}

type visitor struct {
	*commons.ErrorListener
	count     int
	formatter Formatter
	tags      *commons.Stack[*tag]
	mode      int
}

func (v *visitor) writeWS() {
	//v.stack.Peek().writeWS()
	v.tags.Peek().writeWS()
}

func (v *visitor) writeString(vv string) {
	v.tags.Peek().writeString(vv)
}

func (v *visitor) addVar(vv ...any) {
	v.tags.Peek().addVar(vv...)
	v.count += len(vv)
}

func (v *visitor) addChildren(node antlr.Tree) {
	v.tags.Peek().children = append(v.tags.Peek().children, node)
}

func (v *visitor) FetchVar(name string) (val any, ok bool) {
	for i := v.tags.Len() - 1; i >= 0; i-- {
		val, ok = v.tags.Index(i).vars[name]
		if ok {
			return
		}
	}
	return
}

func (v *visitor) VisitContent(nodes []antlr.Tree) {
	for _, c := range nodes {
		if v.Error() != nil {
			return
		}

		fmt.Println("visit content:", v.mode, reflect.ValueOf(c).MethodByName("GetText").Call([]reflect.Value{})[0].Interface())

		if v.mode == lazy {
			if t, ok := c.(*EndContext); ok {
				// TODO 退出层级需要对等
				if t.NAME().GetText() == v.tags.Peek().ctx.NAME().GetText() {
					v.mode = accept
				}
			}
		}

		switch v.mode {
		case reject:
			continue
		case lazy:
			v.addChildren(c)
		default:
			switch t := c.(type) {
			case *StartContext:
				v.visitStart(t)
			case *EndContext:
				v.visitEnd(t)
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
}

func newTag(ctx *StartContext, vars map[string]any) *tag {
	return &tag{
		ctx:      ctx,
		Fragment: &Fragment{},
		vars:     vars,
		children: []antlr.Tree{},
	}
}

func (v *visitor) visitStart(ctx *StartContext) {
	if v.Error() != nil {
		return
	}

	fmt.Println("visit start:", ctx.GetText())

	if ctx.SLASH() != nil {
		v.AddError(fmt.Errorf("unsupport self closed tag: %s", ctx.GetText()))
		return
	}

	switch ctx.NAME().GetText() {
	case tagForeach:
		v.enterForeach()
	case tagIf:
		v.enterIf(ctx)
	case tagWhen:
		v.enterWhen(ctx)
	case tagChoose:
		v.enterChoose(ctx)
	case tagOtherwise,
		tagTrim,
		tagWhere,
		tagSet:
		v.tags.Push(newTag(ctx, nil))
	default:
		v.AddError(fmt.Errorf("invalid tag: %s", ctx.GetText()))
	}
}

func (v *visitor) visitEnd(ctx *EndContext) {
	if v.Error() != nil {
		return
	}

	t := v.tags.Peek()

	fmt.Println("visit end:", ctx.GetText())
	if ctx.NAME().GetText() != t.ctx.NAME().GetText() {
		v.AddError(fmt.Errorf("tag: %s not closed, get: %s", t.ctx.GetText(), ctx.GetText()))
		return
	}

	switch ctx.NAME().GetText() {
	case tagIf:
		v.exitIf()
	case tagWhen:
		v.exitWhen()
	case tagForeach:
		v.exitForeach()
	case tagChoose:
		v.exitChoose()
	case tagOtherwise:
		v.exitOtherwise()
	case tagTrim:
		v.exitTrim()
	case tagWhere:
		v.exitWhere()
	case tagSet:
		//v.visitSet()
		v.exitSet()
	default:
		v.AddError(fmt.Errorf("unsupport close tag: %s", ctx.GetText()))
	}

	fmt.Println("visit end:", t.ctx.GetText(), "stack len:", v.tags.Len())
}

func (v *visitor) enterIf(ctx *StartContext) {
	if l := len(ctx.AllAttribute()); l != 1 {
		v.AddError(fmt.Errorf("<if> expect 1 attribte got: %d", l))
		return
	}
	if ctx.Attribute(0).NAME().GetText() != "test" {
		v.AddError(fmt.Errorf("<if> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
		return
	}

	if v.test(ctx.Attribute(0).STRING().GetText()) {
		v.tags.Push(newTag(ctx, nil))
	} else {
		v.mode = reject
	}
}

func (v *visitor) exitIf() {
	if v.mode != accept {
		v.mode = accept
		return
	}
	v.merge(v.tags.Pop())
}

func (v *visitor) enterWhen(ctx *StartContext) {
	if l := len(ctx.AllAttribute()); l != 1 {
		v.AddError(fmt.Errorf("<when> expect 1 attribte got: %d", l))
		return
	}
	if ctx.Attribute(0).NAME().GetText() != "test" {
		v.AddError(fmt.Errorf("<when> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
		return
	}
	if v.test(ctx.Attribute(0).STRING().GetText()) {
		v.merge(v.tags.Pop())
		v.tags.Peek().test = true
	} else {
		v.mode = reject
	}
	return
}

func (v *visitor) enterForeach() {

}

func (v *visitor) enterChoose(ctx *StartContext) {
	v.mode = lazy
	v.tags.Push(newTag(ctx, nil))
}

func (v *visitor) exitWhen() {

}

func (v *visitor) exitForeach() {

}

func (v *visitor) exitWhere() {

}

func (v *visitor) exitSet() {

}

func (v *visitor) exitChoose() {
	v.mode = accept
	v.VisitContent(v.tags.Peek().children)
	v.merge(v.tags.Pop())

	//	var otherwise *tag
	//	var whens []*tag
	//	for i := 0; i < len(v.tmp); i++ {
	//		if v.tmp[i].ctx.NAME().GetText() == "otherwise" {
	//			if otherwise != nil {
	//				v.AddError(fmt.Errorf("in the choose block, otherwise is allowed to appear only once"))
	//				return
	//			}
	//			otherwise = v.tmp[i]
	//		} else {
	//			whens = append(whens, v.tmp[i])
	//		}
	//	}
	//
	//	v.tags.Pop()
	//	for _, vv := range whens {
	//		if v.visitWhen(vv) || v.Error() != nil {
	//			return
	//		}
	//	}
	//
	//	if otherwise != nil {
	//		v.merge(otherwise)
	//	}
	//
	//	v.tmp = make([]*tag, 0)
}

func (v *visitor) exitTrim() {

}

func (v *visitor) exitOtherwise() {
	t := v.tags.Pop()
	if !v.tags.Peek().test {
		v.merge(t)
	}
}

func (v *visitor) test(s string) bool {
	r, err := expr.Parse(s[1:len(s)-1], v.FetchVar)
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

func (v *visitor) merge(t *tag) {
	v.writeString(t.Statement())
	v.addVar(t.Fragment.vars...)
}

//func (v *visitor) visitWhenOtherwise() {
//	v.tmp = append(v.tmp, v.tags.Pop())
//}

func (v *visitor) visitTrim() {

	return
}

func (v *visitor) visitWhere() {

}

func (v *visitor) visitSet() {
	t := v.tags.Pop()
	s := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(t.Statement()), ","))
	if s != "" {
		t.statement.Reset()
		t.statement.WriteString(s)
		v.merge(t)
	}
}

func (v *visitor) visitForeach() {

}

func (v *visitor) visitExpr(ctx *ExprContext) {
	if v.Error() != nil {
		return
	}
	rv, err := expr.Parse(ctx.GetVal().GetText(), v.FetchVar)
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
