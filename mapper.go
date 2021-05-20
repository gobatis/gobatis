package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
)

var errorType reflect.Type

func isErrorType(_type reflect.Type) bool {
	return _type.Implements(reflect.TypeOf((*error)(nil)).Elem())
}

func init() {
	errorType = reflect.TypeOf((*error)(nil)).Elem()
}

type fragmentManager struct {
	mu        sync.RWMutex
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
	return nil
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
}

type queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type psr struct {
	sql       string
	cacheable bool
}

func (p *psr) merge(s ...string) {
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v != "" {
			if p.sql != "" {
				p.sql += " " + v
			} else {
				p.sql += v
			}
		}
	}
}

const (
	fragment_result_type = iota + 1
	fragment_result_map
	fragment_result
)

type fragment struct {
	db         *DB
	logger     Logger
	id         string
	statement  *xmlNode
	cacheable  bool
	sql        string
	in         []*param
	out        []*param
	resultType *resultType
	resultAttr int
}

func (p *fragment) proxy(field reflect.Value) {
	field.Set(reflect.MakeFunc(field.Type(), func(args []reflect.Value) (results []reflect.Value) {
		return p.call(field.Type(), args...)
	}))
}

func (p *fragment) call(_type reflect.Type, in ...reflect.Value) []reflect.Value {
	
	c := &caller{fragment: p, params: in}
	for i := 0; i < _type.NumOut()-1; i++ {
		if _type.Out(i).Kind() == reflect.Ptr {
			c.values = append(c.values, reflect.New(_type.Out(i).Elem()))
		} else {
			c.values = append(c.values, reflect.New(_type.Out(i)))
		}
	}
	
	err := c.call()
	if err != nil {
		c.values = append(c.values, reflect.ValueOf(err))
	} else {
		c.values = append(c.values, reflect.Zero(errorType))
	}
	
	for i := 0; i < _type.NumOut()-1; i++ {
		if _type.Out(i).Kind() != reflect.Ptr {
			c.values[i] = c.values[i].Elem()
		}
	}
	
	return c.values
}

func (p *fragment) setResultAttribute() {
	
	if p.statement.Name != dtd.SELECT &&
		p.statement.Name != dtd.INSERT &&
		p.statement.Name != dtd.UPDATE &&
		p.statement.Name != dtd.DELETE {
		return
	}
	
	if p.statement.HasAttribute(dtd.RESULT_TYPE) {
		p.resultAttr = fragment_result_type
		_type, slice := isSlice(p.statement.GetAttribute(dtd.RESULT_TYPE))
		kind, err := varToReflectKind(_type)
		if err != nil {
			throw(p.statement.File, p.statement.ctx, varToReflectKindErr).with(err)
		}
		p.resultType = &resultType{
			kind:  kind,
			slice: slice,
		}
	} else if p.statement.HasAttribute(dtd.RESULT_MAP) {
		p.resultAttr = fragment_result_map
	} else if p.statement.HasAttribute(dtd.RESULT) {
		p.resultAttr = fragment_result
	}
}

func (p *fragment) checkParameter(ft reflect.Type, mn, fn string) {
	if len(p.in) != ft.NumIn() {
		throw(p.statement.File, p.statement.ctx, checkParameterErr).
			format("%s expected %d bind parameter, got %d", fn, len(p.in), ft.NumIn())
	}
}

func (p *fragment) checkResult(ft reflect.Type, mn, fn string) {
	
	// ft.out[last] expected err
	
	switch p.statement.Name {
	case dtd.SELECT:
		switch p.resultAttr {
		case fragment_result_type:
			if ft.NumOut() > 1 {
				if p.resultType.slice {
					if ft.Out(0).Kind() != reflect.Slice ||
						(ft.Out(0).Elem().Kind() != reflect.Ptr && ft.Out(0).Elem().Kind() != reflect.Struct) ||
						(ft.Out(0).Elem().Kind() == reflect.Ptr && ft.Out(0).Elem().Elem().Kind() != reflect.Struct) {
						throw(p.statement.File, p.statement.ctx, checkResultErr).
							format("%s.%s out[0] expect [](*)struct, got: %s", mn, fn, ft.Out(0))
					}
				} else {
					if (ft.Out(0).Kind() != reflect.Ptr && ft.Out(0).Kind() != reflect.Struct) ||
						(ft.Out(0).Elem().Kind() == reflect.Ptr && ft.Out(0).Elem().Elem().Kind() != reflect.Struct) {
						throw(p.statement.File, p.statement.ctx, checkResultErr).
							format("%s.%s out[0] expect (*)struct, got: %s", mn, fn, ft.Out(0))
					}
				}
			}
		case fragment_result:
			for i, v := range p.out {
				if !v.expected(ft.Out(i)) {
					throw(p.statement.File, p.statement.ctx, varBindErr).
						format("%s.%s bind result '%s' expect '%s', got '%s'", mn, fn, v.name, v.Type(), ft.Out(i))
				}
			}
		}
	case dtd.INSERT, dtd.UPDATE, dtd.DELETE:
		if ft.NumOut() > 1 {
			if (ft.Out(0).Kind() != reflect.Ptr && ft.Out(0).Kind() != reflect.Int64) ||
				(ft.Out(0).Kind() == reflect.Ptr && ft.Out(0).Elem().Kind() != reflect.Int64) {
				throw(p.statement.File, p.statement.ctx, checkResultErr).
					format("%s.%s out[0] expect (*)int64, got: %s", mn, ft.Name(), ft.Out(0).Name())
			}
		}
	}
	
	return
}

