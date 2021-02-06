package compiler

import "fmt"

const (
	TT_TEXT         = "text"
	TT_START_TAG    = "startTag"
	TT_END_TAG      = "endTag"
	TT_SELF_END_TAG = "selfEndTag"
	TT_ATTR_VALUE   = "attrValue"
	TT_ATTR_NAME    = "attrName"
)

func NewTokenizer(content []byte) Tokenizer {
	return Tokenizer{
		pos: Position{
			index:  -1,
			line:   1,
			column: 0,
		},
		chars: []rune(string(content)),
		state: STAT_LITERAL,
	}
}

type Tokenizer struct {
	file   string
	pos    Position
	start  Position
	chars  []rune
	peek   rune
	state  int
	tokens []Token
}

func (p *Tokenizer) next() {
	p.pos.next(p.peek)
	if p.pos.index < len(p.chars) {
		p.peek = p.chars[p.pos.index]
	} else {
		p.peek = EOF
	}
}

func (p *Tokenizer) Parse() []Token {
	p.next()
	p.start = p.pos.fork()
	for p.peek != EOF {
		switch p.state {
		case STAT_LITERAL:
			p.parseLiteral()
		case STAT_START_TAG:
			p.parseTagStart()
		case STAT_START_TAG_NAME:
			p.parseStartTagName()
		case STAT_END_TAG_NAME:
			p.parseEndTagName()
		case STAT_ATTRIBUTE:
			p.parseAttribute()
		case STAT_ATTRIBUTE_NAME:
			p.parseAttributeName()
		case STAT_ATTRIBUTE_EQUAL:
			p.parseAttributeEqual()
		case STAT_ATTRIBUTE_VALUE_START:
			p.parseAttributeValueStart()
		case STAT_ATTRIBUTE_VALUE_END:
			p.parseAttributeValueEnd()
		case STAT_SELF_END_TAG:
			p.parseSelfEndTag()
		}
		p.next()
	}
	switch p.state {
	case STAT_LITERAL:
		p.addToken(TT_TEXT)
	}
	return p.tokens
}

func (p *Tokenizer) addToken(tokenType string) {
	value := p.fetchValue(p.start.index, p.pos.index)
	if tokenType == TT_TEXT && value == "" {
		return
	}
	p.tokens = append(p.tokens, Token{
		Value: value,
		Type:  tokenType,
		Start: TokenLoc{Line: p.start.line, Column: p.start.column},
		End:   TokenLoc{Line: p.pos.line, Column: p.pos.column},
	})
}

func (p *Tokenizer) fetchValue(start, end int) string {
	v := ""
	for i := start; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *Tokenizer) parseLiteral() {
	if p.peek == LESS_THAN {
		p.addToken(TT_TEXT)
		p.expectStatus(STAT_START_TAG)
	}
}

func (p *Tokenizer) parseTagStart() {
	if IsLetter(p.peek) {
		p.expectStatus(STAT_START_TAG_NAME)
	} else if p.peek == FORWARD_SLASH {
		p.next()
		p.expectStatus(STAT_END_TAG_NAME)
	}
}

func (p *Tokenizer) parseStartTagName() {

	if IsBlank(p.peek) {
		// <m
		p.addToken(TT_START_TAG)
		p.expectStatus(STAT_ATTRIBUTE)
	} else if p.peek == GREATER_THAN {
		// <m>
		p.addToken(TT_START_TAG)
		p.next()
		p.expectStatus(STAT_LITERAL)
	} else if p.peek == FORWARD_SLASH {
		// <m/
		p.addToken(TT_START_TAG)
		p.next()
		p.expectStatus(STAT_SELF_END_TAG)
	}
}

func (p *Tokenizer) parseEndTagName() {
	if p.peek == GREATER_THAN {
		p.addToken(TT_END_TAG)
		p.next()
		p.expectStatus(STAT_LITERAL)
	}
}

func (p *Tokenizer) parseSelfEndTag() {
	if p.peek == GREATER_THAN {
		p.addToken(TT_SELF_END_TAG)
		p.expectStatus(STAT_LITERAL)
	}
}

func (p *Tokenizer) parseAttribute() {
	if IsLetter(p.peek) {
		p.expectStatus(STAT_ATTRIBUTE_NAME)
	} else if p.peek == FORWARD_SLASH {
		// <m ... /
		p.next()
		p.expectStatus(STAT_SELF_END_TAG)
	} else if p.peek == GREATER_THAN {
		// <m ... >
		p.next()
		p.expectStatus(STAT_LITERAL)
	} else {
		// TODO 非法字符
	}
}

func (p *Tokenizer) parseAttributeName() {
	if IsBlank(p.peek) {
		p.addToken(TT_ATTR_NAME)
		p.expectStatus(STAT_ATTRIBUTE_EQUAL)
	} else if p.peek == EQUAL_SIGN {
		// <m a=
		p.addToken(TT_ATTR_NAME)
		p.expectStatus(STAT_ATTRIBUTE_VALUE_START)
	}
}

func (p *Tokenizer) parseAttributeEqual() {
	if p.peek == EQUAL_SIGN {
		p.expectStatus(STAT_ATTRIBUTE_VALUE_START)
	} else if !IsBlank(p.peek) {
		// TODO 报错，非法字符，期待 =，属性未赋值
	}
}

func (p *Tokenizer) parseAttributeValueStart() {
	if p.peek == DOUBLE_QUOTE {
		// <m a="
		p.next()
		p.expectStatus(STAT_ATTRIBUTE_VALUE_END)
	} else {
		// TODO 非法字符，期待 "
	}
}

func (p *Tokenizer) parseAttributeValueEnd() {
	if p.peek == DOUBLE_QUOTE {
		p.addToken(TT_ATTR_VALUE)
		p.expectStatus(STAT_ATTRIBUTE)
	}
}

func (p *Tokenizer) char() string {
	return fmt.Sprintf("%c", p.peek)
}

func (p *Tokenizer) expectStatus(status int) {
	p.state = status
	p.start = p.pos.fork()
}
