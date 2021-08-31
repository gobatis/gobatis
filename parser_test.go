package gobatis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
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

const (
	errLogFile = "err.md"
)

func init() {
	
	logPath := filepath.Join(pwd, errLogFile)
	_, err := os.Stat(logPath)
	if !os.IsNotExist(err) {
		_ = os.Remove(logPath)
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
		{Id: "QueryTestByStatues", Parameter: []interface{}{[]string{"ok", "success"}}, SQL: "select * from test where status in('$1','$2') and name > 1 and names in('$3','$4')", Vars: 4},
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
	
	testCorrectParseExprExpression(t, []testExpression{
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
		
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Age + b.Age", Result: 38},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() + b.Weight()", Result: 100},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() > b.Weight()", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Parent == nil", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent && nil != b.Parent", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "parent1, child2", Expr: "child2.Parent.Age + child2.Age", Result: 38},
	})
}

func TestErrorParseExprExpression(t *testing.T) {
	testErrorParseExprParameter(t, []testExpression{
		//{In: []interface{}{2, 4}, Parameter: `a:int64,b`, Expr: "a + b", Result: 6, Err: 1},
	})
}

func TestAnyExprParam(t *testing.T) {
	params, err := testAnyExprParam(" * ")
	require.Equal(t, 0, len(params))
	require.NoError(t, err)
	
	_, err = testAnyExprParam("*,")
	require.Error(t, err)
	
	_, err = testAnyExprParam("*,a")
	require.Error(t, err)
	
	_, err = testAnyExprParam(",*")
	require.Error(t, err)
}

func testAnyExprParam(tokens string) (params []*param, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	parser := newParamParser("test.xml")
	parser.walkMethods(initExprParser(tokens))
	params = parser.params
	return
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
	tds := td.String()
	if !strings.HasSuffix(tds, "\n") {
		tds += "\n"
	}
	es := _err.Error()
	if !strings.HasSuffix(es, "\n") {
		es += "\n"
	}
	_, err = f.WriteString(fmt.Sprintf("**%s**\n\ndata:\n```\n%s```\nerror:\n```\n%s```\n", title, tds, es))
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
		sql, exprs, _vars, dynamic, err := frag.parseStatement(vars...)
		require.NoError(t, err)
		_ = dynamic
		_ = exprs
		if test.Err > 0 {
			require.Error(t, err, test)
		} else {
			require.NoError(t, err, test)
			require.Equal(t, reg.ReplaceAllString(test.SQL, ""), reg.ReplaceAllString(sql, ""), test)
			require.Equal(t, len(_vars), test.Vars, test)
		}
	}
}

func testCorrectParseExprExpression(t *testing.T, tests []testExpression) {
	for i, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.In {
			vars = append(vars, rv(v))
		}
		_expr := newExprParser(vars...)
		_expr.file = "tmp.xml"
		
		params, err := testParseParams(test.Parameter)
		require.NoError(t, err, test)
		
		for ii, vv := range params {
			err = _expr.paramsStack.list.Front().Next().Value.(*exprParams).bind(vv, ii)
			require.NoError(t, err, test, ii, vv)
		}
		
		result, _, err := _expr.parseExpression(nil, test.Expr)
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

func testErrorParseExprParameter(t *testing.T, tests []testExpression) {
	for i, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.In {
			vars = append(vars, rv(v))
		}
		_, err := testParseParams(test.Parameter)
		require.Error(t, err, test)
		writeError(t, fmt.Sprintf("test error parse parameter: %d", i), test, err)
	}
}

func testParseParams(tokens string) (params []*param, err error) {
	defer func() {
		e := recover()
		err = castRecoverError("", e)
	}()
	params = (&fragment{node: new(xmlNode)}).parseParams(tokens)
	return
}
