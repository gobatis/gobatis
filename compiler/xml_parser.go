package compiler

import (
	"strings"
)

const (
	ST_TEXT = "text"
	ST_NODE = "node"
)

type XMLParser struct {
	tokens []*Token
	peek   *Token
	body   []*XMLNode
	nodes  []*XMLNode
	state  int
	index  int
}

func NewXMLParser(tokens []*Token) *XMLParser {
	return &XMLParser{tokens: tokens, index: -1}
}

func (p *XMLParser) next() {
	p.index += 1
	if p.index < len(p.tokens) {
		p.peek = p.tokens[p.index]
	} else {
		p.peek = nil
	}
}

func (p *XMLParser) Parse() []*XMLNode {
	p.next()
	for p.peek != nil {
		p.parse()
		p.next()
	}
	return p.body
}

func (p *XMLParser) parse() {

	switch p.peek.Type {
	case TT_TEXT:
		p.addNode(&XMLNode{
			Type:  ST_TEXT,
			Value: p.peek.Value,
		})
	case TT_START_TAG:
		p.addNode(&XMLNode{
			Type: ST_NODE,
			Name: p.peek.Value,
		})
	case TT_ATTR_NAME:
		attr := &XMLAttribute{
			Name: strings.ToLower(p.peek.Value),
		}
		p.next()
		if p.peek.Type != TT_ATTR_VALUE {
			// TODO 报错，属性未赋值
		}
		attr.Value = p.peek.Value
		p.setAttribute(p.lastUnclosedNode(), attr)
	case TT_SELF_END_TAG:
		p.lastUnclosedNode().closed = true
		p.nodes = p.nodes[:len(p.nodes)-1]
	case TT_END_TAG:
		if p.lastUnclosedNode().Name != strings.ToLower(p.peek.Value) {
			// TODO 标签不匹配
		}
		p.lastUnclosedNode().closed = true
		p.nodes = p.nodes[:len(p.nodes)-1]
	default:
		// TODO 报错，未期待 Token
	}

}

func (p *XMLParser) setAttribute(node *XMLNode, attr *XMLAttribute) {
	if node.attributeMap == nil {
		node.attributeMap = map[string]string{}

	}
	_, ok := node.attributeMap[attr.Name]
	if ok {
		// TODO 报错，属性重复
	}

	node.attributeMap[attr.Name] = attr.Value
	node.Attributes = append(node.Attributes, attr)
}

func (p *XMLParser) addNode(node *XMLNode) {
	node.Start = p.peek.Start
	node.End = p.peek.End
	lastNode := p.lastUnclosedNode()
	if lastNode != nil && !lastNode.closed {
		lastNode.AppendBody(node)
	} else {
		p.body = append(p.body, node)
	}
	if node.Type == ST_NODE {
		p.nodes = append(p.nodes, node)
	}
}

func (p *XMLParser) lastUnclosedNode() *XMLNode {
	// TODO 跑跑看，是否需要强制获取
	l := len(p.nodes)
	if l > 0 {
		return p.nodes[l-1]
	}
	return nil
}
