package main

import "github.com/koyeo/gobatis/compiler"

func main() {

	lexer := compiler.NewLexer(`<insert id="insertAuthor" useGeneratedKeys="true"
    keyProperty="id">
  insert into Author (username,password,email,bio)
  values (#{username},#{password},#{email},#{bio})
</insert>
`)

	lexer.Tokenlize()
	lexer.PrintTokens()
}
