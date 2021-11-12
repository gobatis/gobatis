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
			In:     nil,
			SQL:    "insert into a(b,c) values ('b','c');",
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
			require.Equal(t, v.Err, out[len(out)-1].Interface().(error).Error())
		} else {
			if len(out) > 1 {
				for i := 0; i < len(out)-1; i++ {
					stmt, ok = out[i].Interface().(*gobatis.Stmt)
					require.True(t, ok, v)
					if i == 0 {
						t.Log(stmt.RealSQL())
						require.Equal(t, v.SQL, stmt.RealSQL())
					} else {
						panic("implement it")
					}
				}
			}
		}
	}
}
