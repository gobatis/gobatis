package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

const (
	must_prefix = "Must"
	tx_suffix   = "Tx"
	stmt_suffix = "Stmt"
)

const (
	result_none = iota
	result_result
	result_result_map
)

type fragmentManager struct {
	mu        sync.RWMutex
	list      []*fragment
	fragments map[string]*fragment
}

func newMethodManager() *fragmentManager {
	return &fragmentManager{}
}

func (p *fragmentManager) add(m *fragment) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.fragments == nil {
		p.fragments = map[string]*fragment{}
	}
	_, ok := p.fragments[m.id]
	if ok {
		return fmt.Errorf("duplicated fragment '%s'", m.id)
	}
	p.fragments[m.id] = m
	p.list = append(p.list, m)
	return nil
}

func (p *fragmentManager) all() []*fragment {
	p.mu.RLock()
	defer func() {
		p.mu.RUnlock()
	}()
	items := make([]*fragment, len(p.list))
	for i, v := range p.list {
		items[i] = v
	}
	return items
}

func (p *fragmentManager) replace(m *fragment) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.fragments == nil {
		p.fragments = map[string]*fragment{}
	}
	_, ok := p.fragments[m.id]
	if !ok {
		return fmt.Errorf("fragment '%s' not exist", m.id)
	}
	p.fragments[m.id] = m
	return nil
}

func (p *fragmentManager) get(id string) (m *fragment, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.fragments == nil {
		return
	}
	m, ok = p.fragments[id]
	return
}

type inserter struct {
	fl   []string
	fm   map[string]bool
	vs   [][]string
	rows bool
}

func (p *inserter) addField(v string) {
	if p.fm == nil {
		p.fm = map[string]bool{}
	}
	p.fm[v] = true
	p.fl = append(p.fl, v)
}

func (p *inserter) hasField(v string) bool {
	if p.fm == nil {
		return false
	}
	_, ok := p.fm[v]
	return ok
}

func (p *inserter) noField() bool {
	return p.fm == nil
}

func (p *inserter) noValue() bool {
	return len(p.vs) == 0
}

type fragment struct {
	db              *DB
	id              string
	node            *xmlNode
	in              []*param
	out             []*param
	resultAttribute int
	must            bool
	stmt            bool
	_stmt           *Stmt
}

func (p *fragment) fork() *fragment {
	n := new(fragment)
	n.db = p.db
	n.id = p.id
	n.node = p.node
	n.in = p.in
	n.out = p.out
	n.resultAttribute = p.resultAttribute
	n.must = p.must
	n.stmt = p.stmt
	n._stmt = p._stmt
	return n
}

func (p *fragment) newCaller(t reflect.Type) *caller {
	c := &caller{t: t, fragment: p}
	if p.db != nil {
		c.logger = p.db.logger
	}
	if t != nil {
		for i := 0; i < t.NumOut()-1; i++ {
			if t.Out(i).Kind() == reflect.Ptr {
				c.result = append(c.result, reflect.New(t.Out(i).Elem()))
			} else {
				c.result = append(c.result, reflect.New(t.Out(i)))
			}
		}
	}
	return c
}

func (p *fragment) proxy(field reflect.Value) {
	field.Set(reflect.MakeFunc(field.Type(), func(in []reflect.Value) []reflect.Value {
		return p.newCaller(field.Type()).call(in...).result
	}))
}

