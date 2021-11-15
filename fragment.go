package gobatis

import (
	"context"
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

type inserter struct {
	table string
	fl    []string
	fm    map[string]bool
	vs    []string
}

func (p *inserter) addField(k, v string) {
	if p.fm == nil {
		p.fm = map[string]bool{}
	}
	p.fm[k] = true
	p.vs = append(p.vs, v)
	p.fl = append(p.fl, k)
}

func (p inserter) hasField(v string) bool {
	if p.fm == nil {
		return false
	}
	_, ok := p.fm[v]
	return ok
}

func (p *inserter) removeField(k string) {
	if p.fm == nil {
		return
	}
	if _, ok := p.fm[k]; ok {
		delete(p.fm, k)
	}
	for i := range p.fl {
		if p.fl[i] == k {
			p.fl = append(p.fl[:i], p.fl[i+1:]...)
			p.vs = append(p.vs[:i], p.vs[i+1:]...)
			break
		}
	}
}

func (p inserter) empty() bool {
	return p.fm == nil
}

type blocks struct {
	items map[string]*xmlNode
}

func (p blocks) get(name string) *xmlNode {
	if p.items == nil {
		return nil
	}
	return p.items[name]
}

func (p blocks) len() int {
	return len(p.items)
}

type fragmentManager struct {
	mu        sync.RWMutex
	list      []*fragment
	fragments map[string]*fragment
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

type fragment struct {
	engine *Engine
	db     *DB
	id     string
	node   *xmlNode
	in     []*param
	out    []*param
	rt     int
	must   bool
	stmt   bool
}

func (p fragment) fork() *fragment {
	n := new(fragment)
	n.db = p.db
	n.id = p.id
	n.node = p.node
	n.in = p.in
	n.out = p.out
	n.rt = p.rt
	n.must = p.must
	return n
}

func (p fragment) proxy(mt reflect.Value) {
	mt.Set(reflect.MakeFunc(mt.Type(), func(in []reflect.Value) []reflect.Value {
		return p.newCaller(mt.Type()).call(in).result
	}))
}

func (p fragment) newCaller(mt reflect.Type) *caller {
	c := &caller{mt: mt, fragment: &p}
	if p.db != nil {
		c.logger = p.db.logger
	}
	if mt != nil {
		for i := 0; i < mt.NumOut()-1; i++ {
			if mt.Out(i).Kind() == reflect.Ptr {
				c.result = append(c.result, reflect.New(mt.Out(i).Elem()))
			} else {
				c.result = append(c.result, reflect.New(mt.Out(i)))
			}
		}
	}
	return c
}

func (p *fragment) setResultAttribute() {
	if p.node.HasAttribute(dtd.RESULT) {
		p.rt = result_result
	} else if p.node.HasAttribute(dtd.RESULT_MAP) {
		p.rt = result_result_map
	} else {
		p.rt = result_none
	}
}

func (p fragment) checkParameter(ft reflect.Type, mn, fn string) {
	ac := 0
	for i := 0; i < ft.NumIn(); i++ {
		if isContext(ft.In(i)) || isTx(ft.In(i)) || isDB(ft.In(i)) {
			continue
		}
		ac++
	}
	if len(p.in) != ac {
		throw(p.node.File, p.node.ctx, check_parameter_err).
			format("%s.%s expected %d bind parameter, got %d", mn, fn, len(p.in), ac)
	}
}

func (p fragment) checkResult(ft reflect.Type, mn, fn string) {
	if ft.NumOut() == 0 || !isError(ft.Out(ft.NumOut()-1)) {
		throw(p.node.File, p.node.ctx, check_result_err).
			format("out expect error at last, got %s", ft.Out(ft.NumOut()-1).String())
	}
	
	switch p.node.Name {
	case dtd.SELECT:
		switch p.rt {
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
					throw(p.node.File, p.node.ctx, check_result_err).
						format("%s.%s out[0] expect (struct | []struct | map | []map), got: %s", mn, fn, ft.Out(0))
				}
			} else {
				for i, v := range p.out {
					if !v.expected(ft.Out(i)) {
						throw(p.node.File, p.node.ctx, var_bind_err).
							format("%s.%s bind result '%s' expect '%s', got '%s'", mn, fn, v.name, v.Type(), ft.Out(i))
					}
				}
			}
			
		}
	case dtd.INSERT, dtd.UPDATE, dtd.DELETE:
		if ft.NumOut() > 1 {
			elem := ft.Out(0)
			if p.stmt {
				if elem.Kind() != reflect.Ptr ||
					elem.Elem().Name() != "Stmt" ||
					elem.Elem().PkgPath() != "github.com/gobatis/gobatis" {
					throw(p.node.File, p.node.ctx, check_result_err).
						format("%s.%s out[0] expect *gobatis.Stmt, got: %s", mn, fn, ft.Out(0).Name())
				}
			} else {
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
					throw(p.node.File, p.node.ctx, check_result_err).
						format("%s.%s out[0] expect integer, got: %s", mn, fn, ft.Out(0).Name())
				}
			}
		}
	}
}

