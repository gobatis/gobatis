package expr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {

	type A struct {
		V int64
	}

	v, err := Parse(`d`, map[string]any{
		"a": A{
			V: 10,
		},
		"b": true,
		"c": false,
		"d": []string{"d1", "d2"},
	})
	require.NoError(t, err)
	t.Log(v.Interface())
}

func TestA(t *testing.T) {
	//fmt.Println()
}

func TestExpressions(t *testing.T) {

	type TestCase struct {
		Expr   string
		Error  bool
		Vars   map[string]any
		Expect interface{}
	}

	var ptr *int64

	type Obj struct {
		A1 string
		b1 string
	}

	tests := []TestCase{

		// nil & pointer
		{Expr: "nil", Error: false, Vars: map[string]any{}, Expect: nil},
		{Expr: "nil==nil", Vars: map[string]any{"a": ptr}, Expect: true},
		{Expr: "nil!=nil", Vars: map[string]any{"a": ptr}, Expect: false},
		{Expr: "a==nil", Vars: map[string]any{"a": nil}, Expect: true},
		{Expr: "a!=nil", Vars: map[string]any{"a": nil}, Expect: false},
		{Expr: "nil==a", Vars: map[string]any{"a": nil}, Expect: true},
		{Expr: "nil!=a", Vars: map[string]any{"a": nil}, Expect: false},
		{Expr: "a==nil", Vars: map[string]any{"a": ptr}, Expect: true},
		{Expr: "a!=nil", Vars: map[string]any{"a": ptr}, Expect: false},
		{Expr: "nil==a", Vars: map[string]any{"a": ptr}, Expect: true},
		{Expr: "nil!=a", Vars: map[string]any{"a": ptr}, Expect: false},

		{Expr: "b==nil", Error: true, Vars: map[string]any{"a": ptr}},
		{Expr: "a==nil", Error: true, Vars: map[string]any{"a": 1}},

		// assignment
		{Expr: "a", Vars: map[string]any{"a": 1}, Expect: 1},
		{Expr: "a", Vars: map[string]any{"a": "abc"}, Expect: "abc"},
		{Expr: "a > b", Vars: map[string]any{"a": 2, "b": 1}, Expect: true},
		{Expr: "a != 0", Vars: map[string]any{"a": 2}, Expect: true},

		// logical
		{Expr: "a==1 && b==2", Vars: map[string]any{"a": 1, "b": 2}, Expect: true},
		{Expr: "a==1 || b!=2", Vars: map[string]any{"a": 1, "b": 2}, Expect: true},
		{Expr: "(a==1 || a==2) && b==2", Vars: map[string]any{"a": 1, "b": 2}, Expect: true},
		{Expr: "a==1 && (b==1 || b==2)", Vars: map[string]any{"a": 1, "b": 2}, Expect: true},

		// member
		{Expr: "o.A1", Vars: map[string]any{"o": Obj{A1: "a1"}}, Expect: "a1"},
		{Expr: "o.b1", Error: true, Vars: map[string]any{"o": Obj{b1: "b1"}}},

		// index
		{Expr: "s[1]", Vars: map[string]any{"s": []string{"a", "b"}}, Expect: "b"},

		// slice
		{Expr: "s[1:][0]", Vars: map[string]any{"s": []string{"a", "b"}}, Expect: "b"},

		// call
		{Expr: "len(s)", Vars: map[string]any{"s": []string{"a", "b"}}, Expect: 2},

		// ternary
		{Expr: "a > 0 ? 1 : -1", Vars: map[string]any{"a": 1}, Expect: int64(1)},
		{Expr: "a < 0 ? 1 : -1", Vars: map[string]any{"a": 1}, Expect: int64(-1)},
	}

	for _, v := range tests {
		r, err := Parse(v.Expr, v.Vars)
		if v.Error {
			require.Error(t, err)
		} else {
			require.NoError(t, err, v.Expr)
			if v.Expect == nil {
				require.True(t, !r.IsValid(), v.Expr)
			} else {
				require.Equal(t, v.Expect, r.Interface(), v.Expr)
			}
		}
	}
}
