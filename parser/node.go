package parser

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

type Node struct {
	Name       string
	attributes map[string]string
	nodes      []*Node
	countMap   map[string]int
}

func (p *Node) addAttribute(name, value string) {
	if p.attributes == nil {
		p.attributes = map[string]string{}
	}
	p.attributes[name] = value
}

func (p *Node) hasAttribute(name string) bool {
	if p.attributes == nil {
		return false
	}
	_, ok := p.attributes[name]
	return ok
}

func (p *Node) addNode(node *Node) {
	if p.countMap == nil {
		p.countMap = map[string]int{}
	}
	p.countMap[node.Name]++
	p.nodes = append(p.nodes, node)
}

func (p *Node) count(name string) int {
	if p.countMap == nil {
		return 0
	}
	return p.countMap[name]
}