func (p fragment) context(in []reflect.Value) (context.Context, int) {
	for i, v := range in {
		if isContext(v.Type()) {
			return v.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
}

func (p fragment) conn(in []reflect.Value) (conn, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(conn)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(conn), i
			}
		}
	}
	return nil, -1
}

func (p fragment) removeParam(a []reflect.Value, i int) []reflect.Value {
	return append(a[:i], a[i+1:]...)
}

func (p fragment) prepareStmt(in []reflect.Value) (s *Stmt) {
	var index int
	s = &Stmt{
		id: p.id,
		in: in,
	}
	s.ctx, index = p.context(in)
	if index > -1 {
		s.in = p.removeParam(s.in, index)
	}
	s.conn, index = p.conn(in)
	if index > -1 {
		s.in = p.removeParam(s.in, index)
	}
	return
}

func (p fragment) prepareBlocks() *blocks {
	bs := new(blocks)
	if len(p.node.Nodes) > 0 {
		bs.items = map[string]*xmlNode{}
	}
	for _, v := range p.node.Nodes {
		if v.Name == dtd.BLOCK {
			bs.items[v.GetAttribute(dtd.TYPE)] = v
		} else if !v.EmptyText() {
			throw(p.node.File, p.node.ctx, parse_fragment_err).
				with(fmt.Errorf("unsupported ohter element"))
		}
	}
	return bs
}

func (p fragment) parseStmt(parser *exprParser, s *Stmt, nodes ...*xmlNode) (err error) {
	for _, node := range nodes {
		if node == nil {
			continue
		}
		p.parseElements(parser, node, s)
	}
	s.vars, err = parser.realVars()
	if err != nil {
		return
	}
	return
}

func (p fragment) prepareParser(in []reflect.Value) (parser *exprParser, err error) {
	if len(p.in) != len(in) {
		throw(p.node.File, p.node.ctx, paras_fragment_err).
			format("expect %d args, got %d", len(p.in), len(in))
	}
	parser = newExprParser(in...)
	for i, v := range p.in {
		err = parser.paramsStack.list.Front().Next().Value.(*exprParams).bind(v, i)
		if err != nil {
			throw(p.node.File, p.node.ctx, paras_fragment_err).with(err)
		}
	}
	return
}

func (p fragment) buildStmt(in []reflect.Value) (s *Stmt, err error) {
	defer func() {
		err = catch(p.node.File, recover())
	}()
	s = p.prepareStmt(in)
	s.query = p.node.Name == dtd.SELECT
	var parser *exprParser
	parser, err = p.prepareParser(s.in)
	if err != nil {
		return
	}
	err = p.parseStmt(parser, s, p.node)
	if err != nil {
		return
	}
	return
}

func (p fragment) buildQuery(in []reflect.Value) (ss []*Stmt, err error) {
	defer func() {
		err = catch(p.node.File, recover())
	}()
	var parser *exprParser
	ss = make([]*Stmt, 2)
	bs := p.prepareBlocks()
	cn := bs.get(dtd.BLOCK_COUNT)
	if cn != nil {
		ss[0] = p.prepareStmt(in)
		ss[0].query = true
	}
	sn := bs.get(dtd.BLOCK_SELECT)
	if sn != nil {
		ss[1] = p.prepareStmt(in)
		ss[1].query = true
	}
	fn := bs.get(dtd.BLOCK_SOURCE)
	ln := bs.get(dtd.BLOCK_PAGING)
	if ss[0] != nil {
		parser, err = p.prepareParser(ss[0].in)
		if err != nil {
			return
		}
		err = p.parseStmt(parser, ss[0], cn, fn)
		if err != nil {
			return
		}
	}
	if ss[1] != nil {
		parser, err = p.prepareParser(ss[1].in)
		if err != nil {
			return
		}
		err = p.parseStmt(parser, ss[1], sn, fn, ln)
		if err != nil {
			return
		}
	}
	return
}

