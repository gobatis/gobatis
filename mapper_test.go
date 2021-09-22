package gobatis

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func wrapConfigureSchema(s string) string {
	return fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>"+
		"<!DOCTYPE configuration PUBLIC \"-//gobatis.co//DTD Config 1.0//EN\" \"gobaits.co/dtd/config.dtd\">"+
		"<configuration>%s</configuration>", s)
}

func wrapMapperSchema(s string) string {
	return fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>"+
		"<!DOCTYPE mapper PUBLIC \"-//gobatis.co//DTD Mapper 1.0//EN\" \"gobatis.co/dtd/mapper.dtd\">"+
		"<mapper>%s</mapper>", s)
}

type testParseMapperCase struct {
	definition string
	method     reflect.Value
	sqls       []*testParseMapperCaseSql
	error      bool
}

type testParseMapperCaseSql struct {
	in    []reflect.Value
	sql   string
	error bool
}

var testParseMappersCases = []testParseMapperCase{
	//{
	//	definition: `
	//	<insert id="TestInserter" parameter="row">
	//		insert into users("name", "age") values(#{row.Name}, #{row.Age});
	//	</insert>`,
	//	method: rv(func(row string) (err error) { return }),
	//	sqls: []*testParseMapperCaseSql{
	//		{
	//			in: []reflect.Value{rv(struct {
	//				Name string
	//				Age  int
	//			}{
	//				Name: "tom",
	//				Age:  18,
	//			})},
	//			sql: `insert into users("name","age") values(?,?)`,
	//		},
	//	},
	//	error: false,
	//},
	{
		definition: `
		<insert id="TestInserter" parameter="row">
			<inserter table="'users'" data="row">
				<field name="*"/>
				<field name="'name'">#{row.Name}</field>
				<field name="'age'">${row.Age}</field>
			</inserter>
		</insert>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{rv(struct {
					Name string
					Age  int
				}{
					Name: "tom",
					Age:  18,
				})},
				sql: `insert into users("name","age") values(?,?)`,
			},
		},
		error: false,
	},
}

// test parse mappers
func TestParseMappers(t *testing.T) {
	var (
		fs  []*fragment
		f   *fragment
		s   *sentence
		err error
	)
	for _, c := range testParseMappersCases {
		fs, err = parseMapper("", wrapMapperSchema(c.definition))
		if c.error {
			require.Equal(t, true, err != nil)
			continue
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, 1, len(fs))
		f = fs[0]
		for _, sql := range c.sqls {
			s, err = f.newCaller(nil).prepare(sql.in...)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			t.Log(s.sql)
			t.Log(s.realSql())
		}
	}
	
}
