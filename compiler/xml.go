package compiler

import (
	"fmt"
	"strings"
)

type XMLNode struct {
	Start        *Point          `json:"-"`
	End          *Point          `json:"-"`
	Type         string          `json:"type"`
	Name         string          `json:"name,omitempty"`
	Value        string          `json:"value,omitempty"`
	RAW          string          `json:"value,omitempty"`
	Attributes   []*XMLAttribute `json:"Attributes,omitempty"`
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
			line: 1,
		},
		index: -1,
		chars: []rune(string(content)),
		state: TS_LITERAL,
	}
}

type XMLTokenizer struct {
	file   string
	pos    *Position
	start  *Position
	index  int
	begin  int // 词素开始位置
	chars  []rune
	look   rune
	state  int
	tokens []*Token
}

func (p *XMLTokenizer) next() {
	//if p.index == -1 && len(p.chars) > 0 {
	//	p.look = p.chars[0]
	//}
	p.pos.next(p.look)
	p.index++
	if p.index < len(p.chars) {
		p.look = p.chars[p.index]
	} else {
		p.look = EOF
	}
}

func (p *XMLTokenizer) peek(length int) rune {
	index := p.index + length
	if index < len(p.chars) {
		return p.chars[index]
	}
	return EOF
}

func (p *XMLTokenizer) skip(length int) {
	for i := 0; i < length; i++ {
		p.next()
	}
}

func (p *XMLTokenizer) Parse() (tokens []*Token, err error) {
	p.next()
	p.start = p.pos.fork()
	for p.look != EOF {
		switch p.state {
		case TS_LITERAL:
			p.parseLiteral()
		case TS_START_TAG:
			err = p.parseTagStart()
			if err != nil {
				return
			}
		case TS_START_TAG_NAME:
			p.parseStartTagName()
		case TS_END_TAG_NAME:
			p.parseEndTagName()
		case TS_ATTRIBUTE:
			err = p.parseAttribute()
			if err != nil {
				return
			}
		case TS_ATTRIBUTE_NAME:
			p.parseAttributeName()
		case TS_ATTRIBUTE_EQUAL:
			err = p.parseAttributeEqual()
			if err != nil {
				return
			}
		case TS_ATTRIBUTE_VALUE_START:
			err = p.parseAttributeValueStart()
			if err != nil {
				return
			}
		case TS_ATTRIBUTE_VALUE_END:
			p.parseAttributeValueEnd()
		case TS_SELF_END_TAG:
			p.parseSelfEndTag()
		case TS_DOCTYPPE:
			p.parseDoctype()
		case TS_STATEMENT:
			err = p.parseStatement()
			if err != nil {
				return
			}
		case TS_COMMENT_START:
			err = p.parseCommentStart()
			if err != nil {
				return
			}
		case TS_COMMENT_END:
			p.parseCommentEnd()
		}
		p.next()
	}
	switch p.state {
	case TS_LITERAL:
		p.emitBuffer(TT_TEXT)
	}
	tokens = p.tokens
	return
}

func (p *XMLTokenizer) parseLiteral() {
	if p.look == LESS_THAN {
		p.emitBuffer(TT_TEXT)
		p.forward(p.index+1, TS_START_TAG)
	}
}

func (p *XMLTokenizer) parseTagStart() (err error) {
	if IsLetter(p.look) {
		p.forward(p.index, TS_START_TAG_NAME)
	} else if p.look == QUESTION_MARK {
		// <?
		if IsLetter(p.peek(1)) {
			p.forward(p.index+1, TS_STATEMENT)
		} else {
			return p.newInvalidCharErr()
		}
	} else if p.look == EXCLAMATION_MARK {
		// <!
		p.forward(p.index+1, TS_COMMENT_START)
	} else if p.look == FORWARD_SLASH {
		// </
		p.forward(p.index+1, TS_END_TAG_NAME)
	}
	return
}

