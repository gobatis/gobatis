package compiler

import (
	"fmt"
	"strings"
)

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
