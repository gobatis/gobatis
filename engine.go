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
	"sync"
)

func NewPostgresql(dsn string) *Engine {
	return NewEngine(NewDB(postgresql.PGX, dsn))
}

func NewMySQL(dsn string) *Engine {
	return NewEngine(NewDB(mysql.MySQL, dsn))
}

func NewEngine(db *DB) *Engine {
	engine := &Engine{DB: db}
	return engine
}

type Engine struct {
	*DB
	mu         sync.RWMutex
	bundle     http.FileSystem
	slaves     []*DB
	logger     Logger
	statements map[string]*xmlNode
	sqlCaches  map[string]string
}

func (p *Engine) SetBundle(bundle http.FileSystem) {
	p.bundle = bundle
}

func (p *Engine) Init() (err error) {
	err = p.initDB()
	if err != nil {
		err = fmt.Errorf("init master db error: %s", err)
		return
	}
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

func (p *Engine) Call(name string, args ...interface{}) (res interface{}, err error) {
	var ok bool
	//sql, ok := p.getSqlCache(name)
	//if !ok {
	var node *xmlNode
	node, ok = p.getStatement(name)
	if !ok {
		err = fmt.Errorf("not found statement: %s", name)
		return
	}
	var r *psr
	r, err = p.parseStatement(node, args...)
	if err != nil {
		return
	}
	//if r.cache {
	//	p.addSqlCache(r.id, r.sql)
	//}
	//sql = r.sql
	//}
	
	fmt.Println("sql:", r.id, r.sql)
	
	return
}

func (p *Engine) BindMapper(mapper ...interface{}) {
	//tx := p.db.Begin()
	//tx.Exec()
	//tx.Query()
	//tx.QueryRow()
	//p.db.QueryRow()
}

func (p *Engine) SQL(name string) string {
	return "hello world!"
}

func (p *Engine) initLogger() {
	if p.logger == nil {
		p.logger = newLogger()
	}
}

func (p *Engine) parseConfig() (err error) {
	if p.bundle == nil {
		err = fmt.Errorf("no set bundle")
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

func (p *Engine) addStatement(file string, ctx antlr.ParserRuleContext, token antlr.Token, id string, node *xmlNode) (err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.statements == nil {
		p.statements = map[string]*xmlNode{}
	}
	_, ok := p.statements[id]
	if ok {
		err = parseError(file, ctx, fmt.Sprintf("duplicate statement: %s", id))
		return
	}
	p.statements[id] = node
	return
}

func (p *Engine) getStatement(id string) (node *xmlNode, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.statements == nil {
		return
	}
	node, ok = p.statements[id]
	return
}

// mapper file duplicate cache tag has filtered
//func (p *Engine) addCache(file string, node *xmlNode) (err error) {
//	p.mu.Lock()
//	defer p.mu.Unlock()
//	if p.caches == nil {
//		p.caches = map[string]*xmlNode{}
//	}
//	p.caches[file] = node
//	return
//}
//
//func (p *Engine) getCache(file string) (node *xmlNode, ok bool) {
//	p.mu.RLock()
//	defer p.mu.RUnlock()
//	if p.caches == nil {
//		return
//	}
//	node, ok = p.caches[file]
//	return
//}

func (p *Engine) addSqlCache(id, sql string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.sqlCaches == nil {
		p.sqlCaches = map[string]string{}
	}
	p.sqlCaches[id] = sql
	return
}

func (p *Engine) getSqlCache(id string) (sql string, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.sqlCaches == nil {
		return
	}
	sql, ok = p.sqlCaches[id]
	return
}

type psr struct {
	sql   string
	cache bool
	id    string
}

func (p *psr) appendSql(s ...string) {
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

//func (p *Engine) makeParams(args ...interface{}) Params {
//	if len(args) == 0 {
//		return nil
//	}
//	baseParams := make(Params)
//	for _, v := range args {
//		//rv :=
//		fmt.Println("make baseParams 参数:", v)
//	}
//	return baseParams
//}

func (p *Engine) parseStatement(node *xmlNode, params ...interface{}) (res *psr, err error) {
	if node == nil {
		err = fmt.Errorf("parse node is nil")
		return
	}
	res = new(psr)
	//params := p.makeParams(args...)
	parser := newExprParser(params...)
	err = parser.parseParameter(node.GetAttribute(dtd.PARAMETER_TYPE))
	if err != nil {
		return
	}
	err = p.parseBlock(parser, node, params, res)
	if err != nil {
		return
	}
	
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

func (p *Engine) parseBlock(parser *exprParser, node *xmlNode, params []interface{}, res *psr) (err error) {
	for _, child := range node.Nodes {
		if child.textOnly {
			res.appendSql(child.Text)
		} else {
			switch child.Name {
			case dtd.IF:
				err = p.parseTest(parser, child, params, res)
			case dtd.WHERE:
				err = p.parseWhere(parser, child, params, res)
			case dtd.CHOOSE:
				err = p.parseChoose(parser, child, params, res)
			case dtd.FOREACH:
				err = p.parseForeach(parser, child, params, res)
			case dtd.TRIM:
				err = p.parseTrim(parser, child, params, res)
			case dtd.SET:
				err = p.parseSet(parser, child, params, res)
			}
		}
	}
	return
}

func (p *Engine) renderSql(parser *exprParser, sql string) (result string, err error) {
	return
}

func (p *Engine) parseTest(parser *exprParser, node *xmlNode, params []interface{}, res *psr) error {
	v, err := parser.parseExpression(node.GetAttribute(dtd.TEST))
	if err != nil {
		return err
	}
	b, err := cast.ToBoolE(v)
	if err != nil {
		return err
	}
	if !b {
		return nil
	}
	return p.parseBlock(parser, node, params, res)
}

func (p *Engine) parseWhere(parser *exprParser, node *xmlNode, params []interface{}, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, params, res, dtd.WHERE, "AND |OR ")
}

func (p *Engine) parseChoose(parser *exprParser, node *xmlNode, params []interface{}, res *psr) (err error) {
	for i, child := range node.Nodes {
		switch child.Name {
		case dtd.WHEN:
			err = p.parseTest(parser, node, params, res)
		case dtd.OTHERWISE:
			if i != len(node.Nodes) {
				err = parseError(parser.file, node.ctx, "otherwise should be last element in choose")
				return
			}
			err = p.parseBlock(parser, node, params, res)
		}
	}
	return
}

func (p *Engine) parseTrim(parser *exprParser, node *xmlNode, params []interface{}, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, params, res, node.GetAttribute(dtd.PREFIX), node.GetAttribute(dtd.PREFIX_OVERRIDES))
}

func (p *Engine) trimPrefixOverrides(parser *exprParser, node *xmlNode, params []interface{}, res *psr, tag, prefixes string) (err error) {
	wr := new(psr)
	for _, child := range node.Nodes {
		if child.textOnly {
			res.appendSql(child.Text)
		} else {
			err = p.parseBlock(parser, child, params, wr)
			if err != nil {
				return
			}
		}
	}
	sql := strings.TrimSpace(wr.sql)
	filters := strings.Split(prefixes, "|")
	for _, v := range filters {
		sql, err = p.trimPrefixOverride(sql, v)
		if err != nil {
			err = parseError(parser.file, node.ctx, fmt.Sprintf("regexp compile error: %s", err))
			return
		}
	}
	res.appendSql(tag, sql)
	return
}

func (p *Engine) parseSet(parser *exprParser, node *xmlNode, params []interface{}, res *psr) (err error) {
	return p.trimPrefixOverrides(parser, node, params, res, dtd.SET, ",")
}

func (p *Engine) parseForeach(parser *exprParser, node *xmlNode, params []interface{}, res *psr) error {
	
	collection, ok := parser.baseParams.get(node.GetAttribute(dtd.COLLECTION))
	if !ok {
		return parseError(parser.file, node.ctx, "can't get foreach collection value")
	}
	
	subParams := fmt.Sprintf("%s,%s", node.GetAttribute(dtd.INDEX), node.GetAttribute(dtd.ITEM))
	open := node.GetAttribute(dtd.OPEN)
	_close := node.GetAttribute(dtd.CLOSE)
	separator := node.GetAttribute(dtd.SEPARATOR)
	parser.foreachParams = newExprParams()
	elem := collection.reflectElem()
	frags := make([]string, 0)
	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < elem.Len()-1; i++ {
			parser.foreachParams.set(i, elem.Index(i).Interface())
			err := parser.parseParameter(subParams)
			if err != nil {
				return err
			}
			err = p.parseForeachChild(parser, node, params, &frags, separator)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, v := range elem.MapKeys() {
			parser.foreachParams.set(v.Interface(), elem.MapIndex(v).Interface())
			err := parser.parseParameter(subParams)
			if err != nil {
				return err
			}
			err = p.parseForeachChild(parser, node, params, &frags, separator)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		for i := 0; i < elem.NumField()-1; i++ {
			parser.foreachParams.set(elem.Field(i).Interface(), elem.Field(i).Elem().Interface())
			err := parser.parseParameter(subParams)
			if err != nil {
				return err
			}
			err = p.parseForeachChild(parser, node, params, &frags, separator)
			if err != nil {
				return err
			}
		}
	default:
		return parseError(parser.file, node.ctx, "foreach collection type can't range")
	}
	parser.foreachParams = nil
	
	if len(frags) > 0 {
		res.appendSql(open, strings.Join(frags, ""), _close)
	}
	
	return nil
}

func (p *Engine) parseForeachChild(parser *exprParser, node *xmlNode, params []interface{}, frags *[]string, separator string) error {
	for _, child := range node.Nodes {
		br := new(psr)
		err := p.parseBlock(parser, child, params, br)
		if err != nil {
			return err
		}
		br.appendSql(separator)
		*frags = append(*frags, br.sql)
	}
	return nil
}
