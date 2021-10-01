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

var testParseMappersCases = []testParseMapperCase{
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
func TestParseMappers(t *testing.T) {
	var (
		fs     []*method
		f      *method
		c      *caller
		s      *segment
		parser *exprParser
		err    error
	)
	for _, item := range testParseMappersCases {
		fs, err = parseMapper("", wrapMapperSchema(item.definition))
		if item.error {
			require.Equal(t, true, err != nil)
			continue
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, 1, len(fs))
		f = fs[0]
		for _, sql := range item.sqls {
			c = f.newCaller(nil)
			s = c.prepareSegment(sql.in)
			parser, err = c.fragment.prepareParser(s.in)
			require.NoError(t, err)
			err = c.fragment.buildSegment(parser, s, c.fragment.node)
			if sql.error {
				require.Equal(t, true, err != nil)
				continue
			} else {
				require.NoError(t, err)
			}
			//f.par
			t.Log(f.id, s.sql)
			d, _ := json.Marshal(s.vars)
			t.Log(f.id, string(d))
		}
	}
}

// test parse inserter
func testParseInserter(t *testing.T, f *method, c testParseMapperCase) {
	//f.parseInserter(f)
	//s := new(segment)
	//f.buildSegment(s,c.)
}
