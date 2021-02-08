package compiler

import (
	"encoding/json"
)

type Token struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Start *Point `json:"start"`
	End   *Point `json:"end"`
}

func (p *Token) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

