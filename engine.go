package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/koyeo/gobatis/driver/mysql"
	"github.com/koyeo/gobatis/driver/postgresql"
	"github.com/koyeo/gobatis/dtd"
	"io/ioutil"
	"net/http"
	"path/filepath"
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
	sql, ok := p.getSqlCache(name)
	if !ok {
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
		if r.cache {
			p.addSqlCache(r.id, r.sql)
		}
		sql = r.sql
	}
	
	fmt.Println(sql)
	
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

func (p *Engine) addStatement(file string, token antlr.Token, id string, node *xmlNode) (err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.statements == nil {
		p.statements = map[string]*xmlNode{}
	}
	_, ok := p.statements[id]
	if ok {
		err = parseError(file, token, fmt.Sprintf("duplicate statement: %s", id))
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

func (p *psr) appendSql(s string) {
	if p.sql != "" && !strings.HasSuffix(p.sql, " ") {
		p.sql += " " + s
	} else {
		p.sql += s
	}
}

func (p *Engine) makeParams(args ...interface{}) Params {
	if len(args) == 0 {
		return nil
	}
	params := make(Params)
	for _, v := range args {
		//rv :=
		fmt.Println(v)
	}
	return params
}

func (p *Engine) parseStatement(node *xmlNode, args ...interface{}) (res *psr, err error) {
	if node == nil {
		err = fmt.Errorf("parse node is nil")
		return
	}
	res = new(psr)
	params := p.makeParams(args...)
	for _, v := range node.Nodes {
		if v.textOnly {
			res.appendSql(v.Text)
		} else {
			switch v.Name {
			case dtd.IF:
				err = p.parseIf(v, params, res)
			case dtd.WHERE:
				err = p.parseWhere(v, params, res)
			}
		}
	}
	return
}

func (p *Engine) parseIf(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseWhere(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseChoose(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseWhen(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseOtherwise(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseTrim(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseSet(node *xmlNode, params Params, res *psr) (err error) {
	return
}

func (p *Engine) parseForeach(node *xmlNode, params Params, res *psr) (err error) {
	return
}
