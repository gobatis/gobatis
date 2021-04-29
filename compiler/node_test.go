package compiler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNodeParser_ParseConfiguration(t *testing.T) {
	parser := NewNodeParser("./config.xml")
	_, err := parser.ParseConfiguration([]byte(`
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-config.dtd">

<configuration>
    <properties>
        <property name="module" value="github.com/gobatis/gobatis"/>
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
`))
	if err != nil {
		t.Error(err)
	}

}

func TestNewNodeParser(t *testing.T) {
	parser := NewNodeParser("./config.xml")
	nodes, err := parser.ParseMapper([]byte(testIf))
	if err != nil {
		t.Error(err)
		return
	}
	d, _ := json.MarshalIndent(nodes, "", "\t")
	fmt.Println(string(d))
}
