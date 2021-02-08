package compiler

import (
	"fmt"
	"strings"
)

const (
	ST_TEXT = "text"
	ST_NODE = "node"
)

type XMLNode struct {
	Start        *Point          `json:"-"`
	End          *Point          `json:"-"`
	Type         string          `json:"type"`
	Name         string          `json:"name,omitempty"`
	Value        string          `json:"value,omitempty"`
	RAW          string          `json:"value,omitempty"`
	Attributes   []*XMLAttribute `json:"attributes,omitempty"`
	Body         []*XMLNode      `json:"body,omitempty"`   // tree body
	Tokens       []*Token        `json:"tokens,omitempty"` // SQL tokens
	attributeMap map[string]string
	closed       bool
}

func (p *XMLNode) AppendBody(node ...*XMLNode) {
	p.Body = append(p.Body, node...)
}

type XMLAttribute struct {
	Start *Point `json:"-"`
	End   *Point `json:"-"`
	Name  string `json:"name"`
	Value string `json:"value"`
}


func NewXMLTokenizer(content []byte) *XMLTokenizer {
	return &XMLTokenizer{
		pos: &Position{
			index:  -1,
			line:   1,
			column: 0,
		},
		chars: []rune(string(content)),
		state: TS_LITERAL,
	}
}

type XMLTokenizer struct {
	file   string
	pos    *Position
	start  *Position
	chars  []rune
	peek   rune
	state  int
	tokens []*Token
}

func (p *XMLTokenizer) next() {
	p.pos.next(p.peek)
	if p.pos.index < len(p.chars) {
		p.peek = p.chars[p.pos.index]
	} else {
		p.peek = EOF
	}
}

func (p *XMLTokenizer) Parse() []*Token {
	p.next()
	p.start = p.pos.fork()
	for p.peek != EOF {
		switch p.state {
		case TS_LITERAL:
			p.parseLiteral()
		case TS_START_TAG:
			p.parseTagStart()
		case TS_START_TAG_NAME:
			p.parseStartTagName()
		case TS_END_TAG_NAME:
			p.parseEndTagName()
		case TS_ATTRIBUTE:
			p.parseAttribute()
		case TS_ATTRIBUTE_NAME:
			p.parseAttributeName()
		case TS_ATTRIBUTE_EQUAL:
			p.parseAttributeEqual()
		case TS_ATTRIBUTE_VALUE_START:
			p.parseAttributeValueStart()
		case TS_ATTRIBUTE_VALUE_END:
			p.parseAttributeValueEnd()
		case TS_SELF_END_TAG:
			p.parseSelfEndTag()
		}
		p.next()
	}
	switch p.state {
	case TS_LITERAL:
		p.addToken(TT_TEXT)
	}
	return p.tokens
}

func (p *XMLTokenizer) parseLiteral() {

	if p.peek == LESS_THAN {
		p.addToken(TT_TEXT)
		p.expectStatus(TS_START_TAG)
	}
}

func (p *XMLTokenizer) parseTagStart() {
	if IsLetter(p.peek) {
		p.expectStatus(TS_START_TAG_NAME)
	} else if p.peek == FORWARD_SLASH {
		p.next()
		p.expectStatus(TS_END_TAG_NAME)
	}
}

func (p *XMLTokenizer) parseStartTagName() {

	if IsBlank(p.peek) {
		// <m
		p.addToken(TT_START_TAG)
		p.expectStatus(TS_ATTRIBUTE)
	} else if p.peek == GREATER_THAN {
		// <m>
		p.addToken(TT_START_TAG)
		p.next()
		p.expectStatus(TS_LITERAL)
	} else if p.peek == FORWARD_SLASH {
		// <m/
		p.addToken(TT_START_TAG)
		p.next()
		p.expectStatus(TS_SELF_END_TAG)
	}
}

func (p *XMLTokenizer) parseEndTagName() {
	if p.peek == GREATER_THAN {
		p.addToken(TT_END_TAG)
		p.next()
		p.expectStatus(TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseSelfEndTag() {
	if p.peek == GREATER_THAN {
		p.addToken(TT_SELF_END_TAG)
		p.next()
		p.expectStatus(TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseAttribute() {
	if IsLetter(p.peek) {
		p.expectStatus(TS_ATTRIBUTE_NAME)
	} else if p.peek == FORWARD_SLASH {
		// <m ... /
		//p.next()
		p.expectStatus(TS_SELF_END_TAG)
	} else if p.peek == GREATER_THAN {
		// <m ... >
		p.next()
		p.expectStatus(TS_LITERAL)
	} else {
		// TODO 非法字符
	}
}

func (p *XMLTokenizer) parseAttributeName() {
	if IsBlank(p.peek) {
		p.addToken(TT_ATTR_NAME)
		p.expectStatus(TS_ATTRIBUTE_EQUAL)
	} else if p.peek == EQUAL_SIGN {
		// <m a=
		p.addToken(TT_ATTR_NAME)
		p.expectStatus(TS_ATTRIBUTE_VALUE_START)
	}
}

func (p *XMLTokenizer) parseAttributeEqual() {
	if p.peek == EQUAL_SIGN {
		p.expectStatus(TS_ATTRIBUTE_VALUE_START)
	} else if !IsBlank(p.peek) {
		// TODO 报错，非法字符，期待 =，属性未赋值
	}
}

func (p *XMLTokenizer) parseAttributeValueStart() {
	if p.peek == DOUBLE_QUOTE {
		// <m a="
		p.next()
		p.expectStatus(TS_ATTRIBUTE_VALUE_END)
	} else {
		// TODO 非法字符，期待 "
	}
}

func (p *XMLTokenizer) parseAttributeValueEnd() {
	if p.peek == DOUBLE_QUOTE {
		p.addToken(TT_ATTR_VALUE)
		p.expectStatus(TS_ATTRIBUTE)
	}
}

func (p *XMLTokenizer) char() string {
	return fmt.Sprintf("%c", p.peek)
}

func (p *XMLTokenizer) expectStatus(status int) {
	p.state = status
	p.start = p.pos.fork()
}

func (p *XMLTokenizer) fetchValue(start, end int) string {
	v := ""
	for i := start; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *XMLTokenizer) addToken(tokenType string) {
	value := p.fetchValue(p.start.index, p.pos.index)
	if tokenType == TT_TEXT && strings.TrimSpace(value) == "" {
		return
	}
	p.tokens = append(p.tokens, &Token{
		Value: value,
		Type:  tokenType,
		Start: &Point{Line: p.start.line, Column: p.start.column},
		End:   &Point{Line: p.pos.line, Column: p.pos.column - 1},
	})
}

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
			Type:   ST_TEXT,
			RAW:    p.peek.Value,
			Tokens: NewSQLTokenizer(p.peek.Start.Line, p.peek.Start.Column, p.peek.Value).Parse(),
		})
	case TT_START_TAG:
		p.addNode(&XMLNode{
			Type: ST_NODE,
			Name: p.peek.Value,
		})
	case TT_ATTR_NAME:
		attr := &XMLAttribute{
			Name:  strings.ToLower(p.peek.Value),
			Start: p.peek.Start,
			End:   p.peek.End,
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
