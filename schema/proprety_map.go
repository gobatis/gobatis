package schema

type PropertyMap struct {
	list []*Property
	_map map[string]*Property
}

func NewPropertyMap() *PropertyMap {
	return &PropertyMap{}
}

// 添加元素，如有重复则覆盖旧元素
func (p *PropertyMap) Add(val ...*Property) {
	for _, v := range val {
		if p._map == nil {
			p._map = map[string]*Property{}
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

func (p *PropertyMap) Get(name string) *Property {
	if p._map == nil {
		return nil
	}
	v, ok := p._map[name]
	if ok {
		return v
	}
	return nil
}

func (p *PropertyMap) Has(name string) bool {
	if p._map == nil {
		return false
	}
	_, ok := p._map[name]
	if ok {
		return true
	}
	return false
}
