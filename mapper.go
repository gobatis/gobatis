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

//type Params = map[string]interface{}

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

type param struct {
	name string
	kind reflect.Kind
}

type execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func newFragment(db *DB, logger Logger, id string, in, out []param, statement *xmlNode) *fragment {
	return &fragment{db: db, logger: logger, id: id, in: in, out: out, statement: statement}
}

type fragment struct {
	db        *DB
	logger    Logger
	id        string
	statement *xmlNode
	cacheable bool
	sql       string
	in        []param
	out       []param
}

func (p *fragment) call(in ...interface{}) *caller {
	return &caller{fragment: p, params: in}
}

type caller struct {
	fragment *fragment
	params   []interface{}
	values   []reflect.Value
}

func (p *caller) Scan(pointers ...interface{}) (err error) {

	start := time.Now()
	defer func() {
		p.fragment.logger.Debugf("[gobatis] [%s cost]: %s", p.fragment.id, time.Since(start))
	}()

	if len(p.fragment.out) != len(pointers) {
		err = fmt.Errorf("expected %d result fileds, but pass %d", len(p.fragment.out), len(pointers))
		return
	}

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
		err = parseError(
			p.fragment.statement.File,
			p.fragment.statement.ctx,
			fmt.Sprintf("unsupported call method '%s'", p.fragment.statement.Name),
		)
		return
	}
}

func (p *caller) exec(in ...interface{}) (err error) {

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

	s, vars, err := p.parseStatement(in)
	if err != nil {
		return
	}
	if len(vars) > 0 {
		in = vars
	}

	p.fragment.logger.Debugf("[gobatis] [%s args]: %s", p.fragment.id, in)

	res := newResult(result_result)
	res.result, err = exec.ExecContext(ctx, s, in...)
	if err != nil {
		return
	}

	return
}

func (p *caller) query(in ...interface{}) (err error) {

	q, index := p.queryer(in...)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	ctx, index := p.context(in...)
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

	s, vars, err := p.parseStatement(in)
	if err != nil {
		return
	}
	if len(vars) > 0 {
		in = vars
	}

	p.fragment.logger.Debugf("[gobatis] [%s args]: %+v", p.fragment.id, in)

	rows, err := q.QueryContext(ctx, s, in...)
	if err != nil {
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	return p.parseQueryResult(rows)
}

func (p *caller) parseQueryResult(rows *sql.Rows) (err error) {

	res := newResult(result_rows)
	res.rows = rows
	res.setSelected(p.fragment.out)
	res.setValues(p.values)
	err = res.scanAll()
	if err != nil {
		return
	}
	return
}

func (p *caller) removeParam(a []interface{}, i int) []interface{} {
	return append(a[:i], a[i+1:]...)
}

func (p *caller) context(in ...interface{}) (context.Context, int) {
	for i, v := range in {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr && rv.Elem().Type().PkgPath() == "context" {
			return rv.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
}

func (p *caller) execer(in ...interface{}) (execer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(execer)).Elem()
		for i, v := range in {
			rv := reflect.ValueOf(v)
			if rv.Type().Implements(t) {
				return rv.Interface().(execer), i
			}
		}
	}
	return nil, -1
}

func (p *caller) queryer(in ...interface{}) (queryer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(queryer)).Elem()
		for i, v := range in {
			rv := reflect.ValueOf(v)
			if rv.Type().Implements(t) {
				return rv.Interface().(queryer), i
			}
		}
	}
	return nil, -1
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

func (p *caller) parseStatement(args []interface{}) (string, []interface{}, error) {
	if p.fragment.cacheable {
		p.fragment.logger.Debugf("[gobatis] [%s cached sql]: %s", p.fragment.id, p.fragment.sql)
		return p.fragment.sql, nil, nil
	}
	parser := newExprParser(args...)
	for i, v := range p.fragment.in {
		err := parser.baseParams.alias(v.name, v.kind, i)
		if err != nil {
			return "", nil, err
		}
	}
	res := new(psr)
	err := p.parseBlocks(parser, p.fragment.statement, res)
	if err != nil {
		return "", nil, err
	}
	if res.cacheable {
		p.fragment.cacheable = res.cacheable
		p.fragment.sql = res.sql
	}

	p.fragment.logger.Debugf("[gobatis] [%s sql]: %s", p.fragment.id, res.sql)

	return res.sql, parser.vars, nil
}

func (p *caller) trimPrefixOverride(sql, prefix string) (r string, err error) {
	reg, err := regexp.Compile(`(?i)^` + prefix)
	if err != nil {
		return
	}
	r = reg.ReplaceAllString(sql, "")
	return
}

func (p *caller) parseSql(parser *exprParser, node *xmlNode, res *psr) error {
	chars := []rune(node.Text)
	begin := false
	var from int
	var next int
	var s string
	for i := 0; i < len(chars); i++ {
		if !begin {
			next = i + 1
			if chars[i] == 35 && next <= len(chars)-1 && chars[next] == 123 {
				begin = true
				i++
				from = i + 1
				continue
			} else {
				s += string(chars[i])
			}
		} else if chars[i] == 125 {
			r, err := parser.parseExpression(string(chars[from:i]))
			if err != nil {
				return err
			}
			parser.varIndex++
			s += fmt.Sprintf("$%d", parser.varIndex)
			parser.vars = append(parser.vars, r)
			begin = false
		}
	}
	res.merge(s)
	return nil
}

