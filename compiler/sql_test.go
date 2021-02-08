package compiler

import (
	"fmt"
	"testing"
)

func TestNewSQLTokenizer(t *testing.T) {
	tokenizer := NewSQLTokenizer(1, 1, `
	  update Author set
		username = #{USER},
		password = #{password},
		email = #{email},
		bio = #{bio}
`)
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
