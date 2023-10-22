package xsql

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/logger"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	input := antlr.NewInputStream("$.User.Roles[*].ID")
	lexer := NewXSQLLexer(input)
	for {
		tok := lexer.NextToken()
		fmt.Println(tok.GetText(), tok.GetTokenType())
		if tok.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}

//func TestReplaceIsolatedLessThan(t *testing.T) {
//	tests := []struct {
//		input    string
//		expected string
//	}{
//		{"This is a <test> string with a single < and another <tag> and one more <.", "This is a <test> string with a single &lt; and another <tag> and one more &lt;."},
//		{"<Hello<world", "&lt;Hello&lt;world"},
//		{"Hello<world>", "Hello<world>"},
//		{"<<<<<", "&lt;&lt;&lt;&lt;&lt;"},
//		{"<", "&lt;"},
//		{"Hello<", "Hello&lt;"},
//		{">Hello", ">Hello"},
//		{"<10", "&lt;10"},
//	}
//
//	for _, test := range tests {
//		result := replaceIsolatedLessThanWithEntity(test.input)
//		require.Equal(t, test.expected, result)
//	}
//}

func TestParser(t *testing.T) {
	//xs, err := Parse(`< ok <if test="ok">abc</if> <a  < = >`)
	//xs, err := Parse(`><`)
	//xs, err := Parse(`<a test="ok" /> / > < !=`)
	//xs, err := Parse(`ok <a/> b > c`)
	//xs, err := Parse(`</>/></a</a><b/><c></c>d<!>`)
	//xs, err := Parse(`<a></a>*`)
	//xs, err := Parse(`*~<<a test="ok">b</a>>kk`)
	//xs, err := Parse(`<a`)
	xs, err := Parse(logger.DefaultLogger().Explain, `
	select $a weight < / > =,
		'a', "b"
		${ weight[1].Value }
		// hi
		<if test="price <= 0">
			and price &lt; #{ price[0].Age }
	
		<where>where2</where>
		<foreach>
			foreach3
		</foreach>
	
		</if>
	
		kkk
	`, nil)
	//xs, err := Parse(`
	//	<a test="123"><b>c</b></a>
	//`)
	require.NoError(t, err)
	//t.Log(xs.Raw())
	t.Log(xs.Statement())
}

func TestCalc(t *testing.T) {
	//xs, err := Parse(`
	//	select * from products
	//    <if test="!status">
	//	</if>
	//	<if test="age > 18">
	//	</if>
	//`, map[string]any{
	//	"status": 18,
	//	"age":    18,
	//})
	//require.NoError(t, err)
	//t.Log(xs.Raw())
}

func TestStmt(t *testing.T) {
	//xs, err := Explain(`
	//select * from products where category in #{category} where color = ${color};
	//`, map[string]any{
	//	"category": []string{"a", "b"},
	//	"color":    "red",
	//})
	//require.NoError(t, err)
	//t.Log(xs)
	//spew.Json(xs.Vars())
}

