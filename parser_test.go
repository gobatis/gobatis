package gobatis

import (
	"encoding/json"
	"fmt"
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
    <!--        <property name="module" value="github.com/koyeo/gobatis"/>-->
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
	engine := NewEngine(NewDB("nil", "nil"))
	require.NoError(t, parseMapper(engine, "user.Mapper.xml", defaultUserMapper))
	d, _ := json.MarshalIndent(engine.statements, "", "\t")
	fmt.Println(string(d))
}

func TestParseExprExpression(t *testing.T) {
	parser := newExprParser(10,20,30)
	result, err := parser.parseExpression("a:int, b:int", "a+b*c")
	require.NoError(t, err)
	//require.Equal(t, int64(2), result)
	t.Log("result:", result)
	//
}
