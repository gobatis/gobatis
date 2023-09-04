package batis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
)

type execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type psr struct {
	sql     string
	dynamic bool
}

func (p *psr) merge(s ...*psr) {
	for _, v := range s {
		p.sql += " " + strings.TrimSpace(v.sql)
		if !p.dynamic && v.dynamic {
			p.dynamic = v.dynamic
		}
	}
}

type fragment struct {
	node *xmlNode
	in   []*param
}

func (p *fragment) fork() *fragment {
	n := new(fragment)
	n.node = p.node
	n.in = p.in
	return n
}

func (p *fragment) proxy(field reflect.Value) {
	field.Set(reflect.MakeFunc(field.Type(), func(args []reflect.Value) (results []reflect.Value) {
		return p.call(field.Type(), args...)
	}))
}

func (p *fragment) call(_type reflect.Type, in ...reflect.Value) []reflect.Value {

	c := &caller{fragment: p, args: in}
	for i := 0; i < _type.NumOut()-1; i++ {
		if _type.Out(i).Kind() == reflect.Ptr {
			c.values = append(c.values, reflect.New(_type.Out(i).Elem()))
		} else {
			c.values = append(c.values, reflect.New(_type.Out(i)))
		}
	}

	var err error
	if err != nil {
		if err == sql.ErrNoRows {
			//if p.must {
			//	c.values = append(c.values, reflect.ValueOf(err))
			//} else {
			//	c.values = append(c.values, reflect.Zero(errorType))
			//}
		} else {
			c.values = append(c.values, reflect.ValueOf(err))
		}
	} else {
		c.values = append(c.values, reflect.Zero(errorType))
	}

	for i := 0; i < _type.NumOut()-1; i++ {
		if _type.Out(i).Kind() == reflect.Ptr {
			if err == sql.ErrNoRows {
				c.values[i] = reflect.Zero(c.values[i].Type())
			}
		} else {
			c.values[i] = c.values[i].Elem()
		}
	}

	return c.values
}

func (p *fragment) checkParameter(ft reflect.Type, mn, fn string) {
	ac := 0
	for i := 0; i < ft.NumIn(); i++ {
		if isContext(ft.In(i)) || isTx(ft.In(i)) || isDB(ft.In(i)) {
			continue
		}
		ac++
	}
	if len(p.in) != ac {
		throw(p.node.File, p.node.ctx, checkParameterErr).
			format("%s.%s expected %d bind parameter, got %d", mn, fn, len(p.in), ac)
	}
}

func (p *fragment) parseStatement(args ...reflect.Value) (sql string, exprs []string, vars []interface{},
	dynamic bool, err error) {

	defer func() {
		e := recover()
		err = castRecoverError(p.node.File, e)
	}()

	if len(p.in) != len(args) {
		throw(p.node.File, p.node.ctx, parasFragmentErr).format("expect %d args, got %d", len(p.in), len(args))
	}

	parser := newExprParser(args...)
	for i, v := range p.in {
		err = parser.paramsStack.list.Front().Next().Value.(*exprParams).bind(v, i)
		if err != nil {
			throw(p.node.File, p.node.ctx, parasFragmentErr).with(err)
		}
	}
	res := new(psr)
	p.parseBlocks(parser, p.node, res)

	sql = res.sql
	vars, err = parser.realVars()
	if err != nil {
		return
	}
	exprs = parser.exprs
	dynamic = res.dynamic

	return
}

func (p *fragment) trimPrefixOverride(sql, prefix string) (r string, err error) {
	reg, err := regexp.Compile(`(?i)^` + prefix)
	if err != nil {
		return
	}
	r = reg.ReplaceAllString(sql, "")
	return
}

