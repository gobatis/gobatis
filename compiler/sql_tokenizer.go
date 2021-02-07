package compiler

import (
	"fmt"
)

func NewSQLTokenizer(line, column int, content string) *SQLTokenizer {
	return &SQLTokenizer{
		pos: &Position{
			index:  -1,
			line:   line,
			column: column,
		},
		chars: []rune(content),
		state: TS_BLANK,
	}
}

type SQLTokenizer struct {
	file   string
	pos    *Position
	start  *Position
	chars  []rune
	peek   rune
	state  int
	tokens []*Token
}

func (p *SQLTokenizer) next() {
	p.pos.next(p.peek)
	if p.pos.index < len(p.chars) {
		p.peek = p.chars[p.pos.index]
	} else {
		p.peek = EOF
	}
}

func (p *SQLTokenizer) Parse() []*Token {
	p.next()
	p.start = p.pos.fork()
	for p.peek != EOF {
		switch p.state {
		case TS_BLANK:
			p.parseBlack()
		case TS_ID:
			p.parseId()
		case TS_SQL_VAR_START:
			p.parseSQLVarStart()
		case TS_SQL_VAR_VALUE_START:
			p.parseSQLVarValueStart()
		case TS_SQL_VAR_VALUE_END:
			p.parseSQLVarValueEnd()
		case TS_SQL_VAR_END:
			p.parseSQLVarEnd()
		}
		p.next()
	}

	return p.tokens
}

func (p *SQLTokenizer) parseBlack() {
	if IsLetter(p.peek) {
		p.expectStatus(TS_ID)
	} else if p.peek == NUMBER_SIGN {
		p.expectStatus(TS_SQL_VAR_START)
	} else if p.peek == COMMA {
		p.parseComma()
	} else if p.peek == EQUAL_SIGN {
		p.parseEqual()
	} else if !IsBlank(p.peek) {
		// TODO 报错，异常字符，出现了非字母字符
	}

}

func (p *SQLTokenizer) parseComma() {
	p.expectStatus(TS_BLANK)
	p.next()
	p.addToken(TT_COMMA)
}

func (p *SQLTokenizer) parseEqual() {
	p.expectStatus(TS_SQL_VAR_START)
	p.next()
	p.addToken(TT_EQUAL)
}

func (p *SQLTokenizer) parseId() {
	if IsBlank(p.peek) {
		p.addToken(TT_ID)
		p.expectStatus(TS_BLANK)
	} else if p.peek == EQUAL_SIGN {
		p.addToken(TT_ID)
		p.parseEqual()
	} else if p.peek == NUMBER_SIGN {
		// select#
		// TODO 与 ID 相连的 # 号，可能不需要
		p.expectStatus(TS_SQL_VAR_START)
	}
}

func (p *SQLTokenizer) parseSQLVarStart() {
	if p.peek == LEFT_BRACE {
		p.expectStatus(TS_SQL_VAR_VALUE_START)
	}
}

func (p *SQLTokenizer) parseSQLVarValueStart() {
	if IsLetter(p.peek) {
		p.expectStatus(TS_SQL_VAR_VALUE_END)
	}
}

func (p *SQLTokenizer) parseSQLVarValueEnd() {
	if p.peek == DOT {
		p.addToken(TT_SQL_STRUCT)
		p.start = p.pos.fork()
		p.next()
		p.addToken(TT_SQL_DOT)
		p.expectStatus(TS_SQL_VAR_VALUE_END)
	} else if p.peek == RIGHT_BRACE {
		p.addToken(TT_SQL_VAR)
		p.expectStatus(TS_BLANK)
	} else if IsBlank(p.peek) {
		p.addToken(TT_SQL_VAR)
		p.expectStatus(TS_SQL_VAR_END)
	} else if !IsLetter(p.peek) {
		// TODO 非法字符
	}
}

func (p *SQLTokenizer) parseSQLVarEnd() {
	if p.peek == RIGHT_BRACE {
		p.expectStatus(TS_BLANK)
	} else if !IsBlank(p.peek) {
		// TODO 非法字符
	}
}

func (p *SQLTokenizer) fetchValue(start, end int) string {
	v := ""
	for i := start; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *SQLTokenizer) char() string {
	return fmt.Sprintf("%c", p.peek)
}

func (p *SQLTokenizer) expectStatus(status int) {
	p.state = status
	p.start = p.pos.fork()
}

func (p *SQLTokenizer) addToken(tokenType string) {
	value := p.fetchValue(p.start.index, p.pos.index)
	//value = strings.TrimSpace(value)
	//if value == "" {
	//	return
	//}
	p.tokens = append(p.tokens, &Token{
		Value: value,
		Type:  tokenType,
		Start: &Point{Line: p.start.line, Column: p.start.column},
		End:   &Point{Line: p.pos.line, Column: p.pos.column - 1},
	})
}
