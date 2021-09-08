package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
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

func (p *fragment) proxy(field reflect.Value) {
	field.Set(reflect.MakeFunc(field.Type(), func(args []reflect.Value) (results []reflect.Value) {
		return p.call(field.Type(), args...)
	}))
}

func (p *fragment) call(_type reflect.Type, in ...reflect.Value) []reflect.Value {
	
	c := &caller{fragment: p, args: in, logger: p.db.logger}
	for i := 0; i < _type.NumOut()-1; i++ {
		if _type.Out(i).Kind() == reflect.Ptr {
			c.values = append(c.values, reflect.New(_type.Out(i).Elem()))
		} else {
			c.values = append(c.values, reflect.New(_type.Out(i)))
		}
	}
	
	err := c.call()
	if err != nil {
		if err == sql.ErrNoRows {
			if p.must {
				c.values = append(c.values, reflect.ValueOf(err))
			} else {
				c.values = append(c.values, reflect.Zero(errorType))
			}
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
	v,_, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TEST))
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
	logger   Logger
	args     []reflect.Value
	values   []reflect.Value
}

//func (p *caller) Scan(pointers ...interface{}) (err error) {
//	for _, v := range pointers {
//		rv := reflect.ValueOf(v)
//		if rv.Kind() != reflect.Ptr {
//			err = fmt.Errorf("scan only accept pointer")
//			return
//		}
//		p.values = append(p.values, rv)
//	}
//
//	return p.call()
//}

func (p *caller) call() (err error) {
	
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.fragment.id, time.Since(start))
	}()
	
	switch p.fragment.node.Name {
	case dtd.SELECT:
		return p.query(p.args...)
	case dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		return p.exec(p.args...)
	default:
		throw(p.fragment.node.File, p.fragment.node.ctx, callerErr).
			format("unsupported call method '%s'", p.fragment.node.Name)
		return
	}
}

func (p *caller) exec(in ...reflect.Value) (err error) {
	
	_execer, index := p.execer(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	ctx, index := p.context(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	tx, _ := _execer.(*Tx)
	if tx != nil {
		stmt := tx.getStmt(p.fragment.id)
		if stmt != nil {
			return stmt.exec(true, ctx, in)
		}
	}
	
	if p.fragment._stmt != nil {
		return p.fragment._stmt.exec(false, ctx, in)
	}
	
	var conn *sql.Conn
	if _execer == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		_execer = conn
	}
	defer func() {
		if conn != nil && p.fragment._stmt == nil {
			if _err := conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
			}
		}
	}()
	
	s, exprs, vars, dynamic, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	
	p.logger.Debugf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
	p.logger.Debugf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
	
	var res sql.Result
	if p.fragment.stmt {
		var _stmt *sql.Stmt
		_stmt, err = _execer.PrepareContext(ctx, s)
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
			return err
		}
		
		if !dynamic {
			stmt := &Stmt{
				stmt:   _stmt,
				exprs:  exprs,
				sql:    s,
				conn:   conn,
				caller: p,
			}
			if tx != nil {
				tx.addStmt(stmt)
			} else {
				p.fragment._stmt = stmt
			}
		}
		res, err = _stmt.ExecContext(ctx, vars...)
	} else {
		res, err = _execer.ExecContext(ctx, s, vars...)
	}
	if err != nil {
		p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
		p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
		p.logger.Errorf("[gobatis] [%s] exec error: %v", p.fragment.id, err)
		return
	}
	
	return p.parseExecResult(res, p.values)
}

func (p *caller) parseExecResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.fragment.must && affected != 1 {
		return fmt.Errorf("expect affect 1 row, got %d", affected)
	}
	return (&execResult{affected: affected, values: values}).scan()
}

func (p *caller) query(in ...reflect.Value) (err error) {
	
	ctx, index := p.context(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	_queryer, index := p.queryer(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	tx, _ := _queryer.(*Tx)
	if tx != nil {
		stmt := tx.getStmt(p.fragment.id)
		if stmt != nil {
			err = stmt.query(true, ctx, in, p.values)
			if err != nil {
				return
			}
			return
		}
	}
	
	if p.fragment._stmt != nil {
		err = p.fragment._stmt.query(false, ctx, in, p.values)
		if err != nil {
			return
		}
		return
	}
	
	var conn *sql.Conn
	if _queryer == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		_queryer = conn
	}
	defer func() {
		if conn != nil && p.fragment._stmt == nil {
			if _err := conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
			}
		}
	}()
	
	s, exprs, vars, dynamic, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	
	p.logger.Debugf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
	p.logger.Debugf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
	
	var rows *sql.Rows
	if p.fragment.stmt {
		var _stmt *sql.Stmt
		_stmt, err = _queryer.PrepareContext(ctx, s)
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
			return err
		}
		
		if p.fragment.stmt && !dynamic {
			stmt := &Stmt{
				stmt:   _stmt,
				exprs:  exprs,
				sql:    s,
				conn:   conn,
				caller: p,
			}
			if tx != nil {
				tx.addStmt(stmt)
			} else {
				p.fragment._stmt = stmt
			}
		}
		
		rows, err = _stmt.QueryContext(ctx, vars...)
	} else {
		rows, err = _queryer.QueryContext(ctx, s, vars...)
	}
	if err != nil {
		p.logger.Errorf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
		p.logger.Errorf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
		p.logger.Errorf("[gobatis] [%s] query error: %v", p.fragment.id, err)
		return
	}
	err = p.parseQueryResult(rows, p.values)
	if err != nil {
		return
	}
	return
}

func (p *caller) parseQueryResult(rows *sql.Rows, values []reflect.Value) (err error) {
	defer func() {
		if _err := rows.Close(); _err != nil {
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.fragment.id, _err)
		}
	}()
	
	res := queryResult{rows: rows}
	err = res.setSelected(p.fragment.resultAttribute, p.fragment.out, values)
	if err != nil {
		return err
	}
	return res.scan()
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
