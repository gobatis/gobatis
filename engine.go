package gobatis

type Engine struct {
}

func (p *Engine) InitDB(driver string) {

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
