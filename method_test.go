package gobatis

import (
	"fmt"
	"github.com/shopspring/decimal"
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

func appendCases(v ...[]testParseMapperCase) []testParseMapperCase {
	items := make([]testParseMapperCase, 0)
	for _, vv := range v {
		items = append(items, vv...)
	}
	return items
}

type testParseMapperCase struct {
	definition string
	method     reflect.Value
	sqls       []*testParseMapperCaseSql
	error      bool
}

type testParseMapperCaseSql struct {
	in      []reflect.Value
	stmtSQL string
	realSQL string
	values  []interface{}
	error   bool
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
				in:      []reflect.Value{rv(1)},
				stmtSQL: `select * from users where id = $1;`,
				realSQL: `select * from users where id = 1;`,
				values:  []interface{}{1},
			},
		},
	},
	{
		definition: `
		<select id="SelectP002" parameter="userId,status">
			select * from orders where user_id = #{userId}
			<if test="status>0">
				and status = #{ status }
			</if>
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(1), rv(0)},
				stmtSQL: `select * from orders where user_id = $1`,
				realSQL: `select * from orders where user_id = 1`,
				values:  []interface{}{1},
			},
			{
				in:      []reflect.Value{rv(1), rv(2)},
				stmtSQL: `select * from orders where user_id = $1 and status = $2`,
				realSQL: `select * from orders where user_id = 1 and status = 2`,
				values:  []interface{}{1, 2},
			},
		},
	},
	{
		definition: `
		<select id="SelectP003" parameter="userId,status">
			select * from orders
			<where>
				<if test="userId>0">
					user_id = #{userId}
				</if>
				<if test="status>0">
					and status = #{ status }
				</if>
			</where>
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(0), rv(0)},
				stmtSQL: `select * from orders`,
				realSQL: `select * from orders`,
				values:  []interface{}{},
			},
			{
				in:      []reflect.Value{rv(1), rv(0)},
				stmtSQL: `select * from orders where user_id = $1`,
				realSQL: `select * from orders where user_id = 1`,
				values:  []interface{}{1},
			},
			{
				in:      []reflect.Value{rv(0), rv(1)},
				stmtSQL: `select * from orders where status = $1`,
				realSQL: `select * from orders where status = 1`,
				values:  []interface{}{1},
			},
		},
	},
	{
		definition: `
		<select id="SelectP004" parameter="id">
			<bind name="table" value="'users'"  />
			select * from ${table} where id = #{id};
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(1)},
				stmtSQL: `select * from users where id = $1;`,
				realSQL: `select * from users where id = 1;`,
				values:  []interface{}{1},
			},
		},
	},
}

var testParseInsertCases = []testParseMapperCase{
	{
		definition: `
		<insert id="InserterP001" parameter="row">
			<inserter table="'users'">
				<field name="'name'">#{row.Name}</field>
				<field name="'age'">#{row.Age}</field>
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
				stmtSQL: `insert into users("name","age") values($1,$2)`,
				realSQL: `insert into users("name","age") values('tom',18)`,
				values:  []interface{}{"tom", 18},
			},
		},
	},
	{
		definition: `
		<insert id="InserterP001" parameter="row">
			<inserter table="'users'" entity="row">
				<field name="*"/>
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
				stmtSQL: `insert into users("name","age") values($1,$2)`,
				realSQL: `insert into users("name","age") values('tom',18)`,
				values:  []interface{}{"tom", 18},
			},
		},
	},
}

var testParseUpdateCases = []testParseMapperCase{
	{
		definition: `
		<update id="UpdateP001" parameter="id, password">
			update users set password = #{password} where id = #{id};
		</update>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(1), rv("123456")},
				stmtSQL: `update users set password = $1 where id = $2;`,
				realSQL: `update users set password = '123456' where id = 1;`,
				values:  []interface{}{"123456", 1},
			},
		},
	},
}

var testParseDeleteCases = []testParseMapperCase{
	{
		definition: `
		<delete id="DeleteP001" parameter="id">
			delete from users where id=#{id};
		</delete>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(1)},
				stmtSQL: `delete from users where id=$1;`,
				realSQL: `delete from users where id=1;`,
				values:  []interface{}{1},
			},
		},
	},
}

func (p testParseMapperCaseSql) check(t *testing.T, c *testParseMapperCase, s *segment) bool {
	require.Equal(t, p.stmtSQL, s.sql, c)
	require.Equal(t, p.realSQL, s.realSQL(), c)
	require.Equal(t, len(p.values), len(s.vars), c)
	for i, v := range s.vars {
		require.Equal(t, p.values[i], v, c)
	}
	return false
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
	
	for _, item := range appendCases(
		testParseSelectCases,
		testParseInsertCases,
		testParseUpdateCases,
		testParseDeleteCases,
	) {
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
			sql.check(t, &item, s)
			t.Log(fmt.Sprintf("[%s]", f.id), s.realSQL())
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
				stmtSQL: `insert into users("name","age") values(?,?)`,
			},
		},
		error: false,
	},
	{
		definition: `
		<query id="TestQueryP002" parameter="age,page,limit">
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
					rv(decimal.NewFromFloat(3.14)),
					rv(2),
					rv(10),
				},
				stmtSQL: `insert into users("name","age") values(?,?)`,
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
			t.Log(f.id, ss[0].realSQL())
			//t.Log(f.id, ss[0].sql)
			//d, _ := json.Marshal(ss[0].vars)
			//t.Log(f.id, string(d))
			//
			t.Log(f.id, ss[1].realSQL())
			//d, _ = json.Marshal(ss[1].vars)
			//t.Log(f.id, string(d))
		}
	}
}