func (p *fragment) parseSql(parser *exprParser, node *xmlNode, res *psr) {
	chars := []rune(node.Text)
	begin := false
	inject := false
	var from int
	var next int
	var s string
	for i := 0; i < len(chars); i++ {
		if !begin {
			next = i + 1
			if (chars[i] == 35 || chars[i] == 36) && next <= len(chars)-1 && chars[next] == 123 {
				if chars[i] == 36 {
					inject = true
				}
				begin = true
				i++
				from = i + 1
				continue
			} else {
				s += string(chars[i])
			}
		} else if chars[i] == 125 {
			_expr := string(chars[from:i])
			r, _, err := parser.parseExpression(node.ctx, _expr)
			if err != nil {
				panic(err)
			}
			if inject {
				s += fmt.Sprintf("%v", cast.Indirect(r))
			} else {
				parser.varIndex++
				s += fmt.Sprintf("$%d", parser.varIndex)
				parser.addVar(_expr, r)
			}
			begin = false
			inject = false
		}
	}

	// to avoid useless space
	res.sql += s
}

func (p *fragment) parseBlocks(parser *exprParser, node *xmlNode, res *psr) {
	for _, child := range node.Nodes {
		p.parseBlock(parser, child, res)
	}
	return
}

func (p *fragment) parseBlock(parser *exprParser, node *xmlNode, res *psr) {
	if node.textOnly {
		p.parseSql(parser, node, res)
	} else {
		res.dynamic = true
		switch node.Name {
		case dtd.IF:
			r := new(psr)
			if p.parseTest(parser, node, r) {
				res.merge(r)
			}
		case dtd.WHERE:
			p.parseWhere(parser, node, res)
		case dtd.CHOOSE:
			p.parseChoose(parser, node, res)
		case dtd.FOREACH:
			p.parseForeach(parser, node, res)
		case dtd.TRIM:
			p.parseTrim(parser, node, res)
		case dtd.SET:
			p.parseSet(parser, node, res)
		}
	}
}

func (p *fragment) parseTest(parser *exprParser, node *xmlNode, res *psr) bool {
	v, _, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TEST))
	if err != nil {
		throw(p.node.File, p.node.ctx, parasFragmentErr).with(err)
	}
	b, err := cast.ToBoolE(v)
	if err != nil {
		throw(p.node.File, p.node.ctx, parasFragmentErr).with(err)
	}
	if !b {
		return false
	}
	p.parseBlocks(parser, node, res)

	return true
}

func (p *fragment) parseWhere(parser *exprParser, node *xmlNode, res *psr) {
	p.trimPrefixOverrides(parser, node, res, dtd.WHERE, "AND |OR ")
}

func (p *fragment) parseChoose(parser *exprParser, node *xmlNode, res *psr) {
	var pass bool
	var oc int
	for _, child := range node.Nodes {
		if pass {
			break
		}
		switch child.Name {
		case dtd.WHEN:
			r := new(psr)
			if p.parseTest(parser, child, r) {
				res.merge(r)
				return
			}
		case dtd.OTHERWISE:
			oc++
			p.parseBlocks(parser, child, res)
			return
		default:
			if child.textOnly {
				if strings.TrimSpace(child.Text) == "" {
					continue
				}
				child.Name = "TEXT"
			}
			throw(parser.file, child.ctx, parasFragmentErr).
				format("unsupported element '%s' element in choose", child.Name)
		}
	}
	if oc != 1 {
		throw(parser.file, node.ctx, parasFragmentErr).format("choose except 1 otherwise, got %d", oc)
	}
}