func (p fragment) buildSave(in []reflect.Value) (s *Stmt, err error) {
	defer func() {
		err = catch(p.node.File, recover())
	}()
	s = p.prepareStmt(in)
	parser, err := p.prepareParser(s.in)
	if err != nil {
		return
	}
	bs := p.prepareBlocks()
	n := bs.get(dtd.BLOCK_INSERT)
	update := true
	if n != nil {
		var v interface{}
		v, _, err = parser.parseExpression(n.ctx, n.GetAttribute(dtd.TEST))
		if err != nil {
			return
		}
		var ok bool
		ok, err = cast.ToBoolE(v)
		if err != nil {
			return
		}
		if ok {
			update = false
			err = p.parseStmt(parser, s, n)
			if err != nil {
				return
			}
		}
	}
	if update {
		n = bs.get(dtd.BLOCK_SELECT)
		if n != nil {
			err = p.parseStmt(parser, s, n)
			if err != nil {
				return
			}
		}
	}
	return
}

func (p fragment) parseElements(parser *exprParser, node *xmlNode, s *Stmt) {
	for _, child := range node.Nodes {
		p.parseElement(parser, node, child, s)
	}
	return
}

func (p fragment) parseElement(parser *exprParser, parent, node *xmlNode, s *Stmt) {
	if node.plain {
		p.parseSQL(parser, node.ctx, node.Text, s)
	} else {
		s.dynamic = true
		switch node.Name {
		case dtd.IF:
			_s := new(Stmt)
			if p.parseTest(parser, node, _s) {
				s.concatSQL(_s.sql)
			}
		case dtd.BIND:
			p.parseBind(parser, node)
		case dtd.SELECT_KEY:
			p.parseSelectKey(parser, node)
		case dtd.WHERE:
			p.parseWhere(parser, node, s)
		case dtd.CHOOSE:
			p.parseChoose(parser, node, s)
		case dtd.FOREACH:
			p.parseForeach(parser, parent, node, s)
		case dtd.TRIM:
			p.parseTrim(parser, node, s)
		case dtd.SET:
			p.parseSet(parser, node, s)
		case dtd.INSERTER:
			p.parseInserter(parser, node, s)
		case dtd.BLOCK:
			// pass
		default:
			throw(node.File, node.ctx, paras_fragment_err).with(fmt.Errorf("unknown tag: %s", node.Name))
		}
	}
}

func (p fragment) trimPrefixOverride(sql, prefix string) (r string, err error) {
	reg, err := regexp.Compile(`(?i)^` + prefix)
	if err != nil {
		return
	}
	r = reg.ReplaceAllString(sql, "")
	return
}

func (p fragment) parseSQL(parser *exprParser, ctx antlr.ParserRuleContext, text string, s *Stmt) {
	if text == "" {
		return
	}
	var (
		chars  = []rune(text)
		begin  = false
		inject = false
		from   int
		next   int
		sql    string
	)
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
	s.concatSQL(sql)
}

func (p fragment) parseBind(parser *exprParser, node *xmlNode) {
	name := node.GetAttribute(dtd.NAME)
	// TODO check var rule
	value, _, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.VALUE))
	if err != nil {
		throw(node.File, node.ctx, parser_bind_err).with(err)
	}
	parser.paramsStack.peak().set(name, exprValue{
		value: value,
	})
}

func (p fragment) parseSelectKey(parser *exprParser, node *xmlNode) {
	// TODO implements
}

func (p fragment) parseTest(parser *exprParser, node *xmlNode, s *Stmt) bool {
	v, _, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TEST))
	if err != nil {
		throw(p.node.File, p.node.ctx, paras_fragment_err).with(err)
	}
	b, err := cast.ToBoolE(v)
	if err != nil {
		throw(p.node.File, p.node.ctx, paras_fragment_err).with(err)
	}
	if !b {
		return false
	}
	p.parseElements(parser, node, s)
	return true
}

