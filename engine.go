package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/dtd"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

func NewPostgresql(dsn string) *Engine {
	return NewEngine(NewDB(postgresql.PGX, dsn))
}

func NewMySQL(dsn string) *Engine {
	return NewEngine(NewDB(mysql.MySQL, dsn))
}

func NewEngine(db *DB) *Engine {
	engine := &Engine{master: db, logger: newLogger(), fragmentManager: newMethodManager()}
	return engine
}

type Engine struct {
	master          *DB
	bundle          http.FileSystem
	slaves          []*DB
	logger          Logger
	fragmentManager *fragmentManager
}

func (p *Engine) SetBundle(bundle http.FileSystem) {
	p.bundle = bundle
}

func (p *Engine) Init() (err error) {
	err = p.master.initDB()
	if err != nil {
		err = fmt.Errorf("init master db error: %s", err)
		return
	}
	err = p.parseBundle()
	if err != nil {
		return
	}
	return
}

func (p *Engine) parseBundle() (err error) {
	err = p.parseConfig()
	if err != nil {
		return
	}

	err = p.parseMappers()
	if err != nil {
		return
	}
	return
}

//func (p *Engine) Call(name string, args ...interface{}) (res []interface{}, err error) {
//	var ok bool
//	//sql, ok := p.getSqlCache(name)
//	//if !ok {
//	var node *xmlNode
//	node, ok = p.getStatement(name)
//	if !ok {
//		err = fmt.Errorf("not found statement: %s", name)
//		return
//	}
//	var r *psr
//	var in []interface{}
//	r, in, err = p.parseStatement(node, args...)
//	if err != nil {
//		return
//	}
//	p.logger.Debugf("[%s] query: %s", name, r.sql)
//	p.logger.Debugf("[%s]  args: %+v", name, in)
//	stmt, err := p.Prepare(r.sql)
//	if err != nil {
//		return
//	}
//	defer func() {
//		_ = stmt.Close()
//	}()
//	_, err = stmt.Exec(in...)
//	if err != nil {
//		return
//	}
//	stmt.QueryRow()
//	//conn, err := p.DB.Conn(context.Background())
//	//conn.B
//	//tx ,_:= p.Begin()
//	//tx.PrepareContext()
//	return
//}

func (p *Engine) BindMapper(mapper ...interface{}) {
	//tx := p.db.Begin()
	//tx.Exec()
	//tx.Query()
	//tx.QueryRow()
	//p.db.QueryRow()
}

func (p *Engine) BindVar(pointer interface{}, result map[string]interface{}) error {
	return nil
}

func (p *Engine) makeMapper() {

}

func (p *Engine) initLogger() {
	if p.logger == nil {
		p.logger = newLogger()
	}
}

func (p *Engine) parseConfig() (err error) {
	if p.bundle == nil {
		err = fmt.Errorf("no bundle")
		return
	}
	d, err := p.readBundleFile(CONFIG_XML)
	if err != nil {
		return
	}
	err = parseConfig(p, CONFIG_XML, string(d))
	return
}

