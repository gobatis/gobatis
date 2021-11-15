package stmt

import (
	"github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type TestCase struct {
	Method interface{}
	In     []interface{}
	SQL    string
	Err    string
}

func rv(item interface{}) reflect.Value {
	return reflect.ValueOf(item)
}

func prepareTestCases(mapper *Mapper) []TestCase {
	return []TestCase{
		{
			Method: mapper.InsertS001Stmt,
			SQL:    "insert into a(b,c) values ('b','c');",
		},
		{
			Method: mapper.InsertS002Stmt,
			SQL:    "insert into a(b,c) values ('b','c');",
		},
		{
			Method: mapper.InsertS003Stmt,
			SQL:    "insert into a(b,c) values( 'b', 'c1' )",
		},
		{
			Method: mapper.InsertS004Stmt,
			SQL:    "insert into a(b,c)\n        values(\n        'b', 'c1' )",
		},
		{
			Method: mapper.InsertS005Stmt,
			In:     []interface{}{[]Item{{B: "b1", C: "c1"}, {B: "b2", C: "c2"}}},
			SQL:    "insert into a(b,c) values ('b1','c1'),('b2','c2')",
		},
		{
			Method: mapper.UpdateS001Stmt,
			SQL:    "update a set b='b', c='c' where id = 1;",
		},
		{
			Method: mapper.UpdateS002Stmt,
			SQL:    "update a set b='b', c='c' where id = 1;",
		},
		{
			Method: mapper.UpdateS003Stmt,
			SQL:    "update a set b='b', c='c' where id = 1;",
		},
		{
			Method: mapper.UpdateS004Stmt,
			SQL:    "update a set b='b', c='c' where id = 1;",
		},
		{
			Method: mapper.UpdateS005Stmt,
			In:     []interface{}{[]Item{{B: "b1", C: "c1"}, {B: "b2", C: "c2"}}},
			SQL:    "insert into a(b,c) values ('b1','c1'),('b2','c2')",
		},
	}
}

func TestStmtMapper(t *testing.T) {
	engine := gobatis.NewEngine(gobatis.NewDB("", ""))
	engine.InitLogger()
	err := engine.RegisterBundle(gobatis.NewBundle("./sql"))
	require.NoError(t, err)
	mapper := NewMapper()
	err = engine.BindMapper(mapper)
	require.NoError(t, err)
	cases := prepareTestCases(mapper)
	var (
		in   []reflect.Value
		out  []reflect.Value
		stmt *gobatis.Stmt
		ok   bool
	)
	for _, v := range cases {
		in = []reflect.Value{}
		for _, vv := range v.In {
			in = append(in, rv(vv))
		}
		out = rv(v.Method).Call(in)
		require.True(t, len(out) >= 1, v)
		if v.Err != "" {
			require.Equal(t, v.Err, out[len(out)-1].Interface().(error).Error(), v)
		} else {
			require.Equal(t, nil, out[len(out)-1].Interface(), v)
			require.True(t, out[len(out)-1].IsValid(), v)
			if len(out) > 1 {
				for i := 0; i < len(out)-1; i++ {
					stmt, ok = out[i].Interface().(*gobatis.Stmt)
					require.True(t, ok, v)
					if i == 0 {
						t.Log(stmt.ID(), stmt.RealSQL())
						//require.Equal(t, v.SQL, stmt.RealSQL(), v)
					} else {
						panic("implement it")
					}
				}
			}
		}
	}
}