func (p fragment) parseWhere(parser *exprParser, node *xmlNode, s *Stmt) {
	p.trimPrefixOverrides(parser, node, s, dtd.WHERE, "AND |OR ")
}

func (p fragment) parseChoose(parser *exprParser, node *xmlNode, s *Stmt) {
	var pass bool
	var oc int
	for _, child := range node.Nodes {
		if pass {
			break
		}
		switch child.Name {
		case dtd.WHEN:
			_s := new(Stmt)
			if p.parseTest(parser, child, _s) {
				s.concatSQL(_s.sql)
				return
			}
		case dtd.OTHERWISE:
			oc++
			p.parseElements(parser, child, s)
			return
		default:
			if child.plain {
				if strings.TrimSpace(child.Text) == "" {
					continue
				}
				child.Name = "TEXT"
			}
			throw(parser.file, child.ctx, paras_fragment_err).
				format("unsupported element '%s' element in choose", child.Name)
		}
	}
	if oc != 1 {
		throw(parser.file, node.ctx, paras_fragment_err).format("choose except 1 otherwise, got %d", oc)
	}
}

func (p fragment) parseTrim(parser *exprParser, node *xmlNode, s *Stmt) {
	p.trimPrefixOverrides(parser, node, s, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p fragment) trimPrefixOverrides(parser *exprParser, node *xmlNode, res *Stmt, tag, prefixes string) {
	wr := new(Stmt)
	p.parseElements(parser, node, wr)
	var err error
	s := strings.TrimSpace(wr.sql)
	filters := strings.Split(prefixes, "|")
	for _, v := range filters {
		s, err = p.trimPrefixOverride(s, v)
		if err != nil {
			throw(p.node.File, p.node.ctx, paras_fragment_err).format("regexp compile error: %s", err)
		}
	}
	if s != "" {
		res.sql = fmt.Sprintf("%s %s %s", res.sql, tag, s)
	}
}

func (p fragment) parseSet(parser *exprParser, node *xmlNode, res *Stmt) {
	p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p fragment) parseForeach(parser *exprParser, parent, node *xmlNode, s *Stmt) {
	
	ca := node.GetAttribute(dtd.COLLECTION)
	cv, ok := parser.paramsStack.getVar(ca)
	if !ok {
		throw(p.node.File, p.node.ctx, paras_fragment_err).
			format("can't get foreach collection '%s' value", ca)
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
	
	oa := node.GetAttribute(dtd.OPEN)
	cla := node.GetAttribute(dtd.CLOSE)
	separator := node.GetAttribute(dtd.SEPARATOR)
	
	parser.paramsStack.push(newExprParams())
	elem := toReflectValueElem(cv.value)
	frags := make([]string, 0)
	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < elem.Len(); i++ {
			parser.paramsStack.peak().values = []exprValue{{value: i}, {value: elem.Index(i).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, parent, node, &frags)
		}
	case reflect.Map:
		for i, v := range elem.MapKeys() {
			parser.paramsStack.peak().values = []exprValue{{value: v.Interface()}, {value: elem.MapIndex(v).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, parent, node, &frags)
		}
	case reflect.Struct:
		for i := 0; i < elem.NumField(); i++ {
			parser.paramsStack.peak().values = []exprValue{{value: elem.Type().Field(i).Name}, {value: elem.Field(i).Interface()}}
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, parent, node, &frags)
		}
	default:
		throw(parser.file, node.ctx, paras_fragment_err).format("foreach collection type '%s' can't range", elem.Type())
	}
	_, err := parser.paramsStack.pop()
	if err != nil {
		throw(parser.file, node.ctx, pop_params_err).with(err)
	}
	if len(frags) > 0 {
		s.concatSQL(oa + strings.Join(frags, separator) + cla)
	}
}

func (p fragment) bindForeachParams(parser *exprParser, indexParam, itemParam *param) {
	parser.paramsStack.peak().check = map[string]int{}
	err := parser.paramsStack.peak().bind(indexParam, 0)
	if err != nil {
		throw(p.node.File, p.node.ctx, var_bind_err).with(err)
	}
	err = parser.paramsStack.peak().bind(itemParam, 1)
	if err != nil {
		throw(p.node.File, p.node.ctx, var_bind_err).with(err)
	}
}

func (p fragment) parseForeachChild(parser *exprParser, parent, node *xmlNode, frags *[]string) {
	r := ""
	for _, child := range node.Nodes {
		br := new(Stmt)
		p.parseElement(parser, parent, child, br)
		r += br.sql
	}
	*frags = append(*frags, r)
}

func (p fragment) parseInserter(parser *exprParser, node *xmlNode, s *Stmt) {
	var it = new(inserter)
	var err error
	table, _, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TABLE))
	if err != nil {
		throw(p.node.File, node.ctx, parse_inserter_err).with(err)
	}
	it.table = table.(string)
	p.parseInserterFields(parser, node, it)
	p.buildInserterSQL(parser, node.ctx, it, s)
}

