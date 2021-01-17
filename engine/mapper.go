package engine

// 绑定Mapper实例
func (p *Engine) BindMapper(mapper interface{}) (err error) {

	return
}

func NewMapper(path string) *Mapper {
	return &Mapper{Path: path}
}

type Mapper struct {
	Path    string
	methods *Methods
}

func (p *Mapper) Methods() *Methods {
	if p.methods == nil {
		p.methods = NewMethods()
	}
	return p.methods
}