func (p *caller) parseBlocks(parser *exprParser, node *xmlNode, res *psr) (err error) {
	for _, child := range node.Nodes {
		err = p.parseBlock(parser, child, res)
		if err != nil {
			return
		}
	}
	return
}

func (p *caller) parseBlock(parser *exprParser, node *xmlNode, res *psr) (err error) {
	if node.textOnly {
		err = p.parseSql(parser, node, res)
	} else {
		switch node.Name {
		case dtd.IF:
			_, err = p.parseTest(parser, node, res)
		case dtd.WHERE:
			err = p.parseWhere(parser, node, res)
		case dtd.CHOOSE:
			err = p.parseChoose(parser, node, res)
		case dtd.FOREACH:
			err = p.parseForeach(parser, node, res)
		case dtd.TRIM:
			err = p.parseTrim(parser, node, res)
		case dtd.SET:
			err = p.parseSet(parser, node, res)
		}
	}
	return
}

func (p *caller) parseTest(parser *exprParser, node *xmlNode, res *psr) (bool, error) {
	v, err := parser.parseExpression(node.GetAttribute(dtd.TEST))
	if err != nil {
		return false, err
	}
	b, err := cast.ToBoolE(v)
	if err != nil {
		return false, err
	}
	if !b {
		return false, nil
	}
	return true, p.parseBlocks(parser, node, res)
}

func (p *caller) parseWhere(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, dtd.WHERE, "AND |OR ")
}

func (p *caller) parseChoose(parser *exprParser, node *xmlNode, res *psr) error {
	var pass bool
	for i, child := range node.Nodes {
		if pass {
			break
		}
		switch child.Name {
		case dtd.WHEN:
			var err error
			pass, err = p.parseTest(parser, child, res)
			if err != nil {
				return err
			}
		case dtd.OTHERWISE:
			if i != len(node.Nodes)-1 {
				return parseError(parser.file, node.ctx, "otherwise should be last element in choose")
			}
			return p.parseBlocks(parser, child, res)
		default:
			return parseError(parser.file, child.ctx, fmt.Sprintf("unsupported element '%s' element in choose", child.Name))
		}
	}
	return nil
}

func (p *caller) parseTrim(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p *caller) trimPrefixOverrides(parser *exprParser, node *xmlNode, res *psr, tag, prefixes string) error {
	wr := new(psr)
	err := p.parseBlocks(parser, node, wr)
	if err != nil {
		return err
	}
	s := strings.TrimSpace(wr.sql)
	filters := strings.Split(prefixes, "|")
	for _, v := range filters {
		s, err = p.trimPrefixOverride(s, v)
		if err != nil {
			err = parseError(parser.file, node.ctx, fmt.Sprintf("regexp compile error: %s", err))
			return err
		}
	}
	if strings.TrimSpace(s) != "" {
		res.merge(tag, s)
	}
	return nil
}

func (p *caller) parseSet(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p *caller) parseForeach(parser *exprParser, node *xmlNode, res *psr) error {

	_var := node.GetAttribute(dtd.COLLECTION)
	collection, ok := parser.baseParams.get(_var)
	if !ok {
		return parseError(parser.file, node.ctx, fmt.Sprintf("can't get foreach collection '%s' value", _var))
	}
	index := node.GetAttribute(dtd.INDEX)
	if index == "" {
		index = dtd.INDEX
	}
	item := node.GetAttribute(dtd.ITEM)
	if item == "" {
		item = dtd.ITEM
	}
	subParams := fmt.Sprintf("%s,%s", index, item)
	open := node.GetAttribute(dtd.OPEN)
	_close := node.GetAttribute(dtd.CLOSE)
	separator := node.GetAttribute(dtd.SEPARATOR)

	parser.foreachParams = newExprParams()
	parser.paramIndex = 0

	elem := realReflectElem(collection.value)
	frags := make([]string, 0)
	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < elem.Len(); i++ {
			parser.foreachParams.set(i, elem.Index(i).Interface())
			if i == 0 {
				err := parser.parseParameter(subParams)
				if err != nil {
					return err
				}
			}
			err := p.parseForeachChild(parser, node, &frags)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		for i, v := range elem.MapKeys() {
			parser.foreachParams.set(v.Interface(), elem.MapIndex(v).Interface())
			if i == 0 {
				err := parser.parseParameter(subParams)
				if err != nil {
					return err
				}
			}
			err := p.parseForeachChild(parser, node, &frags)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		for i := 0; i < elem.NumField(); i++ {
			parser.foreachParams.set(elem.Field(i).Interface(), elem.Field(i).Elem().Interface())
			if i == 0 {
				err := parser.parseParameter(subParams)
				if err != nil {
					return err
				}
			}
			err := p.parseForeachChild(parser, node, &frags)
			if err != nil {
				return err
			}
		}
	default:
		return parseError(parser.file, node.ctx,
			fmt.Sprintf("foreach collection type '%s' can't range", elem.Kind()))
	}
	parser.foreachParams = nil

	if len(frags) > 0 {
		res.merge(open + strings.Join(frags, separator) + _close)
	}

	return nil
}

func (p *caller) parseForeachChild(parser *exprParser, node *xmlNode, frags *[]string) error {
	for _, child := range node.Nodes {
		br := new(psr)
		err := p.parseBlock(parser, child, br)
		if err != nil {
			return err
		}
		*frags = append(*frags, br.sql)
	}
	return nil
}
