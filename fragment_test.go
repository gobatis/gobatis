package gobatis

import (
	"fmt"
	"github.com/gozelle/decimal"
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
		// SelectP001 func(id int64)
		definition: `
		<select id="SelectP001" parameter="id">
			select * from users where id = #{id};
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(1)},
				realSQL: `select * from users where id = 1;`,
				stmtSQL: `select * from users where id = $1;`,
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
	{
		definition: `
		<select id="SelectP005" parameter="status">
			select * from users where
			<choose>
				<when test="status > 1">
					status = #{status}
				</when>
				<when test="status == 1 ">
					status > #{status}
				</when>
				<otherwise>
					status > 0
				</otherwise>
			</choose>
		</select>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in:      []reflect.Value{rv(0)},
				stmtSQL: `select * from users where status > 0`,
				realSQL: `select * from users where status > 0`,
				values:  []interface{}{},
			},
			{
				in:      []reflect.Value{rv(1)},
				stmtSQL: `select * from users where status > $1`,
				realSQL: `select * from users where status > 1`,
				values:  []interface{}{1},
			},
			{
				in:      []reflect.Value{rv(2)},
				stmtSQL: `select * from users where status = $1`,
				realSQL: `select * from users where status = 2`,
				values:  []interface{}{2},
			},
		},
	},
}

var testParseInsertCases = []testParseMapperCase{
	{
		definition: `
		<insert id="InserterP001" parameter="member">
			<inserter table="'members'">
				<field name="'username'">#{member.Username}</field>
				<field name="'password'">#{member.Password}</field>
				<field name="'status'">#{member.Status}</field>
			</inserter>
		</insert>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{rv(struct {
					Username string
					Password string
					Status   int
				}{
					Username: "InserterP001",
					Password: "123456",
					Status:   1,
				})},
				stmtSQL: `insert into members("username","password","status") values($1,$2,$3)`,
				realSQL: `insert into members("username","password","status") values('InserterP001','123456',1)`,
				values:  []interface{}{"InserterP001", "123456", 1},
			},
		},
	},
	{
		definition: `
		<insert id="InserterP002" parameter="member">
			<inserter table="'members'" entity="member">
				<field name="*" />
			</inserter>
		</insert>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{rv(struct {
					Username string
					Password string
					Status   int
				}{
					Username: "InserterP002",
					Password: "123456",
					Status:   1,
				})},
				stmtSQL: `insert into members("username","password","status") values($1,$2,$3)`,
				realSQL: `insert into members("username","password","status") values('InserterP002','123456',1)`,
				values:  []interface{}{"InserterP002", "123456", 1},
			},
		},
	},
	{
		definition: `
		<insert id="InserterP003" parameter="member">
			<inserter table="'members'" entity="member">
				<field name="*" />
				<exclude name="'email'" />
				<exclude name="'mobile'" />
			</inserter>
		</insert>`,
		method: rv(func(row string) (err error) { return }),
		sqls: []*testParseMapperCaseSql{
			{
				in: []reflect.Value{rv(struct {
					Username string
					Email    string
					Mobile   string
					Password string
					Status   int
				}{
					Username: "InserterP003",
					Password: "123456",
					Status:   1,
				})},
				stmtSQL: `insert into members("username","password","status") values($1,$2,$3)`,
				realSQL: `insert into members("username","password","status") values('InserterP003','123456',1)`,
				values:  []interface{}{"InserterP003", "123456", 1},
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

func (p testParseMapperCaseSql) check(t *testing.T, c *testParseMapperCase, s *Stmt) bool {
	require.Equal(t, p.stmtSQL, s.sql, c)
	require.Equal(t, p.realSQL, s.RealSQL(), c)
	require.Equal(t, len(p.values), len(s.vars), c)
	for i, v := range s.vars {
		require.Equal(t, p.values[i], v, c)
	}
	return false
}

// test parse mappers
func TestParseSelectCases(t *testing.T) {
	var (
		fs  []*fragment
		f   *fragment
		c   *caller
		s   *Stmt
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
			s, err = c.fragment.buildStmt(sql.in)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			sql.check(t, &item, s)
			t.Log(fmt.Sprintf("[%s]", f.id), s.RealSQL())
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
			<block type="SOURCE">
				from users where age >= #{age}
			</block>
			<block type="PAGING">
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
			<block type="SOURCE">
				from users
				<where>
				age &lt;= #{age}
				</where>
			</block>
			<block type="PAGING">
				order by age desc
				limit #{limit} offset #{ paging(page,limit)}
			
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
		fs []*fragment
		f  *fragment
		c  *caller
		//s  *stmt
		ss []*Stmt
		//parser *exprParser
		err error
	)
	for _, item := range testParseQueryCases {
		fs, err = parseMapper("fragment_test.go", wrapMapperSchema(item.definition))
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
			ss, err = c.fragment.buildQuery(sql.in)
			//parser, err = c.fragment.prepareParser(s.in)
			require.NoError(t, err)
			//err = c.fragment.buildStmt(parser, s, c.fragment.node)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			t.Log(f.id, ss[0].RealSQL())
			//t.Log(f.id, ss[0].sql)
			//d, _ := json.Marshal(ss[0].vars)
			//t.Log(f.id, string(d))
			//
			t.Log(f.id, ss[1].RealSQL())
			//d, _ = json.Marshal(ss[1].vars)
			//t.Log(f.id, string(d))
		}
	}
}
