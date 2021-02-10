package compiler

import (
	"fmt"
	"testing"
)

const normalConfig = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
    <properties>
        <property name="module" value="github.com/koyeo/gobatis"/>
    </properties>
    <typeAliases>
        <typeAlias alias="User" type="test/entity@User"/>
    </typeAliases>
    <environments default="development">
        <environment id="development">
            <transactionManager type="JDBC"/>
            <dataSource type="POOLED">
                <property name="driver" value="mysql"/>
                <property name="url" value="localhost:3306"/>
                <property name="username" value="root"/>
                <property name="password" value="123456"/>
                <property name="database" value="antq"/>
            </dataSource>
        </environment>
    </environments>
    <mappers>
        <mapper resource="test/sql/user.mapper.xml"/>
    </mappers>
</configuration>
`

const normalMapper = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="test/mapper@UserMapper">
    <cache blocking="123">
    </cache>
    <cache blocking="123">
    </cache>
    <select id="GetUser" parameterType="int64" resultType="test/entity@User">
        select * from users where id = #{id}
        <where>
            <if test="a != nil">
                a
                <if test="a!=nil">ok</if>
                <where>
                </where>
            </if>
        </where>
    </select>
    <!--    <select id="selectUsers" parameterType="id:int64,name:string" resultMap="Blog[]">-->
    <!--        select * from blog-->
    <!--    </select>-->
    <!--    <select id="countUsers" parameterType="id:int64,name:string" resultType="int,int,string">-->
    <!--        select * from users where id=#{id} and name=#{name}-->
    <!--    </select>-->
</mapper>
`

const ignoreStatement = `
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!-- just some comment -->
<mapper namespace="test/mapper@UserMapper">
	<select id="123"></select>
</mapper>
`

func TestNewXMLTokenizer2(t *testing.T) {
	_, err := NewXMLTokenizer([]byte(normalConfig)).Parse()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewXMLTokenizer(t *testing.T) {
	_, err := NewXMLTokenizer([]byte(normalMapper)).Parse()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIgnoreStatement(t *testing.T) {
	tokens, err := NewXMLTokenizer([]byte(ignoreStatement)).Parse()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range tokens {
		fmt.Println(v.String())
	}
	//assert.Equal(t, len(tokens), 8, "tokens length should be 8")
	//assert.Equal(t, tokens[0].Value, "mapper")
	//assert.Equal(t, tokens[1].Value, "namespace")
	//assert.Equal(t, tokens[2].Value, "test/mapper@UserMapper")
	//assert.Equal(t, tokens[3].Value, "select")
	//assert.Equal(t, tokens[4].Value, "id")
	//assert.Equal(t, tokens[5].Value, "123")
	//assert.Equal(t, tokens[6].Value, "select")
	//assert.Equal(t, tokens[7].Value, "mapper")
}

func TestNewXMLParser(t *testing.T) {
	_, err := NewXMLParser().Parse([]byte(normalMapper))
	if err != nil {
		t.Error(err)
		return
	}
}
