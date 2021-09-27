package gobatis

type inserter struct {
	table string
	fl    []string
	fm    map[string]bool
	vs    [][]string
	rows  bool
	item  string
	index string
}

func (p *inserter) addField(v string) {
	if p.fm == nil {
		p.fm = map[string]bool{}
	}
	p.fm[v] = true
	p.fl = append(p.fl, v)
}

func (p *inserter) hasField(v string) bool {
	if p.fm == nil {
		return false
	}
	_, ok := p.fm[v]
	return ok
}

func (p *inserter) noField() bool {
	return p.fm == nil
}

func (p *inserter) noValue() bool {
	return len(p.vs) == 0
}
