package compiler

import (
	"errors"
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
	if index >= 0 && index < len(p.chars) {
		return p.chars[index]
	}
	return EOF
}

func (p *SQLTokenizer) last() rune {
	index := p.index - 1
	if index >= 0 && index < len(p.chars) {
		return p.chars[index]
	}
	return EOF
}

func (p *SQLTokenizer) skip(length int) {
	for i := 0; i < length; i++ {
		p.next()
	}
}

func (p *SQLTokenizer) Parse() (tokens []*Token, err error) {
	p.next()
	p.start = p.pos.fork()
	for p.look != EOF {
		switch p.state {
		case TS_BLANK:
			err = p.parseBlank()
			if err != nil {
				return
			}
		case TS_ID:
			err = p.parseId()
			if err != nil {
				return
			}
		case TS_SQL_VAR_START:
			p.parseSQLVarStart()
		case TS_SQL_VAR_VALUE_START:
			p.parseSQLVarValueStart()
		case TS_SQL_VAR_VALUE_END:
			err = p.parseSQLVarValueEnd()
			if err != nil {
				return
			}
		case TS_SQL_VAR_END:
			err = p.parseSQLVarEnd()
			if err != nil {
				return
			}
		case TS_SQL_ASTERISK:
			err = p.parseSQLAsterisk()
			if err != nil {
				return
			}
		case TS_SQL_OPEN_APOSTROPHE:
			p.parseOpenApostrophe()
		case TS_SQL_SINGLE_QUOTE:
			p.parseSingleQuote()
		case TS_SQL_DOUBLE_QUOTE:
			p.parseDoubleQuote()
		}
		p.next()
	}
	tokens = p.tokens
	return
}

func (p *SQLTokenizer) parseBlank() (err error) {
	if IsLetter(p.look) {
		p.forward(p.index, TS_ID)
	} else if p.look == OPEN_APOSTROPHE {
		p.forward(p.index+1, TS_SQL_OPEN_APOSTROPHE)
	} else if p.look == SINGLE_QUOTE {
		p.forward(p.index+1, TS_SQL_SINGLE_QUOTE)
	} else if p.look == DOUBLE_QUOTE {
		p.forward(p.index+1, TS_SQL_DOUBLE_QUOTE)
	} else if p.look == NUMBER_SIGN {
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else if p.look == COMMA {
		p.emitChar(TT_COMMA)
		p.forward(p.index+1, TS_BLANK)
	} else if p.look == EQUAL_SIGN {
		p.emitChar(TT_EQUAL)
		p.forward(p.index+1, TS_SQL_VAR_START)
	} else if p.look == ASTERISK {
		p.forward(p.index, TS_SQL_ASTERISK)
	} else if !IsBlank(p.look) {

		return p.newInvalidCharErr()
	}
	return
}

func (p *SQLTokenizer) parseId() (err error) {
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
	} else if !IsLetter(p.look) {
		return p.newInvalidCharErr()
	}
	return
}

func (p *SQLTokenizer) parseSQLAsterisk() (err error) {
	if IsBlank(p.look) {
		p.emitBuffer(TT_SQL_ASTERISK)
		p.forward(p.index, TS_BLANK)
	} else {
		return p.newInvalidCharErr()
	}
	return
}

func (p *SQLTokenizer) parseOpenApostrophe() {
	if p.look == OPEN_APOSTROPHE {
		p.emitBuffer(TT_ID)
		p.forward(p.index+1, TS_BLANK)
	}
}
func (p *SQLTokenizer) parseSingleQuote() {
	if p.look == SINGLE_QUOTE && p.last() != BACK_SLASH {
		p.emitBuffer(TT_ID)
		p.forward(p.index+1, TS_BLANK)
	}
}
func (p *SQLTokenizer) parseDoubleQuote() {
	if p.look == DOUBLE_QUOTE && p.last() != BACK_SLASH {
		p.emitBuffer(TT_ID)
		p.forward(p.index+1, TS_BLANK)
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

func (p *SQLTokenizer) parseSQLVarValueEnd() (err error) {
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
		return p.newInvalidCharErr("letter")
	}
	return
}

func (p *SQLTokenizer) parseSQLVarEnd() (err error) {
	if p.look == RIGHT_BRACE {
		p.forward(p.index+1, TS_BLANK)
	} else if !IsBlank(p.look) {
		return p.newInvalidCharErr()
	}
	return
}

func (p *SQLTokenizer) newInvalidCharErr(expects ...string) error {
	msg := fmt.Sprintf("invalid char %c at line %d column %d",
		p.look, p.pos.line, p.pos.column)
	if len(expects) > 0 {
		msg += " expect"
		for _, v := range expects {
			msg += " " + v
		}
	}
	return errors.New(msg)
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