func (p *fragment) parseStatement(args ...reflect.Value) (sql string, vars []interface{}, err error) {
	
	defer func() {
		e := recover()
		err = castRecoverError(p.statement.File, e)
	}()
	
	if len(p.in) != len(args) {
		throw(p.statement.File, p.statement.ctx, parasFragmentErr).format("expect %d args, got %d", len(p.in), len(args))
	}
	
	parser := newExprParser(args...)
	for i, v := range p.in {
		err = parser.baseParams.bind(v, i)
		if err != nil {
			throw(p.statement.File, p.statement.ctx, parasFragmentErr).with(err)
		}
	}
	res := new(psr)
	p.parseBlocks(parser, p.statement, res)
	if res.cacheable {
		p.cacheable = res.cacheable
		p.sql = res.sql
	}
	
	sql = res.sql
	vars = parser.vars
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
			r, err := parser.parseExpression(node.ctx, string(chars[from:i]))
			if err != nil {
				panic(err)
			}
			if inject {
				s += fmt.Sprintf("%v", r)
			} else {
				parser.varIndex++
				s += fmt.Sprintf("$%d", parser.varIndex)
				parser.addVar(r)
			}
			begin = false
			inject = false
		}
	}
	res.merge(s)
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
		switch node.Name {
		case dtd.IF:
			p.parseTest(parser, node, res)
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
	v, err := parser.parseExpression(node.ctx, node.GetAttribute(dtd.TEST))
	if err != nil {
		throw(p.statement.File, p.statement.ctx, parasFragmentErr).with(err)
	}
	b, err := cast.ToBoolE(v)
	if err != nil {
		throw(p.statement.File, p.statement.ctx, parasFragmentErr).with(err)
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
	for i, child := range node.Nodes {
		if pass {
			break
		}
		switch child.Name {
		case dtd.WHEN:
			pass = p.parseTest(parser, child, res)
		case dtd.OTHERWISE:
			if i != len(node.Nodes)-1 {
				throw(parser.file, child.ctx, parasFragmentErr).format("otherwise should be last element in choose")
			}
			p.parseBlocks(parser, child, res)
		default:
			throw(parser.file, child.ctx, parasFragmentErr).
				format("unsupported element '%s' element in choose", child.Name)
			
		}
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
			throw(p.statement.File, p.statement.ctx, parasFragmentErr).format("regexp compile error: %s", err)
		}
	}
	if strings.TrimSpace(s) != "" {
		res.merge(tag, s)
	}
}

