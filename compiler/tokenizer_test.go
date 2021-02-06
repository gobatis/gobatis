package compiler

import (
	"fmt"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	tokenizer := NewTokenizer([]byte(`abc`))
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
