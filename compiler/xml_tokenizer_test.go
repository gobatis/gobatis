package compiler

import (
	"fmt"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	tokenizer := NewXMLTokenizer([]byte(`
<insert id="insertAuthor">
  insert into Author (id,username,password,email,bio)
  values (#{id},#{username},#{password},#{email},#{bio})
</insert>

<update id="updateAuthor">
  update Author set
    username = #{username},
    password = #{password},
    email = #{email},
    bio = #{bio}
  where id = #{id}
</update>

<delete id="deleteAuthor">
  delete from Author where id = #{id}
</delete>
`))
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v.String())
	}
}