func TestVisitor(t *testing.T) {

	type TestCase struct {
		Vars       map[string]any
		Error      bool
		ExpectSQL  string
		ExpectVars []any
	}

	type TestSource struct {
		Source string
		Cases  []TestCase
	}

	tests := []TestSource{
		// if
		//{
		//	Source: `select <if test="a > 0">${b}</if>`,
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"a": 1, "b": "b"}, Error: false, ExpectSQL: "select 'b'", ExpectVars: nil},
		//		{Vars: map[string]any{"a": 0}, Error: false, ExpectSQL: "select ", ExpectVars: nil},
		//	},
		//},
		//{
		//	Source: `<if test="a==-1">-1</if>`,
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"a": -1}, Error: false, ExpectSQL: "-1", ExpectVars: nil},
		//	},
		//},
		//
		//// choose
		//{
		//	Source: `
		//		select <choose>
		//				<when test="a >= 1">1</when>
		//				<when test="a >= 0">0</when>
		//				<otherwise>-1</otherwise>
		//               </choose>
		//	`,
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"a": 1}, Error: false, ExpectSQL: " select 1 ", ExpectVars: nil},
		//		{Vars: map[string]any{"a": 0}, Error: false, ExpectSQL: " select 0 ", ExpectVars: nil},
		//		{Vars: map[string]any{"a": -1}, Error: false, ExpectSQL: " select -1 ", ExpectVars: nil},
		//	},
		//},

		// foreach
		//{
		//	Source: `
		//		select * from products where "id" in
		//		<foreach item="id" collection="ids" open="(" close=")" separator=",">
		//			#{id}
		//		</foreach>
		//	`,
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"ids": []int{1, 2, 3}}},
		//	},
		//},

		// set
		//{
		//	Source: `
		//	  UPDATE some_table
		//	  <set>
		//		<if test="name != nil">name = #{name},</if>
		//		<if test="age != nil">age = #{age},</if>
		//		<if test="address != nil">address = #{address},</if>
		//	  </set>
		//	  WHERE id = #{id}
		//	`,
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"name": "tom", "age": nil, "address": nil, "id": 1}, ExpectSQL: "", ExpectVars: []any{"tom"}},
		//	},
		//},

		// where
		//{
		//	Source: `
		//	SELECT * products
		//	<where>
		//		<if test="a != nil">field1 = #{a}</if>
		//		<if test="b != nil">AnD field2 = #{b}</if>
		//		<if test="c != nil">oR field3 = #{c}</if>
		//	</where>
		//    `,
		//	Cases: []TestCase{
		//		//{Vars: map[string]any{"a": 1, "b": nil, "c": nil}},
		//		//{Vars: map[string]any{"a": nil, "b": 2, "c": nil}},
		//		{Vars: map[string]any{"a": nil, "b": nil, "c": 3}},
		//	},
		//},

		// trim
		{
			Source: `
			select * from table_name 
			<trim prefix="where" prefixOverrides="and |or ">
				<if test="a != nil">
					title = #{a}
				</if>
				<if test="b != nil">
					and content = #{b}
				</if>
				<if test="c != nil">
					or owner = #{c}
				</if>
			</trim>
            `,
			Cases: []TestCase{
				//{Vars: map[string]any{"a": 1, "b": nil, "c": nil}},
				{Vars: map[string]any{"a": nil, "b": 2, "c": nil}},
				//{Vars: map[string]any{"a": nil, "b": nil, "c": 3}},
			},
		},

		// complex
		//{
		//	Source: "select * from products where category in #{category}",
		//	Cases: []TestCase{
		//		{Vars: map[string]any{"category": 1}, Error: false, ExpectSQL: "select * from products where category in $1", ExpectVars: []any{1}},
		//		{Vars: map[string]any{"category": "a"}, Error: false, ExpectSQL: "select * from products where category in $1", ExpectVars: []any{"a"}},
		//		{Vars: map[string]any{"category": []int{1, 2, 3}}, Error: false, ExpectSQL: "select * from products where category in ($1,$2,$3)", ExpectVars: []any{1, 2, 3}},
		//	},
		//},
	}

	for _, v := range tests {
		for _, vv := range v.Cases {
			r, err := Parse(logger.DefaultLogger().Explain, v.Source, vv.Vars)
			if vv.Error {
				require.Error(t, err, v.Source)
			} else {
				require.NoError(t, err, v.Source)
				require.Equal(t, vv.ExpectSQL, r.Statement())
				require.Equal(t, len(vv.ExpectVars), len(r.Vars()), v.Source)
				for i, vvv := range vv.ExpectVars {
					require.Equal(t, vvv, r.Vars()[i], v.Source)
				}
			}
		}
	}
}

func TestRegex(t *testing.T) {

	r, err := regexp.Compile(`(and | or )`)
	require.NoError(t, err)

	t.Log(r.MatchString("a and b"))
	t.Log(r.MatchString("a or c"))
	t.Log(r.MatchString("and"))

}
