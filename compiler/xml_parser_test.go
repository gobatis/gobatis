package compiler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXMLParser_Parse(t *testing.T) {
	tokenizer := NewXMLTokenizer([]byte(`
<mapper>
	<update id="updateAuthor">
	  update Author set
		username = #{ User.Name },
		password = #{password},
		email = #{email},
		bio = #{bio}
	  <where>
		<if test="id != nil">
		id = #{id}
		</if>
	  </where> 
	</update>
</mapper>
`))
	tokens := tokenizer.Parse()
	for _, v := range tokens {
		fmt.Println(v)
	}
	parser := NewXMLParser(tokens)
	d, _ := json.MarshalIndent(parser.Parse(), "", "\t")
	fmt.Println(string(d))
}
