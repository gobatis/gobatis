package compiler

import (
	"fmt"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	tokenizer := NewTokenizer([]byte(`abc<div id="123">你好呀</div>`))
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
