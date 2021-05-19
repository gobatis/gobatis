package gobatis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/parser/expr"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

const defaultConfigXML = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "../dtd/config.dtd">

<configuration>
    <module base="github.com/gobatis/gobatis"/>
<!--    <properties>-->
    <!--        <property name="module" value="github.com/gobatis/gobatis"/>-->
    <!--        <property name="null2zero" value="true"/>-->
    <!--    </properties>-->
    <typeAliases>
        <typeAlias alias="User" type="test/entity@User"/>
    </typeAliases>
<!--    <environments default="development">-->
<!--        <environment id="development">-->
<!--            <transactionManager type="JDBC"/>-->
<!--            <dataSource type="POOLED">-->
<!--                <property name="driver" value="mysql"/>-->
<!--                <property name="url" value="localhost:3306"/>-->
<!--                <property name="username" value="root"/>-->
<!--                <property name="password" value="123456"/>-->
<!--                <property name="database" value="antq"/>-->
<!--            </dataSource>-->
<!--        </environment>-->
<!--        <environment id="development2">-->
<!--            <dataSource type="POOLED">-->
<!--                <property name="driver" value="mysql"/>-->
<!--                <property name="url" value="localhost:3306"/>-->
<!--                <property name="username" value="root"/>-->
<!--                <property name="password" value="123456"/>-->
<!--                <property name="database" value="antq"/>-->
<!--            </dataSource>-->
<!--        </environment>-->
<!--    </environments>-->
<!--    <mappers>-->
<!--        <mapper resource="test/sql/user.mapper.xml"/>-->
<!--    </mappers>-->
</configuration>
`

const defaultCorrectTestMapper = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "../../../dtd/mapper.dtd">
<mapper namespace="test/mapper@UserMapper">
    <select id="SelectTestById" parameter="id:int64, name:string">
        select * from test where id = #{ id } and name = '${ name }' and age= #{id+2}
    </select>

	<select id="QueryTestByStatues" parameter="statuses:array">
		select * from test where status in
		<foreach item="item" index="index" collection="statuses" open="(" separator="," close=")">
		        '#{item}'
		  </foreach>
		and name > 1 and
		names in <foreach item="item" index="index" collection="statuses" open="(" separator="," close=")">
		        '#{item}'
		  </foreach>
	</select>
</mapper>
`

type testEntity struct {
	Name     string           `json:"name"`
	Age      int              `json:"age"`
	Weight   func() int       `json:"weight"`
	Children map[string]int64 `json:"children"`
	Products map[int][]int    `json:"products"`
	Parent   *testEntity      `json:"parent"`
	auth     bool
}

type testExpression struct {
	In        []interface{} `json:"in"`
	Parameter string        `json:"parameter"`
	Expr      string        `json:"expr"`
	Result    interface{}   `json:"result"`
	Err       int           `json:"err"`
}

type testMapper struct {
	File    string `json:"file"`
	Content string `json:"content"`
	Err     int    `json:"err"`
}

type testFragment struct {
	Id        string        `json:"id"`
	Parameter []interface{} `json:"parameter"`
	SQL       string        `json:"sql"`
	Vars      int           `json:"vars"`
	Err       int           `json:"err"`
}

var (
	u1 testEntity
	u2 testEntity
)

const (
	errLogFile = "err.md"
)

func init() {
	
	logPath := filepath.Join(pwd, errLogFile)
	_, err := os.Stat(logPath)
	if !os.IsNotExist(err) {
		_ = os.Remove(logPath)
	}
	
	u1 = testEntity{
		Name:     "foo",
		Age:      18,
		Weight:   func() int { return 60 },
		auth:     true,
		Children: map[string]int64{"michael": 8},
		Products: map[int][]int{1: []int{11, 12, 13}},
	}
	u2 = testEntity{
		Name:     "foo parent",
		Age:      20,
		Weight:   func() int { return 40 },
		auth:     true,
		Children: map[string]int64{"alice": 8},
		Products: map[int][]int{2: []int{21, 22, 23}},
		Parent:   &u1,
	}
}

func TestParseConfig(t *testing.T) {
	engine := NewEngine(NewDB("nil", "nil"))
	require.NoError(t, parseConfig(engine, "gobatis.xml", defaultConfigXML))
}

func TestCorrectParseFragment(t *testing.T) {
	engine := NewEngine(&DB{})
	err := parseMapper(engine, "defaultCorrectTestMapper", defaultCorrectTestMapper)
	require.NoError(t, err)
	
	execTestFragment(t, engine, []testFragment{
		{Id: "QueryTestByStatues", Parameter: []interface{}{"ok", "success"}, SQL: "", Vars: 0},
	})
}

func TestErrorParseMapper(t *testing.T) {
	engine := NewEngine(&DB{})
	execTestErrorMapper(t, engine, []testMapper{
		{Err: syntaxErr, File: "mapper.xml", Content: `<mapper>...</mapper`},
		{Err: syntaxErr, File: "mapper.xml", Content: `<mapper</mapper`},
		{Err: syntaxErr, File: "mapper.xml", Content: `mapper>...</mapper`},
		{Err: syntaxErr, File: "mapper.xml", Content: `mapper>.../mapper>`},
	})
}

