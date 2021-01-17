package engine

type Methods struct {
	list []*Method
	_map map[string]*Method
}

func NewMethods() *Methods {
	return &Methods{}
}

// 添加元素，如有重复则覆盖旧元素
func (p *Methods) Add(val ...*Method) {
	for _, v := range val {
		if p._map == nil {
			p._map = map[string]*Method{}
		}
		if p.Has(v.Name) {
			for i, vv := range p.list {
				if vv.Name == v.Name {
					p.list[i] = v
					break
				}
			}
			p._map[v.Name] = v
		} else {
			p.list = append(p.list, v)
			p._map[v.Name] = v
		}
	}
}

func (p *Methods) Get(name string) *Method {
	if p._map == nil {
		return nil
	}
	v, ok := p._map[name]
	if ok {
		return v
	}
	return nil
}

func (p *Methods) Has(name string) bool {
	if p._map == nil {
		return false
	}
	_, ok := p._map[name]
	if ok {
		return true
	}
	return false
}