func (p *fragment) parseSet(parser *exprParser, node *xmlNode, res *psr) {
	p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p *fragment) parseForeach(parser *exprParser, node *xmlNode, res *psr) {
	
	_var := node.GetAttribute(dtd.COLLECTION)
	collection, ok := parser.baseParams.get(_var)
	if !ok {
		throw(p.statement.File, p.statement.ctx, parasFragmentErr).
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
		item, slice = isSlice(item)
	}
	indexParam := &param{name: index, _type: reflect.Interface.String(), slice: false}
	itemParam := &param{name: item, _type: reflect.Interface.String(), slice: slice}
	
	open := node.GetAttribute(dtd.OPEN)
	_close := node.GetAttribute(dtd.CLOSE)
	separator := node.GetAttribute(dtd.SEPARATOR)
	
	parser.foreachParams = newExprParams()
	parser.paramIndex = 0
	
	elem := reflectValueElem(collection.value)
	frags := make([]string, 0)
	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < elem.Len(); i++ {
			parser.foreachParams.set(elem.Index(i), reflect.ValueOf(i))
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	case reflect.Map:
		for i, v := range elem.MapKeys() {
			parser.foreachParams.set(elem.MapIndex(v), v)
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	case reflect.Struct:
		for i := 0; i < elem.NumField(); i++ {
			parser.foreachParams.set(elem.Field(i).Elem(), elem.Field(i))
			if i == 0 {
				p.bindForeachParams(parser, indexParam, itemParam)
			}
			p.parseForeachChild(parser, node, &frags)
		}
	default:
		throw(parser.file, node.ctx, parasFragmentErr).format("foreach collection type '%s' can't range", elem.Type())
	}
	parser.foreachParams = nil
	
	if len(frags) > 0 {
		res.merge(open + strings.Join(frags, separator) + _close)
	}
}

func (p *fragment) bindForeachParams(parser *exprParser, indexParam, itemParam *param) {
	parser.foreachParams.check = map[string]int{}
	err := parser.foreachParams.bind(indexParam, 0)
	if err != nil {
		throw(p.statement.File, p.statement.ctx, varBindErr).with(err)
	}
	err = parser.foreachParams.bind(itemParam, 1)
	if err != nil {
		throw(p.statement.File, p.statement.ctx, varBindErr).with(err)
	}
}

func (p *fragment) parseParams(nodeCtx antlr.ParserRuleContext, tokens string) []*param {
	parser := newParamParser(p.statement.File)
	parser.walkMethods(initExprParser(tokens))
	return parser.params
}

func (p *fragment) parseForeachChild(parser *exprParser, node *xmlNode, frags *[]string) {
	for _, child := range node.Nodes {
		br := new(psr)
		p.parseBlock(parser, child, br)
		*frags = append(*frags, br.sql)
	}
}

type caller struct {
	fragment *fragment
	params   []reflect.Value
	values   []reflect.Value
}

func (p *caller) Scan(pointers ...interface{}) (err error) {
	for _, v := range pointers {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			err = fmt.Errorf("scan only accept pointer")
			return
		}
		p.values = append(p.values, rv)
	}
	
	return p.call()
}

func (p *caller) call() (err error) {
	
	start := time.Now()
	defer func() {
		p.fragment.logger.Debugf("[gobatis] [%s] cost: %s", p.fragment.id, time.Since(start))
	}()
	
	if len(p.fragment.in) != len(p.params) {
		err = fmt.Errorf("expected %d params, but pass %d", len(p.fragment.in), len(p.params))
		return
	}
	
	switch p.fragment.statement.Name {
	case dtd.SELECT:
		return p.query(p.params...)
	case dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		return p.exec(p.params...)
	default:
		throw(p.fragment.statement.File, p.fragment.statement.ctx, callerErr).
			format("unsupported call method '%s'", p.fragment.statement.Name)
		return
	}
}

func (p *caller) exec(in ...reflect.Value) (err error) {
	
	exec, index := p.execer(in...)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	ctx, index := p.context(in...)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	var conn *sql.Conn
	if exec == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		exec = conn
	}
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	s, vars, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	
	p.fragment.logger.Debugf("[gobatis] [%s] statement: %s", p.fragment.id, s)
	p.fragment.logger.Debugf("[gobatis] [%s] parameter: %s", p.fragment.id, p.printVars(vars))
	
	res, err := exec.ExecContext(ctx, s, vars...)
	if err != nil {
		return
	}
	
	return newExecResult(res, p.values).scan()
}

func (*caller) printVars(vars []interface{}) string {
	r := "\n"
	for i, v := range vars {
		r += fmt.Sprintf("   $%d => (%s) %+v\n", i+1, reflect.TypeOf(v), v)
	}
	return r
}

func (p *caller) query(in ...reflect.Value) (err error) {
	
	ctx, index := p.context(in...)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	q, index := p.queryer(in...)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	var conn *sql.Conn
	if q == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		q = conn
	}
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	
	s, vars, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	p.fragment.logger.Debugf("[gobatis] [%s] statement: %s", p.fragment.id, s)
	p.fragment.logger.Debugf("[gobatis] [%s] parameter: %+v", p.fragment.id, p.printVars(vars))
	
	rows, err := q.QueryContext(ctx, s, vars...)
	if err != nil {
		p.fragment.logger.Debugf("[gobatis] [%s] exec query error: %v", p.fragment.id, err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()
	
	return p.parseQueryResult(rows)
}

func (p *caller) parseQueryResult(rows *sql.Rows) error {
	
	res := queryResult{rows: rows}
	res.rows = rows
	
	err := res.setSelected(p.fragment.resultType, p.fragment.out, p.values)
	if err != nil {
		return err
	}
	
	return res.scan()
}

func (p *caller) removeParam(a []reflect.Value, i int) []reflect.Value {
	return append(a[:i], a[i+1:]...)
}

func (p *caller) context(in ...reflect.Value) (context.Context, int) {
	for i, v := range in {
		if v.Kind() == reflect.Ptr && v.Elem().Type().PkgPath() == "context" {
			return v.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
}

func (p *caller) execer(in ...reflect.Value) (execer, int) {
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

func (p *caller) queryer(in ...reflect.Value) (queryer, int) {
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
