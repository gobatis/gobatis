package compiler

import (
	"fmt"
	"testing"
)

const testBaseSql = `id=#{ User.Id},name=#{Name}`

func TestNewSQLTokenizer(t *testing.T) {
	tokenizer := NewSQLTokenizer(1, 0, testBaseSql)
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
