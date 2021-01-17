package engine

type Mappers struct {
	list []*Mapper
	_map map[string]*Mapper
}

func NewMappers() *Mappers {
	return &Mappers{}
}

// 添加元素，如有重复则覆盖旧元素
func (p *Mappers) Add(val ...*Mapper) {
	for _, v := range val {
		if p._map == nil {
			p._map = map[string]*Mapper{}
		}
		if p.Has(v.Path) {
			for i, vv := range p.list {
				if vv.Path == v.Path {
					p.list[i] = v
					break
				}
			}
			p._map[v.Path] = v
		} else {
			p.list = append(p.list, v)
			p._map[v.Path] = v
		}
	}
}

func (p *Mappers) Get(name string) *Mapper {
	if p._map == nil {
		return nil
	}
	v, ok := p._map[name]
	if ok {
		return v
	}
	return nil
}

func (p *Mappers) Has(name string) bool {
	if p._map == nil {
		return false
	}
	_, ok := p._map[name]
	if ok {
		return true
	}
	return false
}
