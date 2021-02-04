package main

import "github.com/koyeo/gobatis/ast"

func main() {

	lexer := ast.NewLexer("./test/sql/user.mapper.xml", `<select id="GetUser" parameterType="int64" resultType="test/entity@User">
        	select * from users #{id}
			<mapper src="ok" />
			<where>
				<if test="a != nil">a</if>
			</where>
    	</select>
`)

	lexer.Parse()
	lexer.PrintTokens()
}