func TestCorrectParseExprExpression(t *testing.T) {
	
	testParseExprExpression(t, []testExpression{
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a:int, b", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a:int, b:int", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a, b:int", Expr: "a + b", Result: 6},
		{In: []interface{}{int8(2), int8(4)}, Parameter: "a,b", Expr: "a - b", Result: -int8(2)},
		{In: []interface{}{int16(2), int16(4)}, Parameter: "a,b", Expr: "a * b", Result: int16(8)},
		{In: []interface{}{int32(2), int32(4)}, Parameter: "a,b", Expr: "a / b ", Result: int32(0)},
		{In: []interface{}{int64(2), int64(4)}, Parameter: "a,b", Expr: "b / a", Result: int64(2)},
		{In: []interface{}{decimal.NewFromFloat(3.12), "2.13"}, Parameter: "a,b", Expr: "a + b", Result: "5.25"},
		{In: []interface{}{decimal.NewFromFloat(3.12), 2.13}, Parameter: "a,b", Expr: "a + b", Result: "5.25"},
		
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: " a + a * a", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "  a + a * b", Result: 10},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: " a + b * a ", Result: 10},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "a + b * b", Result: 18},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "(a + b) * b", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( a + b  ) * b)", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( (  a + b ) ) * b)", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( a + b) / b", Result: 1},
		
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + b * b", Result: 20},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + b * a ", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + a * b", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + a * a", Result: 8},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( b + a ) * a", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( b+ a ) * a)", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( ( b + a ) * a ) )", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( b + a ) / a", Result: 3},
		
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Age + b.Age", Result: 38},
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() + b.Weight()", Result: 100},
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() > b.Weight()", Result: true},
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Parent == nil", Result: true},
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent", Result: true},
		{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent && nil != b.Parent", Result: true},
		{In: []interface{}{u1, u2}, Parameter: "parent1, child2", Expr: "child2.Parent.Age + child2.Age", Result: 38},
	})
}

func TestErrorParseExprExpression(t *testing.T) {
	testParseExprExpression(t, []testExpression{
		{In: []interface{}{2, 4}, Parameter: "a, b", Expr: "a + b", Result: 6, Err: 0},
	})
}

func execTestErrorMapper(t *testing.T, engine *Engine, tests []testMapper) {
	for i, test := range tests {
		err := parseMapper(engine, test.File, test.Content)
		require.Error(t, err)
		writeError(t, fmt.Sprintf("test error mapper: %d", i), test, err)
		_err, ok := err.(*_error)
		require.True(t, ok, "expected *_error")
		require.Equal(t, test.Err, _err.code, err)
	}
}

func writeError(t *testing.T, title string, test interface{}, _err error) {
	f, err := os.OpenFile(filepath.Join(pwd, errLogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	require.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()
	td := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(td)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", " ")
	err = jsonEncoder.Encode(test)
	require.NoError(t, err)
	//td, err := json.MarshalIndent(test, "", "")
	require.NoError(t, err)
	_, err = f.WriteString(fmt.Sprintf(
		"**%s**\n\ndata:\n```\n%s```\nerror:\n```\n%s```\n",
		title, td.String(), _err.Error(),
	))
	require.NoError(t, err)
}

func execTestFragment(t *testing.T, engine *Engine, tests []testFragment) {
	reg := regexp.MustCompile(`\s+`)
	for _, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.Parameter {
			vars = append(vars, rv(v))
		}
		
		frag, ok := engine.fragmentManager.get(test.Id)
		require.True(t, ok, test)
		sql, args, err := frag.parseStatement(vars...)
		require.NoError(t, err)
		
		if test.Err > 0 {
			require.Error(t, err, test)
		} else {
			require.NoError(t, err, test)
			require.Equal(t, reg.ReplaceAllString(sql, ""), reg.ReplaceAllString(test.SQL, ""), test)
			require.Equal(t, len(args), test.Vars, test)
		}
	}
}

func newTestNodeCtx() antlr.ParserRuleContext {
	nodeCtx := expr.NewEmptyExpressionContext()
	nodeCtx.SetStart(antlr.NewCommonToken(&antlr.TokenSourceCharStreamPair{}, 0, 0, 10, 10))
	return nodeCtx
}

func testParseExprExpression(t *testing.T, tests []testExpression) {
	nodeCtx := newTestNodeCtx()
	for i, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.In {
			vars = append(vars, rv(v))
		}
		parser := newExprParser(vars...)
		parser.file = "tmp.xml"
		err := parser.parseParameter(nodeCtx, test.Parameter)
		require.NoError(t, err)
		result, err := parser.parseExpression(nodeCtx, test.Expr)
		if test.Err > 0 {
			require.Error(t, err, test)
			writeError(t, fmt.Sprintf("test parse expression: %d", i), test, err)
		} else {
			require.NoError(t, err, test)
			dr, ok := result.(decimal.Decimal)
			if ok {
				require.Equal(t, test.Result, dr.String(), test)
			} else {
				require.Equal(t, test.Result, result, test)
			}
		}
	}
}
