package ast

// 记录代码位置
type Position struct {
	Index  int    `json:"index,omitempty"`  // 索引
	Line   int    `json:"line,omitempty"`   // 行号
	Column int    `json:"column,omitempty"` // 列号
	// File   string `json:"file,omitempty"`   // 所属文件
}

func (p *Position) fork() *Position {
	return &Position{
		//Index:  p.Index,
		Line:   p.Line,
		Column: p.Column,
		//File:   p.File,
	}
}

func (p *Position) next(char string) {
	p.Index++
	if char == "\n" {
		p.Column = 0
		p.Line++
	} else {
		p.Column++
	}
}

func (p *Position) peek(length int) int {
	return p.Index + length
}
