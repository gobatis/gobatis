package gobatis

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

const config_file = "gobatis.xml"

func NewEngine(bundle Bundle) (engine *Engine, err error) {
	engine = &Engine{
		bundle: bundle,
	}
	err = engine.init()
	return
}

type Engine struct {
	bundle Bundle
	db     sql.DB
}

func (p *Engine) init() (err error) {
	err = p.parseConfig()
	if err != nil {
		return
	}
	return
}

func (p *Engine) parseConfig() (err error) {
	if p.bundle == nil {
		err = fmt.Errorf("no resource bundle")
		return
	}
	c, err := p.bundle.Open(config_file)
	if err != nil {
		err = fmt.Errorf("open %s error: %s", config_file, err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	d, err := ioutil.ReadAll(c)
	if err != nil {
		err = fmt.Errorf("read %s content error: %s", config_file, err)
		return
	}
	return parseConfig(p, config_file, d)
}

func (p *Engine) initDB() {

}

func (p *Engine) parseSql() (err error) {
	return
}

func (p *Engine) BindMapper(mapper ...interface{}) {
	//tx := p.db.Begin()
	//tx.Exec()
	//tx.Query()
	//tx.QueryRow()
	//p.db.QueryRow()
}

func (p *Engine) GetSQL(name string) string {
	return "hello world!"
}
