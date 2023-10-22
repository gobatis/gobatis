package xsql

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser"
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
	attributeTest            = "test"
	attributeItem            = "item"
	attributeCollection      = "collection"
	attributeOpen            = "open"
	attributeClose           = "close"
	attributeSeparator       = "separator"
	attributeIndex           = "index"
	attributePrefix          = "prefix"
	attributeSuffix          = "suffix"
	attributePrefixOverrides = "prefixOverrides"
	attributeSuffixOverrides = "suffixOverrides"
)

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

func (x *Fragment) writeSpace() {
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

	el := &parser.ErrorListener{}

	lexer := NewXSQLLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewXSQLParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.Document()
	if el.Error() != nil {
		err = el.Error()
		return
	}

	defer func() {
		e := recover()
		if e != nil {
			err = parser.RecoverError(e)
		}
	}()

	f = &Fragment{}

	v := &visitor{
		formatter: formatter,
		vars:      NewStack[map[string]any](),
	}
	v.vars.Push(vars)
	v.VisitDocument(f, tree.Content().(*ContentContext).GetChildren())

	return
}

type visitor struct {
	formatter Formatter
	vars      *Stack[map[string]any]
}

func (v visitor) FetchVar(name string) any {
	for i := v.vars.Len() - 1; i >= 0; i-- {
		r, ok := v.vars.Index(i)[name]
		if ok {
			return r
		}
	}
	return nil
}

func (v visitor) VisitDocument(f *Fragment, nodes []antlr.Tree) {
	for _, node := range nodes {
		v.visitContent(f, node)
	}
}

