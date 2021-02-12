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
		index: -1,
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
	p.index++
	if p.index < len(p.chars) {
		p.look = p.chars[p.index]
	} else {
		p.look = EOF
	}
}

func (p *SQLTokenizer) peekString(length int) string {
	r := make([]rune, 0)
	for i := 0; i < length; i++ {
		index := p.index + i
		if index < len(p.chars) {
			r = append(r, p.chars[index])
		}
	}
	return string(r)
}

func (p *SQLTokenizer) peek(length int) rune {
	index := p.index + length
	if index < len(p.chars) {
		return p.chars[index]
	}
	return EOF
}

func (p *SQLTokenizer) skip(length int) {
	for i := 0; i < length; i++ {
		p.next()
	}
}

func (p *SQLTokenizer) Parse() []*Token {
	p.next()
	p.start = p.pos.fork()
	for p.look != EOF {
		switch p.state {
		case TS_BLANK:
			p.parseBlank()
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

func (p *SQLTokenizer) parseBlank() {
	if IsLetter(p.look) {
		p.forward(p.index, TS_ID)
	} else if p.look == NUMBER_SIGN {
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else if p.look == COMMA {
		p.emitBuffer(TT_COMMA)
		p.forward(p.index+1, TS_BLANK)
	} else if p.look == EQUAL_SIGN {
		p.emitBuffer(TT_EQUAL)
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else if !IsBlank(p.look) {
		// TODO 报错，异常字符，出现了非字母字符开头的字符

	}

}

func (p *SQLTokenizer) parseId() {
	if IsBlank(p.look) {
		p.emitBuffer(TT_ID)
		p.forward(p.index, TS_BLANK)
	} else if p.look == EQUAL_SIGN {
		// id=
		p.emitBuffer(TT_ID)
		p.emitChar(TT_EQUAL)
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else if p.look == NUMBER_SIGN {
		// select#
		// TODO 与 ID 相连的 # 号，可能不需要
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else {
		// TODO 非法变量，ID 只能是字母开头，数字+下划线
	}
}

func (p *SQLTokenizer) parseSQLVarStart() {
	if p.look == LEFT_BRACE {
		p.forward(p.index+1, TS_SQL_VAR_VALUE_START)
	}
}

func (p *SQLTokenizer) parseSQLVarValueStart() {
	if IsLetter(p.look) {
		p.forward(p.index, TS_SQL_VAR_VALUE_END)
	}
}

func (p *SQLTokenizer) parseSQLVarValueEnd() {
	if p.look == DOT {
		p.emitBuffer(TT_SQL_STRUCT)
		p.emitChar(TT_SQL_DOT)
		p.forward(p.index+1, TS_SQL_VAR_VALUE_END)
	} else if p.look == RIGHT_BRACE {
		p.emitBuffer(TT_SQL_VAR)
		p.forward(p.index+1, TS_BLANK)
	} else if IsBlank(p.look) {
		p.emitBuffer(TT_SQL_VAR)
		p.forward(p.index+1, TS_SQL_VAR_END)
	} else if !IsLetter(p.look) {
		// TODO 非法字符
	}
}

func (p *SQLTokenizer) parseSQLVarEnd() {
	if p.look == RIGHT_BRACE {
		p.forward(p.index+1, TS_BLANK)
	} else if !IsBlank(p.look) {
		// TODO 非法字符
	}
}

func (p *SQLTokenizer) subString(end int) string {
	v := ""
	for i := p.begin; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *SQLTokenizer) char() string {
	return fmt.Sprintf("%c", p.look)
}

func (p *SQLTokenizer) forward(index, status int) {
	p.state = status
	p.start = p.pos.fork()
	p.begin = index
}

func (p *SQLTokenizer) emitBuffer(tokenType string) {
	value := p.subString(p.index)
	p.tokens = append(p.tokens, &Token{
		Value: value,
		Type:  tokenType,
		Start: &Point{Line: p.start.line, Column: p.start.column},
		End:   &Point{Line: p.pos.line, Column: p.pos.column - 1},
	})
}

func (p *SQLTokenizer) emitChar(tokenType string) {
	value := string(p.chars[p.index])
	p.tokens = append(p.tokens, &Token{
		Value: value,
		Type:  tokenType,
		Start: &Point{Line: p.pos.line, Column: p.pos.column},
		End:   &Point{Line: p.pos.line, Column: p.pos.column},
	})
}
