package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/dtd"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
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

func (p *Engine) Call(name string, args ...reflect.Value) *caller {
	f, ok := p.fragmentManager.get(name)
	if !ok {
		panic(fmt.Errorf("method '%s' not exist", name))
	}
	return &caller{fragment: f, params: args}
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

func (p *Engine) BindMapper(ptr ...interface{}) (err error) {
	for _, v := range ptr {
		err = p.bindMapper(v)
		if err != nil {
			return errors.Wrap(err, "bind mapper:")
		}
	}
	return
}

func (p *Engine) bindMapper(mapper interface{}) (err error) {
	rv := reflect.ValueOf(mapper)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("exptect *struct, got: %s", rv.Type())
	}
	rv = rv.Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		if rv.Field(i).Kind() != reflect.Func {
			continue
		}
		id := rt.Field(i).Name
		m, ok := p.fragmentManager.get(id)
		if !ok {
			return fmt.Errorf("%s.%s not defined", rt.Name(), id)
		}
		ft := rv.Field(i).Type()
		if ft.NumOut() == 0 || !isErrorType(ft.Out(ft.NumOut()-1)) {
			return fmt.Errorf("method out expect error at last")
		}
		err = m.checkParameter(rt, ft)
		if err != nil {
			return err
		}
		err = m.checkResult(rt, ft)
		if err != nil {
			return
		}
		m.proxy(rv.Field(i))
	}
	return
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

func (p *Engine) makeDest(node *xmlNode) (*dest, error) {

	if node.Name != dtd.SELECT {
		return nil, nil
	}

	v := node.GetAttribute(dtd.RESULT_TYPE)

	isArray := strings.HasPrefix(v, "[]")
	if isArray {
		v = strings.TrimSpace(strings.TrimPrefix(v, "[]"))
	}

	if v == "" {
		if node.GetAttribute(dtd.RESULT) != "" {
			return nil, nil
		} else {
			return &dest{
				kind:    reflect.Struct,
				isArray: isArray,
			}, nil
		}
	}

	kind, err := toReflectKind(v)
	if err != nil {
		return nil, err
	}

	return &dest{
		kind:    kind,
		isArray: isArray,
	}, nil
}

func (p *Engine) addFragment(file string, ctx antlr.ParserRuleContext, id string, node *xmlNode) (err error) {

	_dest, err := p.makeDest(node)
	if err != nil {
		return
	}

	m, err := parseFragment(
		p.master, p.logger, file, id,
		node.GetAttribute(dtd.PARAMETER),
		node.GetAttribute(dtd.RESULT),
		_dest,
		node,
	)
	if err != nil {
		return
	}
	err = p.fragmentManager.add(m)
	if err != nil {
		return parseError(file, ctx, err.Error())
	}
	return
}
