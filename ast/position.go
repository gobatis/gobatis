package ast

// 记录代码位置
type Position struct {
	index   int    // 索引
	row     int    // 行号
	col     int    // 列号
	file    string // 所属文件
	content string // 文件内容
}

func (p *Position) fork() Position {
	return Position{
		index:   p.index,
		row:     p.row,
		col:     p.col,
		file:    p.file,
		content: p.content,
	}
}

func (p *Position) next(char string) {
	p.index++
	if char == "\n" {
		p.col = 0
		p.row++
	} else {
		p.col++
	}
}

func (p *Position) peek(length int) int {
	return p.index + length
}