func (p *fragment) parseTrim(parser *exprParser, node *xmlNode, res *psr) {
	p.trimPrefixOverrides(parser, node, res, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p *fragment) trimPrefixOverrides(parser *exprParser, node *xmlNode, res *psr, tag, prefixes string) {
	wr := new(psr)
	p.parseBlocks(parser, node, wr)
	var err error
	s := strings.TrimSpace(wr.sql)
	filters := strings.Split(prefixes, "|")
	for _, v := range filters {
		s, err = p.trimPrefixOverride(s, v)
		if err != nil {
			throw(p.node.File, p.node.ctx, parasFragmentErr).format("regexp compile error: %s", err)
		}
	}
	res.sql = fmt.Sprintf("%s %s %s", res.sql, tag, s)
}

func (p *fragment) parseSet(parser *exprParser, node *xmlNode, res *psr) {
	p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p *fragment) parseForeach(parser *exprParser, node *xmlNode, res *psr) {

	_var := node.GetAttribute(dtd.COLLECTION)
	collection, ok := parser.paramsStack.getVar(_var)
	if !ok {
		throw(p.node.File, p.node.ctx, parasFragmentErr).
			format("can't get foreach collection '%s' value", _var)
	}
	index := node.GetAttribute(dtd.INDEX)
	if index == "" {
		index = dtd.INDEX
	}
	item := node.GetAttribute(dtd.ITEM)
	slice := false
	if item == "" {
		item = dtd.ITEM
	} else {
		item, slice = handleSlice(item)
	}
	indexParam := &param{name: index, rt: reflect.Interface.String(), slice: false}
	itemParam := &param{name: item, rt: reflect.Interface.String(), slice: slice}
	open := node.GetAttribute(dtd.OPEN)
	_close := node.GetAttribute(dtd.CLOSE)
	separator := node.GetAttribute(dtd.SEPARATOR)

	parser.paramsStack.push(newExprParams())
	elem := toReflectValueElem(collection.value)
	frags := make([]string, 0)
	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < elem.Len(); i++ {
			parser.paramsStack.peak().values = []exprValue{{value: i}, {value: elem.Index(i).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	case reflect.Map:
		for i, v := range elem.MapKeys() {
			parser.paramsStack.peak().values = []exprValue{{value: v.Interface()}, {value: elem.MapIndex(v).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	case reflect.Struct:
		for i := 0; i < elem.NumField(); i++ {
			parser.paramsStack.peak().values = []exprValue{{value: elem.Type().Field(i).Name}, {value: elem.Field(i).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	default:
		throw(parser.file, node.ctx, parasFragmentErr).format("foreach collection type '%s' can't range", elem.Type())
	}
	_, err := parser.paramsStack.pop()
	if err != nil {
		throw(parser.file, node.ctx, popParamsStackErr).with(err)
	}
	if len(frags) > 0 {
		res.merge(&psr{
			sql: open + strings.Join(frags, separator) + _close,
		})
	}
}

func (p *fragment) bindForeachParams(parser *exprParser, indexParam, itemParam *param) {
	parser.paramsStack.peak().check = map[string]int{}
	err := parser.paramsStack.peak().bind(indexParam, 0)
	if err != nil {
		throw(p.node.File, p.node.ctx, varBindErr).with(err)
	}
	err = parser.paramsStack.peak().bind(itemParam, 1)
	if err != nil {
		throw(p.node.File, p.node.ctx, varBindErr).with(err)
	}
}

func (p *fragment) parseParams(tokens string) []*param {
	parser := newParamParser(p.node.File)
	parser.walkMethods(initExprParser(tokens))
	return parser.params
}

func (p *fragment) parseForeachChild(parser *exprParser, node *xmlNode, frags *[]string) {
	r := ""
	for _, child := range node.Nodes {
		br := new(psr)
		p.parseBlock(parser, child, br)
		r += br.sql
	}
	*frags = append(*frags, r)
}

type caller struct {
	fragment *fragment
	args     []reflect.Value
	values   []reflect.Value
}

func (p *caller) removeParam(a []reflect.Value, i int) []reflect.Value {
	return append(a[:i], a[i+1:]...)
}

func (p *caller) context(in []reflect.Value) (context.Context, int) {
	for i, v := range in {
		if isContext(v.Type()) {
			return v.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
}

func (p *caller) execer(in []reflect.Value) (execer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(execer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(execer), i
			}
		}
	}
	return nil, -1
}

func (p *caller) queryer(in []reflect.Value) (queryer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(queryer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(queryer), i
			}
		}
	}
	return nil, -1
}