func (p *fragment) setResultAttribute() {
	if p.node.Name != dtd.SELECT &&
		p.node.Name != dtd.INSERT &&
		p.node.Name != dtd.UPDATE &&
		p.node.Name != dtd.DELETE {
		return
	}
	if p.node.HasAttribute(dtd.RESULT) {
		p.resultAttribute = result_result
	} else if p.node.HasAttribute(dtd.RESULT_MAP) {
		p.resultAttribute = result_result_map
	} else {
		p.resultAttribute = result_none
	}
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

func (p *fragment) checkResult(ft reflect.Type, mn, fn string) {
	
	if ft.NumOut() == 0 || !isErrorType(ft.Out(ft.NumOut()-1)) {
		throw(p.node.File, p.node.ctx, checkResultErr).
			format("out expect error at last, got %s", ft.Out(ft.NumOut()-1).String())
	}
	
	switch p.node.Name {
	case dtd.SELECT:
		switch p.resultAttribute {
		case result_result:
			if len(p.out) == 0 {
				if ft.NumOut() > 1 {
					switch ft.Out(0).Kind() {
					case reflect.Ptr:
						switch ft.Out(0).Elem().Kind() {
						case reflect.Struct, reflect.Map:
							return
						}
					case reflect.Slice:
						switch ft.Out(0).Elem().Kind() {
						case reflect.Ptr:
							switch ft.Out(0).Elem().Elem().Kind() {
							case reflect.Struct, reflect.Map:
								return
							}
						case reflect.Struct, reflect.Map:
							return
						}
					case reflect.Struct, reflect.Map:
						return
					}
					throw(p.node.File, p.node.ctx, checkResultErr).
						format("%s.%s out[0] expect (struct | []struct | map | []map), got: %s", mn, fn, ft.Out(0))
				}
			} else {
				for i, v := range p.out {
					if !v.expected(ft.Out(i)) {
						throw(p.node.File, p.node.ctx, varBindErr).
							format("%s.%s bind result '%s' expect '%s', got '%s'", mn, fn, v.name, v.Type(), ft.Out(i))
					}
				}
			}
			
		}
	case dtd.INSERT, dtd.UPDATE, dtd.DELETE:
		if ft.NumOut() > 1 {
			elem := ft.Out(0)
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			switch elem.Kind() {
			case reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64:
				return
			default:
				throw(p.node.File, p.node.ctx, checkResultErr).
					format("%s.%s out[0] expect integer, got: %s", mn, fn, ft.Out(0).Name())
			}
		}
	}
}

func (p *fragment) build(s *sentence, args ...reflect.Value) (err error) {
	defer func() {
		e := recover()
		err = catch(p.node.File, e)
	}()
	
	if len(p.in) != len(args) {
		throw(p.node.File, p.node.ctx, parasFragmentErr).
			format("expect %d args, got %d", len(p.in), len(args))
	}
	
	parser := newExprParser(args...)
	for i, v := range p.in {
		err = parser.paramsStack.list.Front().Next().Value.(*exprParams).bind(v, i)
		if err != nil {
			throw(p.node.File, p.node.ctx, parasFragmentErr).with(err)
		}
	}
	p.parseBlocks(parser, p.node, s)
	s.exprs = parser.exprs
	s.vars, err = parser.realVars()
	if err != nil {
		return
	}
	
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

func (p *fragment) parseSql(parser *exprParser, ctx antlr.ParserRuleContext, text string, s *sentence) {
	chars := []rune(text)
	begin := false
	inject := false
	var from int
	var next int
	var sql string
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
				sql += string(chars[i])
			}
		} else if chars[i] == 125 {
			_expr := string(chars[from:i])
			r, _, err := parser.parseExpression(ctx, _expr)
			if err != nil {
				panic(err)
			}
			if inject {
				sql += fmt.Sprintf("%v", cast.Indirect(r))
			} else {
				parser.varIndex++
				sql += fmt.Sprintf("$%d", parser.varIndex)
				parser.addVar(_expr, r)
			}
			begin = false
			inject = false
		}
	}
	
	// to avoid useless space
	s.sql += sql
}

func (p *fragment) parseBlocks(parser *exprParser, node *xmlNode, s *sentence) {
	for _, child := range node.Nodes {
		p.parseBlock(parser, child, s)
	}
	return
}

func (p *fragment) parseBlock(parser *exprParser, node *xmlNode, s *sentence) {
	if node.textOnly {
		p.parseSql(parser, node.ctx, node.Text, s)
	} else {
		s.dynamic = true
		switch node.Name {
		case dtd.IF:
			r := new(sentence)
			if p.parseTest(parser, node, r) {
				s.merge(r)
			}
		case dtd.WHERE:
			p.parseWhere(parser, node, s)
		case dtd.CHOOSE:
			p.parseChoose(parser, node, s)
		case dtd.FOREACH:
			p.parseForeach(parser, node, s)
		case dtd.TRIM:
			p.parseTrim(parser, node, s)
		case dtd.SET:
			p.parseSet(parser, node, s)
		case dtd.INSERTER:
			p.parseInserter(parser, node, s)
		}
	}
}

func (p *fragment) parseTest(parser *exprParser, node *xmlNode, s *sentence) bool {
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
	p.parseBlocks(parser, node, s)
	
	return true
}

func (p *fragment) parseWhere(parser *exprParser, node *xmlNode, s *sentence) {
	p.trimPrefixOverrides(parser, node, s, dtd.WHERE, "AND |OR ")
}

