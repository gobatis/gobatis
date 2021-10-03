package gobatis

import (
	"encoding/json"
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

var testParseSelectCases = []testParseMapperCase{
	{
		definition: `
		<select id="SelectP001" parameter="id">
			select * from users where id = #{id};
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:  []reflect.Value{rv(100)},
				sql: `insert into users("name","age") values(?,?)`,
			},
		},
		error: false,
	},
	{
		definition: `
		<insert id="TestInserter1" parameter="row">
			<inserter table="'users'" entity="row">
				<field name="*" />
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
func TestParseSelectCases(t *testing.T) {
	var (
		fs  []*method
		f   *method
		c   *caller
		s   *segment
		err error
	)
	for _, item := range testParseSelectCases {
		fs, err = parseMapper("", wrapMapperSchema(item.definition))
		if item.error {
			require.Equal(t, true, err != nil)
			continue
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, 1, len(fs), item.definition)
		f = fs[0]
		for _, sql := range item.sqls {
			c = f.newCaller(nil)
			s, err = c.method.buildSegment(sql.in)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			t.Log(f.id, s.sql)
			d, _ := json.Marshal(s.vars)
			t.Log(f.id, string(d))
		}
	}
}

var testParseQueryCases = []testParseMapperCase{
	{
		definition: `
		<query id="TestQueryP001" parameter="age,page,limit">
			<block type="COUNT">
				select count(1)
			</block>
			<block type="SELECT">
				select username,email,status
			</block>
			<block type="FROM">
				from users where age >= #{age}
			</block>
			<block type="LIMIT">
				order by age desc limit #{limit} offset #{ paging(page,limit)}
			</block>
		</query>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{
					rv(18),
					rv(2),
					rv(10),
				},
				sql: `insert into users("name","age") values(?,?)`,
			},
		},
		error: false,
	},
	{
		definition: `
		<query id="TestQueryP001" parameter="age,page,limit">
			<block type="COUNT">
				select count(1)
			</block>
			<block type="SELECT">
				select username,email,status
			</block>
			<block type="FROM">
				from users
				<where>
				age &lt;= #{age}
				</where>
			</block>
			<block type="LIMIT">
				order by age desc
				limit #{limit} offset #{ paging(page,limit)}
			</block>
		</query>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{
					rv(18),
					rv(2),
					rv(10),
				},
				sql: `insert into users("name","age") values(?,?)`,
			},
		},
		error: false,
	},
}

// test parse mappers
func TestParseQueryCases(t *testing.T) {
	var (
		fs []*method
		f  *method
		c  *caller
		//s  *segment
		ss []*segment
		//parser *exprParser
		err error
	)
	for _, item := range testParseQueryCases {
		fs, err = parseMapper("", wrapMapperSchema(item.definition))
		if item.error {
			require.Equal(t, true, err != nil)
			continue
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, 1, len(fs), item.definition)
		f = fs[0]
		for _, sql := range item.sqls {
			c = f.newCaller(nil)
			ss, err = c.method.buildQuery(sql.in)
			//parser, err = c.method.prepareParser(s.in)
			require.NoError(t, err)
			//err = c.method.buildSegment(parser, s, c.method.node)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			t.Log(f.id, ss[0].sql)
			d, _ := json.Marshal(ss[0].vars)
			t.Log(f.id, string(d))
			
			t.Log(f.id, ss[1].sql)
			d, _ = json.Marshal(ss[1].vars)
			t.Log(f.id, string(d))
		}
	}
}
