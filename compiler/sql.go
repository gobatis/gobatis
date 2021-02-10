package compiler

import (
	"fmt"
)

func NewSQLTokenizer(line, column int, content string) *SQLTokenizer {
	return &SQLTokenizer{
		pos: &Position{
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
	look   rune
	state  int
	tokens []*Token
	index  int
	begin  int
}

func (p *SQLTokenizer) next() {
	p.pos.next(p.look)
	if p.index < len(p.chars) {
		p.look = p.chars[p.index]
	} else {
		p.look = EOF
	}
}

func (p *SQLTokenizer) Parse() []*Token {
	p.next()
	p.start = p.pos.fork()
	for p.look != EOF {
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
	if IsLetter(p.look) {
		p.expectStatus(TS_ID)
	} else if p.look == NUMBER_SIGN {
		p.expectStatus(TS_SQL_VAR_START)
	} else if p.look == COMMA {
		p.parseComma()
	} else if p.look == EQUAL_SIGN {
		p.parseEqual()
	} else if !IsBlank(p.look) {
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
	if IsBlank(p.look) {
		p.addToken(TT_ID)
		p.expectStatus(TS_BLANK)
	} else if p.look == EQUAL_SIGN {
		p.addToken(TT_ID)
		p.parseEqual()
	} else if p.look == NUMBER_SIGN {
		// select#
		// TODO 与 ID 相连的 # 号，可能不需要
		p.expectStatus(TS_SQL_VAR_START)
	}
}

func (p *SQLTokenizer) parseSQLVarStart() {
	if p.look == LEFT_BRACE {
		p.expectStatus(TS_SQL_VAR_VALUE_START)
	}
}

func (p *SQLTokenizer) parseSQLVarValueStart() {
	if IsLetter(p.look) {
		p.expectStatus(TS_SQL_VAR_VALUE_END)
	}
}

func (p *SQLTokenizer) parseSQLVarValueEnd() {
	if p.look == DOT {
		p.addToken(TT_SQL_STRUCT)
		p.start = p.pos.fork()
		p.next()
		p.addToken(TT_SQL_DOT)
		p.expectStatus(TS_SQL_VAR_VALUE_END)
	} else if p.look == RIGHT_BRACE {
		p.addToken(TT_SQL_VAR)
		p.expectStatus(TS_BLANK)
	} else if IsBlank(p.look) {
		p.addToken(TT_SQL_VAR)
		p.expectStatus(TS_SQL_VAR_END)
	} else if !IsLetter(p.look) {
		// TODO 非法字符
	}
}

func (p *SQLTokenizer) parseSQLVarEnd() {
	if p.look == RIGHT_BRACE {
		p.expectStatus(TS_BLANK)
	} else if !IsBlank(p.look) {
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
	return fmt.Sprintf("%c", p.look)
}

func (p *SQLTokenizer) expectStatus(status int) {
	p.state = status
	p.start = p.pos.fork()
}

func (p *SQLTokenizer) addToken(tokenType string) {
	value := p.fetchValue(p.begin, p.index)
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