func (p *Engine) readBundleFile(path string) (d []byte, err error) {
	c, err := p.bundle.Open(path)
	if err != nil {
		err = fmt.Errorf("open %s error: %s", path, err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	d, err = ioutil.ReadAll(c)
	if err != nil {
		err = fmt.Errorf("read %s content error: %s", path, err)
		return
	}
	return
}

func (p *Engine) parseMappers() (err error) {
	files, err := p.walkMappers("/")
	if err != nil {
		return
	}
	for _, v := range files {
		var d []byte
		d, err = p.readBundleFile(v)
		if err != nil {
			return
		}
		err = parseMapper(p, v, string(d))
	}
	return
}

func (p *Engine) walkMappers(root string) (files []string, err error) {
	handle, err := p.bundle.Open(root)
	if err != nil {
		return
	}
	defer func() {
		_ = handle.Close()
	}()
	res, err := handle.Readdir(-1)
	for _, v := range res {
		path := filepath.Join(root, v.Name())
		if v.IsDir() {
			var _files []string
			_files, err = p.walkMappers(path)
			if err != nil {
				return
			}
			files = append(files, _files...)
		} else {
			if path != filepath.Join("/", CONFIG_XML) {
				files = append(files, path)
			}
		}
	}
	return
}

func (p *Engine) addFragment(file string, ctx antlr.ParserRuleContext, id string, node *xmlNode) (err error) {
	m, err := parseFragment(file, id, node.GetAttribute(dtd.PARAMETER_TYPE), node.GetAttribute(dtd.RESULT_TYPE), node)
	if err != nil {
		return
	}
	err = p.fragmentManager.add(m)
	if err != nil {
		return parseError(file, ctx, err.Error())
	}
	return
}

type psr struct {
	sql   string
	cache bool
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

func (p *Engine) parseStatement(node *xmlNode, params ...interface{}) (res *psr, in []interface{}, err error) {
	if node == nil {
		err = fmt.Errorf("parse node is nil")
		return
	}
	res = new(psr)
	parser := newExprParser(params...)
	err = parser.parseParameter(node.GetAttribute(dtd.PARAMETER_TYPE))
	if err != nil {
		return
	}
	err = p.parseBlocks(parser, node, res)
	if err != nil {
		return
	}
	in = parser.vars
	return
}

func (p *Engine) trimPrefixOverride(sql, prefix string) (r string, err error) {
	reg, err := regexp.Compile(`(?i)^` + prefix)
	if err != nil {
		return
	}
	r = reg.ReplaceAllString(sql, "")
	return
}

func (p *Engine) parseSql(parser *exprParser, node *xmlNode, res *psr) error {
	chars := []rune(node.Text)
	begin := false
	var from int
	var next int
	var sql string
	for i := 0; i < len(chars); i++ {
		if !begin {
			next = i + 1
			if chars[i] == 35 && next <= len(chars)-1 && chars[next] == 123 {
				begin = true
				i++
				from = i + 1
				continue
			} else {
				sql += string(chars[i])
			}
		} else if chars[i] == 125 {
			r, err := parser.parseExpression(string(chars[from:i]))
			if err != nil {
				return err
			}
			i++
			parser.varIndex++
			sql += fmt.Sprintf("$%d", parser.varIndex)
			parser.vars = append(parser.vars, r)
			begin = false
		}
	}
	res.merge(sql)
	return nil
}

func (p *Engine) parseBlocks(parser *exprParser, node *xmlNode, res *psr) (err error) {
	for _, child := range node.Nodes {
		err = p.parseBlock(parser, child, res)
		if err != nil {
			return
		}
	}
	return
}

func (p *Engine) parseBlock(parser *exprParser, node *xmlNode, res *psr) (err error) {
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

func (p *Engine) parseTest(parser *exprParser, node *xmlNode, res *psr) (bool, error) {
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

func (p *Engine) parseWhere(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, dtd.WHERE, "AND |OR ")
}

func (p *Engine) parseChoose(parser *exprParser, node *xmlNode, res *psr) error {
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

func (p *Engine) parseTrim(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p *Engine) trimPrefixOverrides(parser *exprParser, node *xmlNode, res *psr, tag, prefixes string) error {
	wr := new(psr)
	err := p.parseBlocks(parser, node, wr)
	if err != nil {
		return err
	}
	sql := strings.TrimSpace(wr.sql)
	filters := strings.Split(prefixes, "|")
	for _, v := range filters {
		sql, err = p.trimPrefixOverride(sql, v)
		if err != nil {
			err = parseError(parser.file, node.ctx, fmt.Sprintf("regexp compile error: %s", err))
			return err
		}
	}
	if strings.TrimSpace(sql) != "" {
		res.merge(tag, sql)
	}
	return nil
}

func (p *Engine) parseSet(parser *exprParser, node *xmlNode, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, res, dtd.SET, ",")
}

func (p *Engine) parseForeach(parser *exprParser, node *xmlNode, res *psr) error {

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

func (p *Engine) parseForeachChild(parser *exprParser, node *xmlNode, frags *[]string) error {
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