func (v visitor) visitContent(f *Fragment, node antlr.Tree) {
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

func (v visitor) visitElement(f *Fragment, ctx *ElementContext) {

	switch ctx.GetName().GetText() {
	case tagForeach:
		v.visitForeach(f, ctx)
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

func (v visitor) enterIf(f *Fragment, ctx *ElementContext) {
	if l := len(ctx.AllAttribute()); l != 1 {
		panic(fmt.Errorf("<if> expect 1 attribte got: %d", l))
	}
	if ctx.Attribute(0).NAME().GetText() != attributeTest {
		panic(fmt.Errorf("<if> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
	}

	if !v.test(ctx.Attribute(0).STRING().GetText()) {
		return
	}
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
}

func (v visitor) visitWhen(f *Fragment, ctx *ElementContext) bool {
	if l := len(ctx.AllAttribute()); l != 1 {
		panic(fmt.Errorf("<when> expect 1 attribte got: %d", l))
	}
	if ctx.Attribute(0).NAME().GetText() != attributeTest {
		panic(fmt.Errorf("<when> only accept test attribte got: %s", ctx.Attribute(0).NAME().GetText()))
	}
	if !v.test(ctx.Attribute(0).STRING().GetText()) {
		return false
	}
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
	return true
}

func (v visitor) visitChoose(f *Fragment, ctx *ElementContext) {

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
}

func (v visitor) test(s string) bool {
	r, err := expr.Parse(s[1:len(s)-1], v.FetchVar)
	if err != nil {
		panic(fmt.Errorf("parse test expression: %s error: %w", s, err))
	}

	if r.Kind() != reflect.Bool {
		panic(fmt.Errorf("if test expression result: %s expect bool, got: %s", s, r.Kind()))
	}
	return r.Bool()
}

func (v visitor) visitOtherwise(f *Fragment, ctx *ElementContext) {
	v.VisitDocument(f, ctx.Content().(*ContentContext).GetChildren())
}

func (v visitor) visitTrim(f *Fragment, ctx *ElementContext) {

	attrs := map[string]string{}
	for _, vv := range ctx.AllAttribute() {
		if _, ok := attrs[vv.NAME().GetText()]; ok {
			panic(fmt.Errorf("duplicated  tag: %s attribute: %s", ctx.GetName().GetText(), vv.NAME().GetText()))
		}
		attrs[vv.NAME().GetText()] = vv.STRING().GetText()[1 : len(vv.STRING().GetText())-1]
	}
	var (
		prefix         string
		suffix         string
		prefixOverride string
		suffixOverride string
	)
	for k, vv := range attrs {
		switch k {
		case attributePrefix:
			prefix = vv
		case attributeSuffix:
			suffix = vv
		case attributePrefixOverrides:
			prefixOverride = vv
		case attributeSuffixOverrides:
			suffixOverride = vv
		default:
			panic(fmt.Errorf("unsupported <trim> attribute: %s", vv))
		}
	}

	if prefix != "" {
		f.writeString(prefix)
	}
	n := &Fragment{}
	v.VisitDocument(n, ctx.Content().GetChildren())
	s := strings.TrimSpace(n.Statement())

	if prefixOverride != "" {
		fmt.Println("prefix rule:", fmt.Sprintf("`%s`", prefixOverride), fmt.Sprintf("^(?i)(%s)\\b", v.splitOverrideRule(prefixOverride)))
		reg, err := regexp.Compile(fmt.Sprintf("^(?i)(%s)\\b", v.splitOverrideRule(prefixOverride)))
		if err != nil {
			panic(fmt.Errorf("prepare <trim> prefixOverride: %s regex rule error: %s", prefixOverride, err))
		}
		s = reg.ReplaceAllString(s, "")
	}

	if suffixOverride != "" {
		reg, err := regexp.Compile(fmt.Sprintf("(?i)\\b(%s)$", v.splitOverrideRule(suffixOverride)))
		if err != nil {
			panic(fmt.Errorf("prepare <trim> suffixOverride: %s regex rule error: %s", suffixOverride, err))
		}
		s = reg.ReplaceAllString(s, "")
	}
	s = strings.TrimSpace(s)
	f.writeSpace()
	f.writeString(s)
	f.writeSpace()
	f.addVar(n.Vars()...)
	if suffix != "" {
		f.writeString(suffix)
	}

	return
}

func (v visitor) splitOverrideRule(s string) string {
	words := strings.Split(s, "|")
	var list []string
	for _, w := range words {
		list = append(list, w)
	}
	return strings.Join(list, "|")
}

func (v visitor) regexpTrim(r, s string) string {
	re := regexp.MustCompile(r)
	return re.ReplaceAllString(s, "")
}

func (v visitor) visitWhere(f *Fragment, ctx *ElementContext) {

	n := &Fragment{}

	v.VisitDocument(n, ctx.Content().GetChildren())

	s := strings.TrimSpace(n.Statement())
	s = strings.TrimSpace(v.regexpTrim(`^(?i)(and|or)\b`, s))
	if s != "" {
		f.writeSpace()
		f.writeString("where")
		f.writeSpace()
		f.writeString(s)
		f.addVar(n.vars...)
	}
}

func (v visitor) visitSet(f *Fragment, ctx *ElementContext) {
	ff := &Fragment{}
	v.VisitDocument(ff, ctx.Content().GetChildren())

	s := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(ff.Statement()), ","))
	if s != "" {
		f.writeSpace()
		f.writeString("set")
		f.writeSpace()
		f.writeString(s)
		f.addVar(ff.vars...)
	}
}

func (v visitor) visitForeach(f *Fragment, ctx *ElementContext) {
	attrs := map[string]string{}
	for _, vv := range ctx.AllAttribute() {
		if _, ok := attrs[vv.NAME().GetText()]; ok {
			panic(fmt.Errorf("duplicated  tag: %s attribute: %s", ctx.GetName().GetText(), vv.NAME().GetText()))
		}
		attrs[vv.NAME().GetText()] = vv.STRING().GetText()[1 : len(vv.STRING().GetText())-1]
	}
	var (
		collection string
		item       string
		open       string
		separator  string
		index      string
		_close     string
	)
	for k, vv := range attrs {
		switch k {
		case attributeSeparator:
			separator = vv
		case attributeOpen:
			open = vv
		case attributeClose:
			_close = vv
		case attributeCollection:
			collection = vv
		case attributeIndex:
			// TODO check var name
			index = vv
		case attributeItem:
			// TODO check var name
			item = vv
		default:
			panic(fmt.Errorf("unsupported <foreach> attribute: %s", vv))
		}
	}
	if collection == "" {
		panic(fmt.Errorf("<foreach> requrie attribute: %s", attributeCollection))
	}
	if item == "" {
		panic(fmt.Errorf("<foreach> requrie attribute: %s", attributeItem))
	}

	f.writeString(open)
	rv := reflect.ValueOf(v.FetchVar(collection))
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			sv := rv.Index(i)
			// TODO Check sv invalid
			v.visitForeachContent(f, ctx, index, item, separator, i, sv.Interface(), i == rv.Len()-1)
		}
	case reflect.Map:
		keys := rv.MapKeys()
		i := -1
		for _, key := range keys {
			i++
			mv := rv.MapIndex(key)
			// TODO Check sv invalid
			v.visitForeachContent(f, ctx, index, item, separator, key, mv.Interface(), i == rv.Len()-1)
		}
	default:
		panic(fmt.Errorf("unsupported <foreach> collection  type: %s", rv.Kind()))
	}
	f.writeString(_close)
}

func (v visitor) visitForeachContent(f *Fragment, ctx *ElementContext, index, item, separator string, ri, rv any, last bool) {
	v.vars.Push(map[string]any{})
	if index != "" {
		v.vars.Peek()[index] = ri
	}
	v.vars.Peek()[item] = rv
	v.VisitDocument(f, ctx.Content().GetChildren())
	v.vars.Pop()

	if separator != "" && !last {
		f.writeString(separator)
	}
}

func (v visitor) visitExpr(f *Fragment, ctx *ExprContext) {
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

func (v visitor) bindExpr(f *Fragment, rv reflect.Value) {
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

func (v visitor) explainExpr(f *Fragment, rv reflect.Value) {
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

func (v visitor) explainVar(rv reflect.Value) (r string) {
	r, err := v.formatter(rv, "'")
	if err != nil {
		panic(err)
	}
	return
}

func (v visitor) visitCharData(f *Fragment, ctx *ChardataContext) {
	if ctx.WS() != nil {
		f.writeSpace()
	} else {
		f.writeString(ctx.GetText())
	}
}

func (v visitor) visitReference(f *Fragment, ctx *ReferenceContext) {
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