func (p *fragment) parseChoose(parser *exprParser, node *xmlNode, s *sentence) {
	var pass bool
	var oc int
	for _, child := range node.Nodes {
		if pass {
			break
		}
		switch child.Name {
		case dtd.WHEN:
			r := new(sentence)
			if p.parseTest(parser, child, r) {
				s.merge(r)
				return
			}
		case dtd.OTHERWISE:
			oc++
			p.parseBlocks(parser, child, s)
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

func (p *fragment) parseTrim(parser *exprParser, node *xmlNode, s *sentence) {
	p.trimPrefixOverrides(parser, node, s, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p *fragment) trimPrefixOverrides(parser *exprParser, node *xmlNode, res *sentence, tag, prefixes string) {
	wr := new(sentence)
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

func (p *fragment) parseSet(parser *exprParser, node *xmlNode, res *sentence) {
	p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p *fragment) parseForeach(parser *exprParser, node *xmlNode, res *sentence) {
	
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
	
	indexParam := &param{name: index, _type: reflect.Interface.String(), slice: false}
	itemParam := &param{name: item, _type: reflect.Interface.String(), slice: slice}
	
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
		res.merge(&sentence{
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
		br := new(sentence)
		p.parseBlock(parser, child, br)
		r += br.sql
	}
	*frags = append(*frags, r)
}

func (p *fragment) parseInserter(parser *exprParser, node *xmlNode, s *sentence) {
	
	//mutilple
	//bool
	//item
	//string
	
	//table, _, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TABLE))
	//if err != nil {
	//	throw(p.node.File, node.ctx, parseInserterErr).with(err)
	//}
	//item := node.GetAttribute("item")
	//index := node.GetAttribute("index")
	//multiple := item != ""
	
	//sql := fmt.Sprintf("insert into %v(%s) values(%s)",
	//	table, strings.Join(fields, ","), strings.Join(values, ","))
	//
	//p.parseSql(parser, node.ctx, sql, s)
}

func (p *fragment) parseInserterFields(parser *exprParser, node *xmlNode, it *inserter) ([]string, [][]string) {
	var err error
	var fn string
	var fv interface{}
	for _, v := range node.Nodes {
		if !it.rows && len(it.vs) == 0 {
			it.vs = append(it.vs, []string{})
		}
		switch v.Name {
		case dtd.FIELD:
			fn = v.GetAttribute(dtd.NAME)
			if fn == "*" {
				data, ok := parser.paramsStack.getVar(node.GetAttribute(dtd.DATA))
				if !ok {
					throw(p.node.File, node.ctx, parseInserterErr).with(fmt.Errorf("data not found"))
				}
				dv := toReflectValueElem(data.value)
				p.extractInserterFields(dv, it)
			} else {
				fv, _, err = parser.parseExpression(node.ctx, fn)
				if err != nil {
					throw(p.node.File, node.ctx, parseInserterErr).with(err)
				}
				it.addField(fmt.Sprintf("\"%s\"", fv))
				if !it.rows {
					it.vs[0] = append(it.vs[0], node.NodeText())
				}
			}
		}
	}
}

func (p *fragment) extractInserterFields(dv reflect.Value, it *inserter) {
	for i := 0; i < dv.NumField(); i++ {
		//dv.Field(i).Type()
		// TODO find by tag first，and convert low snake name
	}
	return
}

func (p *fragment) parseInserterRows(parser *exprParser, node *xmlNode, it *inserter) {
	data, ok := parser.paramsStack.getVar(node.GetAttribute(dtd.DATA))
	if !ok {
		throw(p.node.File, node.ctx, parseInserterErr).with(fmt.Errorf("data not found"))
	}
	dv := toReflectValueElem(data.value)
	switch dv.Kind() {
	case reflect.Slice, reflect.Array:
		var dvv reflect.Value
		for i := 0; i < dv.Len(); i++ {
			dvv = toReflectElem(dv.Index(i))
			switch dvv.Kind() {
			case reflect.Struct:
				if it.noField() {
					p.extractInserterFields(dvv, it)
				}
				//dvv.Field()
				//p.extractInserterFields(dvv,)
			case reflect.String,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64:
				
			}
		}
	default:
		throw(p.node.File, node.ctx, parseInserterErr).with(fmt.Errorf("uniterable data"))
	}
	return
}
