package dtd

func NewElement(name string) Element {
	return Element{Name: name}
}

type Element struct {
	Name       string
	Nodes      map[string]int
	Attributes map[string]int
}

func (p *Element) AddNode(node string, times int) {
	if p.Nodes == nil {
		p.Nodes = map[string]int{}
	}
	
	p.Nodes[node] = times
	
}

func (p *Element) AddAttribute(attribute string, check int) {
	if p.Attributes == nil {
		p.Attributes = map[string]int{}
	}
	
	p.Attributes[attribute] = check
	
}

func (p Element) HasNode(name string) bool {
	if p.Nodes == nil {
		return false
	}
	_, ok := p.Nodes[name]
	return ok
}

func (p Element) GetNodeCount(name string) int {
	if p.Nodes == nil {
		return 0
	}
	return p.Nodes[name]
}

func (p Element) HasAttribute(name string) bool {
	if p.Attributes == nil {
		return false
	}
	_, ok := p.Attributes[name]
	return ok
}