func (p fragment) parseInserterFields(parser *exprParser, node *xmlNode, it *inserter) {
	var err error
	var fn string
	var fv interface{}
	for _, v := range node.Nodes {
		switch v.Name {
		case dtd.FIELD:
			fn = v.GetAttribute(dtd.NAME)
			if fn == "*" {
				entity, ok := parser.paramsStack.getVar(node.GetAttribute(dtd.ENTITY))
				if !ok {
					throw(p.node.File, node.ctx, parse_inserter_err).
						with(fmt.Errorf("inserter entity attribute not defined"))
				}
				if it.empty() {
					dv := toReflectValueElem(entity.value)
					p.extractInserterFields(parser, dv, it)
				}
			} else {
				fv, _, err = parser.parseExpression(node.ctx, fn)
				if err != nil {
					throw(p.node.File, node.ctx, parse_inserter_err).with(err)
				}
				name := fv.(string)
				if !it.hasField(name) {
					it.addField(name, v.NodeText())
				}
			}
		case dtd.EXCLUDE:
			fv, _, err = parser.parseExpression(node.ctx, v.GetAttribute(dtd.NAME))
			if err != nil {
				throw(p.node.File, node.ctx, parse_inserter_err).with(err)
			}
			name := fv.(string)
			it.removeField(name)
		}
	}
}

// extract reflect entity fields as insert fields
func (p fragment) extractInserterFields(parser *exprParser, dv reflect.Value, it *inserter) {
	switch dv.Kind() {
	case reflect.Slice, reflect.Array:
		if dv.Len() > 0 {
			p.extractInserterFields(parser, dv.Index(0), it)
		}
	case reflect.Struct:
		var name string
		var iv string
		for i := 0; i < dv.Type().NumField(); i++ {
			name = p.extractFiledName(dv.Type().Field(i))
			iv = innerVar(p.id, name)
			parser.paramsStack.peak().set(iv, exprValue{
				value: dv.Field(i).Interface(),
			})
			it.addField(name, innerExpr(iv))
		}
	default:
		throw(p.node.File, p.node.ctx, parse_inserter_err).with(fmt.Errorf("unsuport inserter data type"))
	}
	return
}

func (p fragment) quoteFiled(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

// extract reflect entity field to sql fields
// if tag found, use tag
// or use lower_snake_name
func (p fragment) extractFiledName(field reflect.StructField) string {
	tag := field.Tag.Get(p.reflectTag())
	if strings.Contains(tag, ",") {
		tag = strings.TrimSpace(strings.Split(tag, ",")[0])
	}
	if tag != "" {
		return tag
	}
	return snake(field.Name)
}

func (p fragment) reflectTag() string {
	if p.engine != nil {
		return p.engine.ReflectTag()
	}
	return default_tag
}

// build inserter sql
func (p fragment) buildInserterSQL(parser *exprParser, ctx antlr.ParserRuleContext, it *inserter, s *Stmt) {
	fl := make([]string, len(it.fl))
	for i, v := range it.fl {
		fl[i] = p.quoteFiled(v)
	}
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("insert into %s(%s)", it.table, strings.Join(fl, ",")))
	str.WriteString(fmt.Sprintf(" values(%s)", strings.Join(it.vs, ",")))
	p.parseSQL(parser, ctx, str.String(), s)
}
