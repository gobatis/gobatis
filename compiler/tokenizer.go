package compiler

const (
	SPACE = 32  // ' '
	BSN   = 10  // \n
	BST   = 9   // \t
	BSR   = 13  // \r
	BSF   = 12  // \f
	LT    = 60  // <
	EP    = 33  // !
	CL    = 45  // -
	SL    = 47  // /
	GT    = 62  // >
	QM    = 63  // ?
	LA    = 97  // a
	LZ    = 122 // z
	UA    = 65  // A
	UZ    = 90  // Z
	EQ    = 61  // =
	SQ    = 39  // '
	DQ    = 34  // "
	LD    = 100 // d
	UD    = 68  //D
)

const (
	TT_LITERAL        = iota
	TT_OPEN_TAG_START // <div>
	TT_OPEN_TAG_END   // </div>
	TT_CLOSE_Tag      // <img />
	TT_ATTR_VALUE
	TT_ATTR_NAME
)

const (
	STAT_LITERAL      = iota // 解析文本
	STAT_PRE_OPEN_TAG        // 预备解析开始标签类型
	STAT_OPEN_TAG            // 解析开标准
	STAT_POS_OPEN_TAG        // 解析结束标签
	STAT_IN_VALUE_NQ         // 无引号的值
	STAT_IN_VALUE_SQ         // 单引号值
	STAT_IN_VALUE_DQ         // 双引号值
	STAT_CLOSING_OPEN_TAG
	STAT_OPENING_NORMAL_COMMENT
	STAT_IM_NORMAL_COMMENT
	STAT_IN_SHORT_COMMENT
	STAT_CLOSING_NORMAL_COMMENT
	STAT_CLOSING_TAG
)

func NewTokenizer(content []byte) Tokenizer {
	return Tokenizer{
		pos: Position{
			index:  -1,
			line:   1,
			column: 0,
		},
		chars:  []rune(string(content)),
		status: STAT_LITERAL,
	}
}

type Tokenizer struct {
	file   string
	pos    Position
	start  Position
	chars  []rune
	peek   rune
	status int
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
		switch p.status {
		case STAT_LITERAL:
			p.parseLiteral()
		case STAT_PRE_OPEN_TAG:
			p.parsePreOpenTag()
		case STAT_OPEN_TAG:
			p.parseOpenTag()
		}
		p.next()
	}

	return p.tokens
}

func (p *Tokenizer) addToken(tokenType, nextStat int, end Position) {
	value := p.fetchValue(p.start.index, end.index)
	p.tokens = append(p.tokens, Token{
		Value: value,
		Type:  tokenType,
		Start: TokenLoc{Line: p.start.line, Column: p.start.column},
		End:   TokenLoc{Line: end.line, Column: end.column},
	})
	p.start = end
	p.status = nextStat
}

func (p *Tokenizer) fetchValue(start, end int) string {
	v := ""
	for i := start; i < end; i++ {
		v += string(p.chars[i])
	}
	return v
}

func (p *Tokenizer) parseLiteral() {
	if p.peek == LT {
		p.addToken(TT_LITERAL, STAT_PRE_OPEN_TAG, p.pos)
	}
}

func (p *Tokenizer) parsePreOpenTag() {
	if (p.peek >= LA && p.peek <= LZ) ||
		(p.peek >= UA && p.peek <= UZ) {
		p.status = STAT_OPEN_TAG
		p.start = p.pos.fork()
	}
}

func (p *Tokenizer) parseOpenTag() {
	if p.isBlack() {
		p.addToken(TT_OPEN_TAG_START, STAT_POS_OPEN_TAG, p.pos)
	}
}

func (p *Tokenizer) isBlack() bool {
	return p.peek == SPACE ||
		p.peek == BSN ||
		p.peek == BSR ||
		p.peek == BST ||
		p.peek == BSF
}
