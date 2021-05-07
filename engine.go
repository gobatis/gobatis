package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/dtd"
	"io/ioutil"
	"net/http"
	"path/filepath"
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

func (p *Engine) Call(name string, args ...interface{}) *caller {
	f, ok := p.fragmentManager.get(name)
	if !ok {
		panic(fmt.Errorf("method '%s' not exist", name))
	}
	return f.call(args...)
}

func (p *Engine) Master() *DB {
	return p.master
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
	m, err := parseFragment(
		p.master, p.logger, file, id,
		node.GetAttribute(dtd.PARAMETER_TYPE), node.GetAttribute(dtd.RESULT_TYPE), node)
	if err != nil {
		return
	}
	err = p.fragmentManager.add(m)
	if err != nil {
		return parseError(file, ctx, err.Error())
	}
	return
}
