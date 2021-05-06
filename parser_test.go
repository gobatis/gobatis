package gobatis

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
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

const defaultUserMapper = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "../../../dtd/mapper.dtd">

<mapper namespace="test/mapper@UserMapper">
    <select id="SelectTestById" parameterType="id:int64,name:string,user:struct" resultType="int64">
        select * from test where id = #{1}
    </select>
</mapper>
`

func TestParseConfig(t *testing.T) {
	engine := NewEngine(NewDB("nil", "nil"))
	require.NoError(t, parseConfig(engine, "gobatis.xml", defaultConfigXML))
}

func TestParseMapper(t *testing.T) {
	//engine := NewEngine(NewDB("nil", "nil"))
	//require.NoError(t, parseMapper(engine, "user.Mapper.xml", defaultUserMapper))
	//d, _ := json.MarshalIndent(engine.statements, "", "\t")
	//fmt.Println(string(d))
}

func TestParseExprExpression(t *testing.T) {
	//parser := newExprParser(1, 2)
	//result, err := parser.parseExpression("a, b", "a+b")
	//parser := newExprParser(10, 20, 30)
	//result, err := parser.parseExpression("a, b, c", "a+b*((c+1))")
	//parser := newExprParser(3,"2")
	//result, err := parser.parseExpression("a, b", "a + b")
	a := 1
	b := 2
	parser := newExprParser(a, b)
	err := parser.parseParameter("a,b")
	require.NoError(t, err)
	result, err := parser.parseExpression("a + b")
	require.NoError(t, err)
	require.Equal(t, int64(3), result)
}

type testStruct struct {
	Name string
	Age  int64
	Map  map[string]int64
	Calc func(val int64) int64
	Max  func(a, b int64) int64
	Dec  func(a decimal.Decimal) string
}

func TestParseExprExpressionMember(t *testing.T) {
	//parser := newExprParser(1, 2)
	//result, err := parser.parseExpression("a, b", "a+b")
	//parser := newExprParser(10, 20, 30)
	//result, err := parser.parseExpression("a, b, c", "a+b*((c+1))")
	//parser := newExprParser(3,"2")
	//result, err := parser.parseExpression("a, b", "a + b")
	a := testStruct{
		Name: "gobatis",
		Age:  64,
		Map: map[string]int64{
			"weight": 60,
		},
		Calc: func(val int64) int64 {
			fmt.Println("this val is :", val)
			return val / 3
		},
		Max: func(a, b int64) int64 {
			if a > b {
				return a
			}
			return b
		},
		Dec: func(a decimal.Decimal) string {
			return a.Add(decimal.NewFromFloat(1.1234)).String()
		},
	}
	b := []int{1, 2, 3, 4, 5}
	parser := newExprParser(a, b)
	//result, err := parser.parseExpression("a:struct, b:array", `a.Max(a.Age, int64(b[2]) + a.Map["weight"])`)
	//result, err := parser.parseExpression("a:struct, b:array", `strings.HasPrefix(a.Name, "go")`)
	//result, err := parser.parseExpression("a:struct, b:array", `b[0:len(b)]`)
	//result, err := parser.parseExpression("a:struct, b:array", `a.Age > int64(1) && b[2] > int64(1)`)
	//result, err := parser.parseExpression("a:struct, b:array", `a.Age > 1 && b[2] > 1`)
	err := parser.parseParameter("a:struct, b:array")
	require.NoError(t, err)
	result, err := parser.parseExpression(`a.Dec(100)`)
	require.NoError(t, err)
	t.Log("result:", result)
}

func TestBindParser(t *testing.T) {

	//m := func(tx *sql.Tx, a, b string) (_a int, _b bool, err error) { return }
	//f := realReflectType(m)

	var err error
	_, err = parseFragment("", "selectUser", "a:string, b:string", "a:int, b:bool", nil)
	require.NoError(t, err)

	//err = parseFragment("", f, "a, b:string", "a:int, b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a, b", "a:int, b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a, b", "a, b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a, b", "a, b")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a:string,b:string", "a:int,b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a,b:string", "a:int,b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a,b", "a:int,b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a,b", "a,b:bool")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a,b", "a,b")
	//require.NoError(t, err)
	//
	//err = parseFragment("", f, "a:int, b:string", "a:string, b:string")
	//require.Error(t, err)
	//
	//err = parseFragment("", f, "a:string, b:string", "a:string, b:string")
	//require.Error(t, err)
	//
	//err = parseFragment("", f, "a:string, b:string, c", "a:string, b:string")
	//require.Error(t, err)
	//
	//err = parseFragment("", f, "a:string, b:string", "a:int, b:string, c")
	//require.Error(t, err)
}
