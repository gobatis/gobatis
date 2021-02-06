package compiler

import (
	"encoding/json"
)

type Token struct {
	Type  int      `json:"type"`
	Value string   `json:"value"`
	Start TokenLoc `json:"start"`
	End   TokenLoc `json:"end"`
}

func (p *Token) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

type TokenLoc struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
