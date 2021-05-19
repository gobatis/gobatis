package gobatis

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
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
	Name     string
	Age      int
	Weight   func() int
	auth     bool
	Children map[string]int64
	Products map[int][]int
	Parent   *testEntity
}

type testExpression struct {
	in        []interface{}
	parameter string
	expr      string
	result    interface{}
	err       int
}

type testMapper struct {
	file    string
	content string
	err     int
}

type testFragment struct {
	id        string
	parameter []interface{}
	sql       string
	vars      int
	err       int
}

var (
	u1 testEntity
	u2 testEntity
)

func init() {
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
		{id: "QueryTestByStatues", parameter: []interface{}{"ok", "success"}, sql: "", vars: 0},
	})
}

func TestErrorParseMapper(t *testing.T) {
	engine := NewEngine(&DB{})
	execTestErrorMapper(t, engine, []testMapper{
		//{err: syntaxErr, file: "mapper.xml", content: `<mapper>...</mapper`},
		{err: syntaxErr, file: "mapper.xml", content: `<mapper</mapper`},
		//{err: syntaxErr, file: "mapper.xml", content: `mapper>...</mapper`},
		//{err: syntaxErr, file: "mapper.xml", content: `mapper>.../mapper>`},
	})
}

func TestCorrectParseExprExpression(t *testing.T) {
	
	testParseExprExpression(t, []testExpression{
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "a + b", result: 6},
		{in: []interface{}{2, 4}, parameter: "a:int, b", expr: "a + b", result: 6},
		{in: []interface{}{2, 4}, parameter: "a:int, b:int", expr: "a + b", result: 6},
		{in: []interface{}{2, 4}, parameter: "a, b:int", expr: "a + b", result: 6},
		{in: []interface{}{int8(2), int8(4)}, parameter: "a,b", expr: "a - b", result: -int8(2)},
		{in: []interface{}{int16(2), int16(4)}, parameter: "a,b", expr: "a * b", result: int16(8)},
		{in: []interface{}{int32(2), int32(4)}, parameter: "a,b", expr: "a / b ", result: int32(0)},
		{in: []interface{}{int64(2), int64(4)}, parameter: "a,b", expr: "b / a", result: int64(2)},
		{in: []interface{}{decimal.NewFromFloat(3.12), "2.13"}, parameter: "a,b", expr: "a + b", result: "5.25"},
		{in: []interface{}{decimal.NewFromFloat(3.12), 2.13}, parameter: "a,b", expr: "a + b", result: "5.25"},
		
		{in: []interface{}{2, 4}, parameter: "a,b", expr: " a + a * a", result: 6},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "  a + a * b", result: 10},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: " a + b * a ", result: 10},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "a + b * b", result: 18},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "(a + b) * b", result: 24},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( ( a + b  ) * b)", result: 24},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( ( (  a + b ) ) * b)", result: 24},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( a + b) / b", result: 1},
		
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "b + b * b", result: 20},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "b + b * a ", result: 12},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "b + a * b", result: 12},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "b + a * a", result: 8},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( b + a ) * a", result: 12},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( ( b+ a ) * a)", result: 12},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( ( ( b + a ) * a ) )", result: 12},
		{in: []interface{}{2, 4}, parameter: "a,b", expr: "( b + a ) / a", result: 3},
		
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "a.Age + b.Age", result: 38},
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "a.Weight() + b.Weight()", result: 100},
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "a.Weight() > b.Weight()", result: true},
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "a.Parent == nil", result: true},
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "nil == a.Parent", result: true},
		{in: []interface{}{u1, u2}, parameter: "a, b", expr: "nil == a.Parent && nil != b.Parent", result: true},
		{in: []interface{}{u1, u2}, parameter: "parent1, child2", expr: "child2.Parent.Age + child2.Age", result: 38},
	})
}

func TestErrorParseExprExpression(t *testing.T) {
	testParseExprExpression(t, []testExpression{
		{in: []interface{}{2, 4}, parameter: "a:int32, b", expr: "a + b", result: 6, err: 1},
	})
	
	//require.IsType(t, )
}

func execTestErrorMapper(t *testing.T, engine *Engine, tests []testMapper) {
	for _, test := range tests {
		_ = parseMapper(engine, test.file, test.content)
		//require.Error(t, err)
		//_err, ok := err.(*_error)
		//require.True(t, ok, "expected *_error")
		//require.Equal(t, test.err, _err.code, err)
	}
}

func execTestFragment(t *testing.T, engine *Engine, tests []testFragment) {
	reg := regexp.MustCompile(`\s+`)
	for _, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.parameter {
			vars = append(vars, rv(v))
		}
		
		frag, ok := engine.fragmentManager.get(test.id)
		require.True(t, ok, test)
		sql, args, err := frag.parseStatement(vars...)
		require.NoError(t, err)
		
		if test.err > 0 {
			require.Error(t, err, test)
		} else {
			require.NoError(t, err, test)
			require.Equal(t, reg.ReplaceAllString(sql, ""), reg.ReplaceAllString(test.sql, ""), test)
			require.Equal(t, len(args), test.vars, test)
		}
	}
}

func testParseExprExpression(t *testing.T, tests []testExpression) {
	for _, test := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range test.in {
			vars = append(vars, rv(v))
		}
		parser := newExprParser(vars...)
		parser.file = "tmp.xml"
		err := parser.parseParameter(test.parameter)
		require.NoError(t, err)
		result, err := parser.parseExpression(test.expr)
		if test.err > 0 {
			require.Error(t, err, test)
		} else {
			require.NoError(t, err, test)
			dr, ok := result.(decimal.Decimal)
			if ok {
				require.Equal(t, test.result, dr.String(), test)
			} else {
				require.Equal(t, test.result, result, test)
			}
		}
	}
}
