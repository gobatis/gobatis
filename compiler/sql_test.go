package compiler

import (
	"fmt"
	"testing"
)

const testBaseSql = `select * from users where id=#{ User.Id},name=#{Name}`
const testBaseSql2 = "select * from users where id = 1"

func TestNewSQLTokenizer(t *testing.T) {
	tokenizer := NewSQLTokenizer(1, 0, testBaseSql2)
	tokens, err := tokenizer.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