func (p *XMLTokenizer) parseDoctype() {
	if p.look == GREATER_THAN {
		p.forward(p.index+1, TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseStatement() (err error) {
	if p.look == QUESTION_MARK {
		p.next()
		if p.look != GREATER_THAN {
			return p.newInvalidCharErr()
		}
		p.forward(p.index+1, TS_LITERAL)
	}
	return
}

func (p *XMLTokenizer) parseCommentStart() (err error) {
	// 解析注释
	if p.look == MINUS {
		if p.peek(1) == MINUS {
			p.skip(1)
			p.forward(p.index+1, TS_COMMENT_END)
		} else {
			return p.newInvalidCharErr()
		}
	} else if strings.ToLower(p.peekString(7)) == DOCTYPE {
		p.skip(7)
		p.forward(p.index+7, TS_DOCTYPPE)
	} else {
		return p.newInvalidCharErr()
	}
	return
}

func (p *XMLTokenizer) parseCommentEnd() {
	if p.look == MINUS && p.peek(1) == MINUS && p.peek(2) == GREATER_THAN {
		p.skip(2) // look = >
		p.forward(p.index+1, TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseStartTagName() {

	if IsBlank(p.look) {
		// <m
		p.emitBuffer(TT_START_TAG)
		p.forward(p.index+1, TS_ATTRIBUTE)
	} else if p.look == GREATER_THAN {
		// <m>
		p.emitBuffer(TT_START_TAG)
		p.forward(p.index+1, TS_LITERAL)
	} else if p.look == FORWARD_SLASH {
		// <m/
		p.emitBuffer(TT_START_TAG)
		p.forward(p.index+1, TS_SELF_END_TAG)
	}
}

func (p *XMLTokenizer) parseEndTagName() {
	if p.look == GREATER_THAN {
		p.emitBuffer(TT_END_TAG)
		p.forward(p.index+1, TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseSelfEndTag() {
	if p.look == GREATER_THAN {
		p.emitBuffer(TT_SELF_END_TAG)
		p.forward(p.index+1, TS_LITERAL)
	}
}

func (p *XMLTokenizer) parseAttribute() (err error) {
	if IsLetter(p.look) {
		p.forward(p.index, TS_ATTRIBUTE_NAME)
	} else if p.look == FORWARD_SLASH {
		// <m ... /
		p.forward(p.index+1, TS_SELF_END_TAG)
	} else if p.look == GREATER_THAN {
		// <m ... >
		p.forward(p.index+1, TS_LITERAL)
	} else {
		return p.newInvalidCharErr()
	}
	return
}

func (p *XMLTokenizer) parseAttributeName() {
	if IsBlank(p.look) {
		p.emitBuffer(TT_ATTR_NAME)
		p.forward(p.index+1, TS_ATTRIBUTE_EQUAL)
	} else if p.look == EQUAL_SIGN {
		// <m a=
		p.emitBuffer(TT_ATTR_NAME)
		p.forward(p.index+1, TS_ATTRIBUTE_VALUE_START)
	}
}

func (p *XMLTokenizer) parseAttributeEqual() (err error) {
	if p.look == EQUAL_SIGN {
		p.forward(p.index+1, TS_ATTRIBUTE_VALUE_START)
	} else if !IsBlank(p.look) {
		return p.newInvalidCharErr()
	}
	return
}

func (p *XMLTokenizer) parseAttributeValueStart() (err error) {
	if p.look == DOUBLE_QUOTE {
		// <m a="
		p.forward(p.index+1, TS_ATTRIBUTE_VALUE_END)
	} else {
		return p.newInvalidCharErr()
	}
	return
}

func (p *XMLTokenizer) parseAttributeValueEnd() {
	if p.look == DOUBLE_QUOTE {
		p.emitBuffer(TT_ATTR_VALUE)
		p.forward(p.index+1, TS_ATTRIBUTE)
	}
}

func (p *XMLTokenizer) char() string {
	return fmt.Sprintf("%c", p.look)
}

func (p *XMLTokenizer) forward(index, status int) {
	p.state = status
	p.start = p.pos.fork()
	p.begin = index
}

func (p *XMLTokenizer) subString(end int) string {
	v := ""
	for i := p.begin; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *XMLTokenizer) peekString(length int) string {
	r := make([]rune, 0)
	for i := 0; i < length; i++ {
		index := p.index + i
		if index < len(p.chars) {
			r = append(r, p.chars[index])
		}
	}
	return string(r)
}

func (p *XMLTokenizer) emitBuffer(tokenType string) {
	value := p.subString(p.index)
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

func (p *XMLTokenizer) emitChar(tokenType string) {
	value := string(p.chars[p.index])
	p.tokens = append(p.tokens, &Token{
		Value: value,
		Type:  tokenType,
		Start: &Point{Line: p.pos.line, Column: p.pos.column},
		End:   &Point{Line: p.pos.line, Column: p.pos.column},
	})
}

func (p *XMLTokenizer) newInvalidCharErr() error {
	return fmt.Errorf(
		"invalid char %c at line %d column %d",
		p.look, p.start.line, p.start.column,
	)
}

type XMLParser struct {
	tokens []*Token
	look   *Token
	body   []*XMLNode
	nodes  []*XMLNode
	state  int
	index  int
}

func NewXMLParser() *XMLParser {
	return &XMLParser{index: -1}
}

func (p *XMLParser) next() {
	p.index += 1
	if p.index < len(p.tokens) {
		p.look = p.tokens[p.index]
	} else {
		p.look = nil
	}
}

func (p *XMLParser) Parse(content []byte) (xmlNodes []*XMLNode, err error) {
	tokenizer := NewXMLTokenizer(content)
	tokens, err := tokenizer.Parse()
	if err != nil {
		return
	}
	return p.ParseTokens(tokens)
}

func (p *XMLParser) ParseTokens(tokens []*Token) (xmlNodes []*XMLNode, err error) {
	p.tokens = tokens
	p.next()
	for p.look != nil {
		err = p.parse()
		if err != nil {
			return
		}
		p.next()
	}
	xmlNodes = p.body
	return
}

func (p *XMLParser) parse() (err error) {

	switch p.look.Type {
	case TT_TEXT:
		var tokens []*Token
		tokens, err = NewSQLTokenizer(p.look.Start.Line, p.look.Start.Column, p.look.Value).Parse()
		if err != nil {
			return
		}
		p.addNode(&XMLNode{
			Type:   ST_TEXT,
			RAW:    p.look.Value,
			Tokens: tokens,
		})
	case TT_START_TAG:
		p.addNode(&XMLNode{
			Type: ST_NODE,
			Name: p.look.Value,
		})
	case TT_ATTR_NAME:
		attr := &XMLAttribute{
			Name:  p.look.Value,
			Start: p.look.Start,
			End:   p.look.End,
		}
		p.next()
		if p.look.Type != TT_ATTR_VALUE {
			err = p.newEmptyAttributeErr(attr.Name)
			return
		}
		attr.Name = strings.ToLower(attr.Name)
		attr.Value = p.look.Value
		err = p.setAttribute(p.lastUnclosedNode(), attr)
		if err != nil {
			return
		}
	case TT_SELF_END_TAG:
		p.lastUnclosedNode().closed = true
		p.nodes = p.nodes[:len(p.nodes)-1]
	case TT_END_TAG:
		if p.lastUnclosedNode() == nil ||
			p.lastUnclosedNode().Name != strings.ToLower(p.look.Value) {
			err = p.newTagNotClosedErr(p.look.Value)
			return
		}
		p.lastUnclosedNode().closed = true
		p.nodes = p.nodes[:len(p.nodes)-1]
	default:
		err = p.newUnexpectedTokenErr(p.look.Value)
	}
	return
}

func (p *XMLParser) setAttribute(node *XMLNode, attr *XMLAttribute) (err error) {
	if node.attributeMap == nil {
		node.attributeMap = map[string]string{}

	}
	_, ok := node.attributeMap[attr.Name]
	if ok {
		err = p.newAttributeDuplicateErr(attr.Name)
		return
	}

	node.attributeMap[attr.Name] = attr.Value
	node.Attributes = append(node.Attributes, attr)
	return
}

func (p *XMLParser) addNode(node *XMLNode) {
	node.Start = p.look.Start
	node.End = p.look.End
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
	l := len(p.nodes)
	if l > 0 {
		return p.nodes[l-1]
	}
	return nil
}

func (p *XMLParser) newTagNotClosedErr(name string) (err error) {
	return fmt.Errorf("tag '%s' not closed at line %d column %d", name, p.look.Start.Line, p.look.Start.Column)
}

func (p *XMLParser) newUnexpectedTokenErr(name string) (err error) {
	return fmt.Errorf("upexected token '%s' at line %d column %d", name, p.look.Start.Line, p.look.Start.Column)
}

func (p *XMLParser) newEmptyAttributeErr(name string) error {
	return fmt.Errorf("empty attribute '%s' at line %d column %d", name, p.look.Start.Line, p.look.Start.Column)
}

func (p *XMLParser) newAttributeDuplicateErr(name string) error {
	return fmt.Errorf("duplicate attribute '%s' at line %d column %d", name, p.look.Start.Line, p.look.Start.Column)
}
