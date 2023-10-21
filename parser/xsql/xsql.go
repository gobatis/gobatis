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

//type tag struct {
//	*Fragment
//	//ctx      *StartContext
//	vars     map[string]any
//	children []antlr.Tree
//	test     bool
//}

type Fragment struct {
	statement strings.Builder
	dynamic   bool
	vars      []any
	ws        bool
	count     int
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

func (x *Fragment) Count() int {
	return x.count
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
	x.count++
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

func parse(source string, formatter Formatter, vars map[string]any) (f *Fragment, err error) {

	el := &commons.ErrorListener{}

	lexer := NewXSQLLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewXSQLParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	//p.AddErrorListener(antlr.NewConsoleErrorListener())
	//p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	//p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.Document()
	if el.Error() != nil {
		err = el.Error()
		return
	}

	defer func() {
		e := recover()
		if e != nil {
			err = commons.RecoverError(e)
		}
	}()

	f = &Fragment{}

	v := &visitor{
		formatter: formatter,
		vars:      commons.NewStack[map[string]any](),
	}
	v.vars.Push(vars)
	v.VisitDocument(f, tree.Content().(*ContentContext).GetChildren())

	return
}

type visitor struct {
	formatter Formatter
	vars      *commons.Stack[map[string]any]
}

func (v *visitor) FetchVar(name string) (val any, ok bool) {
	for i := v.vars.Len() - 1; i >= 0; i-- {
		val, ok = v.vars.Index(i)[name]
		if ok {
			return
		}
	}
	return
}

func (v *visitor) VisitDocument(f *Fragment, nodes []antlr.Tree) {
	for _, node := range nodes {
		v.visitContent(f, node)
	}
}

func (v *visitor) visitContent(f *Fragment, node antlr.Tree) {
	//fmt.Println("visit content:", reflect.ValueOf(c).MethodByName("GetText").Call([]reflect.Value{})[0].Interface())
	switch t := node.(type) {
	case *ContentContext:
		v.VisitDocument(f, t.GetChildren())
	case *ElementContext:
		v.visitElement(f, t)
	case *ExprContext:
		v.visitExpr(f, t)
	case *ReferenceContext:
		v.visitReference(f, t)
	case *ChardataContext:
		v.visitCharData(f, t)
	default:
		panic(fmt.Errorf("unsupported node: %v", t))
	}
}

func (v *visitor) visitElement(f *Fragment, ctx *ElementContext) {

	switch ctx.GetName().GetText() {
	case tagForeach:
		v.enterForeach()
	case tagIf:
		v.enterIf(f, ctx)
	case tagChoose:
		v.visitChoose(f, ctx)
	case tagTrim:
		v.visitTrim(f, ctx)
	case tagWhere:
		v.visitWhere(f, ctx)
	case tagSet:
		v.visitSet(f, ctx)
	case tagWhen, tagOtherwise:
		panic(fmt.Errorf("tag <%s> should be included directly in the choose tag", ctx.GetName().GetText()))
	default:
		panic(fmt.Errorf("unsupported tag: %s", ctx.GetName().GetText()))
	}
}

//func newTag(ctx *StartContext, vars map[string]any) *tag {
//	return &tag{
//		ctx:      ctx,
//		Fragment: &Fragment{},
//		vars:     vars,
//		children: []antlr.Tree{},
//	}
//}

//func (v *visitor) visitStart(ctx *StartContext) {
//	if v.Error() != nil {
//		return
//	}
//
//	fmt.Println("visit start:", ctx.GetText())
//
//	if ctx.SLASH() != nil {
//		v.AddError(fmt.Errorf("unsupport self closed tag: %s", ctx.GetText()))
//		return
//	}
//
//	switch ctx.NAME().GetText() {
//	case tagForeach:
//		v.enterForeach()
//	case tagIf:
//		v.enterIf(ctx)
//	case tagWhen:
//		v.enterWhen(ctx)
//	case tagChoose:
//		v.enterChoose(ctx)
//	case tagOtherwise,
//		tagTrim,
//		tagWhere,
//		tagSet:
//		v.tags.Push(newTag(ctx, nil))
//	default:
//		v.AddError(fmt.Errorf("invalid tag: %s", ctx.GetText()))
//	}
//}

//func (v *visitor) visitEnd(ctx *EndContext) {
//	if v.Error() != nil {
//		return
//	}
//
//	t := v.tags.Peek()
//
//	fmt.Println("visit end:", ctx.GetText())
//	if ctx.NAME().GetText() != t.ctx.NAME().GetText() {
//		v.AddError(fmt.Errorf("tag: %s not closed, get: %s", t.ctx.GetText(), ctx.GetText()))
//		return
//	}
//
//	switch ctx.NAME().GetText() {
//	case tagIf:
//		v.exitIf()
//	case tagWhen:
//		v.exitWhen()
//	case tagForeach:
//		v.exitForeach()
//	case tagChoose:
//		v.exitChoose()
//	case tagOtherwise:
//		v.exitOtherwise()
//	case tagTrim:
//		v.exitTrim()
//	case tagWhere:
//		v.exitWhere()
//	case tagSet:
//		//v.visitSet()
//		v.exitSet()
//	default:
//		v.AddError(fmt.Errorf("unsupport close tag: %s", ctx.GetText()))
//	}
//
//	fmt.Println("visit end:", t.ctx.GetText(), "stack len:", v.tags.Len())
//}

func (v *visitor) enterIf(f *Fragment, ctx *ElementContext) {
	if l := len(ctx.AllAttribute()); l != 1 {
		panic(fmt.Errorf("<if> expect 1 attribte got: %d", l))
	}
	if ctx.Attribute(0).NAME().GetText() != "test" {
		panic(fmt.Errorf("<if> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
	}

	if !v.test(ctx.Attribute(0).STRING().GetText()) {
		return
	}
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
}

func (v *visitor) visitWhen(f *Fragment, ctx *ElementContext) bool {
	if l := len(ctx.AllAttribute()); l != 1 {
		panic(fmt.Errorf("<when> expect 1 attribte got: %d", l))
	}
	if ctx.Attribute(0).NAME().GetText() != "test" {
		panic(fmt.Errorf("<when> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
	}
	if !v.test(ctx.Attribute(0).STRING().GetText()) {
		return false
	}
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
	return true
}

func (v *visitor) enterForeach() {

}

//func (v *visitor) enterChoose(ctx *StartContext) {
//	v.mode = lazy
//	v.tags.Push(newTag(ctx, nil))
//}

func (v *visitor) exitWhen() {

}

func (v *visitor) exitForeach() {

}

func (v *visitor) exitWhere() {

}

func (v *visitor) exitSet() {

}

func (v *visitor) visitChoose(f *Fragment, ctx *ElementContext) {

	test := false
	var otherwise *ElementContext

	for _, node := range ctx.Content().GetChildren() {
		elem, ok := node.(*ElementContext)
		if ok {
			if elem.GetName().GetText() == tagOtherwise {
				if otherwise != nil {
					panic(fmt.Errorf("in the <choose> block, otherwise is allowed to appear only once"))
				}
				otherwise = elem
				continue
			}

			if elem.GetName().GetText() == tagWhen {
				if !test {
					test = v.visitWhen(f, elem)
				}
				continue
			}
		}
		v.visitContent(f, node)
	}

	if !test && otherwise != nil {
		v.visitOtherwise(f, otherwise)
	}

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
	//t := v.tags.Pop()
	//if !v.tags.Peek().test {
	//	v.merge(t)
	//}
}

func (v *visitor) test(s string) bool {
	r, err := expr.Parse(s[1:len(s)-1], v.FetchVar)
	if err != nil {
		panic(fmt.Errorf("parse test expression: %s error: %w", s, err))
	}

	if r.Kind() != reflect.Bool {
		panic(fmt.Errorf("if test expression result: %s expect bool, got: %s", s, r.Kind()))
	}
	return r.Bool()
}

func (v *visitor) visitOtherwise(f *Fragment, ctx *ElementContext) {
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
}

func (v *visitor) visitTrim(f *Fragment, ctx *ElementContext) {

	return
}

func (v *visitor) visitWhere(f *Fragment, ctx *ElementContext) {

}

func (v *visitor) visitSet(f *Fragment, ctx *ElementContext) {
	//t := v.tags.Pop()
	//s := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(t.Statement()), ","))
	//if s != "" {
	//	t.statement.Reset()
	//	t.statement.WriteString(s)
	//	v.merge(t)
	//}
}

func (v *visitor) visitForeach() {

}

func (v *visitor) visitExpr(f *Fragment, ctx *ExprContext) {
	rv, err := expr.Parse(ctx.GetVal().GetText(), v.FetchVar)
	if err != nil {
		panic(fmt.Errorf("parse expression: %s error: %w", ctx.GetVal().GetText(), err))
	}
	if ctx.HASH() != nil && v.formatter == nil {
		v.bindExpr(f, rv)
	} else {
		v.explainExpr(f, rv)
	}
}

func (v *visitor) bindExpr(f *Fragment, rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			f.addVar(rv.Index(i).Interface())
			// TODO handle count
			s = append(s, fmt.Sprintf("$%d", f.count))
		}
		f.writeString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		f.addVar(rv.Interface())
		f.writeString(fmt.Sprintf("$%d", f.count))
	}
}

func (v *visitor) explainExpr(f *Fragment, rv reflect.Value) {
	if rv.Kind() == reflect.Slice {
		var s []string
		for i := 0; i < rv.Len(); i++ {
			s = append(s, fmt.Sprintf("%s", v.explainVar(rv.Index(i))))
		}
		f.writeString(fmt.Sprintf("(%s)", strings.Join(s, ",")))
	} else {
		f.writeString(fmt.Sprintf("%s", v.explainVar(rv)))
	}
}

func (v *visitor) explainVar(rv reflect.Value) (r string) {
	r, err := v.formatter(rv, "'")
	if err != nil {
		panic(err)
	}
	return
}

func (v *visitor) visitCharData(f *Fragment, ctx *ChardataContext) {
	if ctx.WS() != nil {
		f.writeWS()
	} else {
		f.writeString(ctx.GetText())
	}
}

func (v *visitor) visitReference(f *Fragment, ctx *ReferenceContext) {
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
			panic(fmt.Errorf("unkonwn reference: %s", ctx.EntityRef().GetText()))
		}
		f.writeString(c)
	}
}
