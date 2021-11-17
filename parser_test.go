package gobatis

import (
	"fmt"
	"github.com/gozelle/decimal"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testExpression struct {
	In        []interface{} `json:"in"`
	Parameter string        `json:"parameter"`
	Expr      string        `json:"expr"`
	Result    interface{}   `json:"result"`
	Err       int           `json:"err"`
}

func TestCorrectParseExprExpression(t *testing.T) {
	
	testCorrectParseExprExpression(t, []testExpression{
		
		// test variable visit
		{In: []interface{}{1, 2}, Parameter: "a,b", Expr: "a + b", Result: 3},
		{In: []interface{}{1, 2}, Parameter: "a:int,b", Expr: "a + b", Result: 3},
		{In: []interface{}{1, 2}, Parameter: "a,b:int", Expr: "a + b", Result: 3},
		{In: []interface{}{1, 2}, Parameter: "a:int,b:int", Expr: "a + b", Result: 3},
		
		// TODO Check must export
		{In: []interface{}{struct{ b int }{b: 1}}, Parameter: "a", Expr: "a.b", Result: 1, Err: expr_syntax_err},
		
		{In: []interface{}{struct{ B int }{B: 1}}, Parameter: "a", Expr: "a.B", Result: 1},
		{In: []interface{}{[]int{1}}, Parameter: "a", Expr: "a[0]", Result: 1},
		{In: []interface{}{map[string]int{"b": 1}}, Parameter: "a", Expr: `a["b"]`, Result: 1},
		{In: []interface{}{func() int { return 1 }}, Parameter: "a", Expr: `a()`, Result: 1},
		
		// test number add calc
		{In: []interface{}{1, 1}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{int8(1), 1}, Parameter: "a,b", Expr: "a + b", Result: int8(2)},
		{In: []interface{}{int16(1), 1}, Parameter: "a,b", Expr: "a + b", Result: int16(2)},
		{In: []interface{}{int32(1), 1}, Parameter: "a,b", Expr: "a + b", Result: int32(2)},
		{In: []interface{}{int64(1), 1}, Parameter: "a,b", Expr: "a + b", Result: int64(2)},
		{In: []interface{}{uint(1), 1}, Parameter: "a,b", Expr: "a + b", Result: uint(2)},
		{In: []interface{}{uint8(1), 1}, Parameter: "a,b", Expr: "a + b", Result: uint8(2)},
		{In: []interface{}{uint16(1), 1}, Parameter: "a,b", Expr: "a + b", Result: uint16(2)},
		{In: []interface{}{uint32(1), 1}, Parameter: "a,b", Expr: "a + b", Result: uint32(2)},
		{In: []interface{}{uint64(1), 1}, Parameter: "a,b", Expr: "a + b", Result: uint64(2)},
		
		{In: []interface{}{1, int8(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, int16(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, int32(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, int64(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, uint(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, uint8(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, uint16(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, uint32(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, uint64(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		
		// TODO FIX
		{In: []interface{}{float32(1), 1}, Parameter: "a,b", Expr: "a + b", Result: float32(2), Err: expr_syntax_err},
		{In: []interface{}{float64(1), 1}, Parameter: "a,b", Expr: "a + b", Result: float64(2), Err: expr_syntax_err},
		
		{In: []interface{}{1, float32(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		{In: []interface{}{1, float64(1)}, Parameter: "a,b", Expr: "a + b", Result: 2},
		
		// test logic calc start
		{In: []interface{}{true, true}, Parameter: "a,b", Expr: "a && b", Result: true},
		{In: []interface{}{true, false}, Parameter: "a,b", Expr: "a && b", Result: false},
		{In: []interface{}{false, true}, Parameter: "a,b", Expr: "a && b", Result: false},
		{In: []interface{}{true, true}, Parameter: "a,b", Expr: "a || b", Result: true},
		{In: []interface{}{true, false}, Parameter: "a,b", Expr: "a || b", Result: true},
		{In: []interface{}{false, true}, Parameter: "a,b", Expr: "a || b", Result: true},
		
		{In: []interface{}{true, true}, Parameter: "a,b", Expr: "a and b", Result: true},
		{In: []interface{}{true, false}, Parameter: "a,b", Expr: "a and b", Result: false},
		{In: []interface{}{false, true}, Parameter: "a,b", Expr: "a and b", Result: false},
		{In: []interface{}{true, true}, Parameter: "a,b", Expr: "a or b", Result: true},
		{In: []interface{}{true, false}, Parameter: "a,b", Expr: "a or b", Result: true},
		{In: []interface{}{false, true}, Parameter: "a,b", Expr: "a or b", Result: true},
		
		{In: []interface{}{true, true, true}, Parameter: "a,b,c", Expr: "a && b and c", Result: true},
		{In: []interface{}{true, true, false}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{true, false, false}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{false, true, true}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{false, false, true}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		{In: []interface{}{false, false, false}, Parameter: "a,b,c", Expr: "a && b and c", Result: false},
		
		{In: []interface{}{true, true, true}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{true, true, false}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{true, false, false}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{false, true, true}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{false, false, true}, Parameter: "a,b,c", Expr: "a or b || c", Result: true},
		{In: []interface{}{false, false, false}, Parameter: "a,b,c", Expr: "a or b || c", Result: false},
		
		{In: []interface{}{true, true, true}, Parameter: "a,b,c", Expr: "a and b or c", Result: true},
		{In: []interface{}{true, true, false}, Parameter: "a,b,c", Expr: "a and b or c", Result: true},
		{In: []interface{}{true, false, false}, Parameter: "a,b,c", Expr: "a and b or c", Result: false},
		{In: []interface{}{false, true, true}, Parameter: "a,b,c", Expr: "a and b or c", Result: true},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a and b or c", Result: false},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a and b or c", Result: false},
		{In: []interface{}{false, false, true}, Parameter: "a,b,c", Expr: "a and b or c", Result: true},
		{In: []interface{}{false, false, false}, Parameter: "a,b,c", Expr: "a and b or c", Result: false},
		
		{In: []interface{}{true, true, true}, Parameter: "a,b,c", Expr: "a or b and c", Result: true},
		{In: []interface{}{true, true, false}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		{In: []interface{}{true, false, false}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		{In: []interface{}{false, true, true}, Parameter: "a,b,c", Expr: "a or b and c", Result: true},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		{In: []interface{}{false, true, false}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		{In: []interface{}{false, false, true}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		{In: []interface{}{false, false, false}, Parameter: "a,b,c", Expr: "a or b and c", Result: false},
		// test logic calc end
		
		// test unary calc start
		{In: []interface{}{true}, Parameter: "a", Expr: "!a", Result: false},
		{In: []interface{}{true}, Parameter: "a", Expr: "!!a", Result: true},
		{In: []interface{}{true}, Parameter: "a", Expr: "!!!a", Result: false},
		{In: []interface{}{false}, Parameter: "a", Expr: "!a", Result: true},
		{In: []interface{}{false}, Parameter: "a", Expr: "!!a", Result: false},
		{In: []interface{}{false}, Parameter: "a", Expr: "!!!a", Result: true},
		{In: []interface{}{true, false, false}, Parameter: "a,b,c", Expr: "!!a and !!!b and !c", Result: true},
		
		{In: []interface{}{1}, Parameter: "a", Expr: "+1", Result: 1},
		{In: []interface{}{1}, Parameter: "a", Expr: "++1", Result: 1, Err: expr_syntax_err},
		{In: []interface{}{1}, Parameter: "a", Expr: "+++1", Result: 1, Err: expr_syntax_err},
		{In: []interface{}{1}, Parameter: "a", Expr: "-1", Result: -1},
		{In: []interface{}{1}, Parameter: "a", Expr: "--1", Result: -1, Err: expr_syntax_err},
		{In: []interface{}{1}, Parameter: "a", Expr: "---1", Result: -1, Err: expr_syntax_err},
		// test unary calc end
		
		// test ternary start
		{In: []interface{}{true, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 1},
		{In: []interface{}{false, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 2},
		{In: []interface{}{0, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 2},
		{In: []interface{}{1, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 1},
		{In: []interface{}{-1, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 1},
		{In: []interface{}{uint(0), 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 2},
		{In: []interface{}{uint(1), 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 1},
		{In: []interface{}{0.0, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{float32(0), 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{float64(0), 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{nil, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Result: 2},
		{In: []interface{}{struct{}{}, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{map[string]int{}, 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{"", 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		{In: []interface{}{" ", 1, 2}, Parameter: "a,b,c", Expr: "a ? b : c", Err: expr_syntax_err},
		
		{In: []interface{}{"", 1, 2}, Parameter: "a,b,c", Expr: "a == '' ? b : c", Result: 1},
		{In: []interface{}{"", 1, 2}, Parameter: "a,b,c", Expr: "a != '' ? b : c", Result: 2},
		{In: []interface{}{"ok", 1, 2}, Parameter: "a,b,c", Expr: "a =='ok' ? b+10 : c", Result: 11},
		{In: []interface{}{"ok", 1, 2}, Parameter: "a,b,c", Expr: "a != nil ? b+10 : c", Result: 11},
		{In: []interface{}{map[string]int{}, 1, 2}, Parameter: "a,b,c", Expr: "a != nil ? b+10 : c", Result: 11},
		{In: []interface{}{nil, 1, 2}, Parameter: "a,b,c", Expr: "a == nil ? b+10 : c", Result: 11},
		// test ternary end
		
		// test priority strat
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + b + c", Result: 6},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + b - c", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + b * c", Result: 7},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + b / c", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + b % c", Result: 3},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a + b) + c", Result: 6},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a + b) - c", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a + b) * c", Result: 9},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a + b) / c", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a + b) % c", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + (b + c)", Result: 6},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + (b - c)", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + (b * c)", Result: 7},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + (b / c)", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a + (b % c)", Result: 3},
		
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - b + c", Result: 2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - b - c", Result: -4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - b * c", Result: -5},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - b / c", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - b % c", Result: -1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a - b) + c", Result: 2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a - b) - c", Result: -4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a - b) * c", Result: -3},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a - b) / c", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(a - b) % c", Result: -1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - (b + c)", Result: -4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - (b - c)", Result: 2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - (b * c)", Result: -5},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - (b / c)", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "a - (b % c)", Result: -1},
		
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + b + c", Result: 4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + b - c", Result: -2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + b * c", Result: 5},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + b / c", Result: -1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + b % c", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(-a + b) + c", Result: 4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(-a + b) - c", Result: -2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(-a + b) * c", Result: 3},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(-a + b) / c", Result: 0},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "(-a + b) % c", Result: 1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + (b + c)", Result: 4},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + (b - c)", Result: -2},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + (b * c)", Result: 5},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + (b / c)", Result: -1},
		{In: []interface{}{1, 2, 3}, Parameter: "a,b,c", Expr: "-a + (b % c)", Result: 1},
		// test priority end
		
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a:int, b", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a:int, b:int", Expr: "a + b", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a, b:int", Expr: "a + b", Result: 6},
		{In: []interface{}{int8(2), int8(4)}, Parameter: "a,b", Expr: "a - b", Result: -int8(2)},
		{In: []interface{}{int16(2), int16(4)}, Parameter: "a,b", Expr: "a * b", Result: int16(8)},
		{In: []interface{}{int32(2), int32(4)}, Parameter: "a,b", Expr: "a / b ", Result: int32(0)},
		{In: []interface{}{int64(2), int64(4)}, Parameter: "a,b", Expr: "b / a", Result: int64(2)},
		{In: []interface{}{decimal.NewFromFloat(3.12), "2.13"}, Parameter: "a,b", Expr: "a + b", Result: "5.25"},
		{In: []interface{}{decimal.NewFromFloat(3.12), 2.13}, Parameter: "a,b", Expr: "a + b", Result: "5.25"},
		
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: " a + a * a", Result: 6},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "  a + a * b", Result: 10},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: " a + b * a ", Result: 10},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "a + b * b", Result: 18},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "(a + b) * b", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( a + b  ) * b)", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( (  a + b ) ) * b)", Result: 24},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( a + b) / b", Result: 1},
		
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + b * b", Result: 20},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + b * a ", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + a * b", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "b + a * a", Result: 8},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( b + a ) * a", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( b+ a ) * a)", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( ( ( b + a ) * a ) )", Result: 12},
		{In: []interface{}{2, 4}, Parameter: "a,b", Expr: "( b + a ) / a", Result: 3},
		
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Age + b.Age", Result: 38},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() + b.Weight()", Result: 100},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Weight() > b.Weight()", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "a.Parent == nil", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "a, b", Expr: "nil == a.Parent && nil != b.Parent", Result: true},
		//{In: []interface{}{u1, u2}, Parameter: "parent1, child2", Expr: "child2.Parent.Age + child2.Age", Result: 38},
	})
}

func TestAnyExprParam(t *testing.T) {
	params, err := testAnyExprParam(" * ")
	require.Equal(t, 0, len(params))
	require.NoError(t, err)
	
	_, err = testAnyExprParam("*,")
	require.Error(t, err)
	
	_, err = testAnyExprParam("*,a")
	require.Error(t, err)
	
	_, err = testAnyExprParam(",*")
	require.Error(t, err)
}

func testAnyExprParam(tokens string) (params []*param, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	parser := newParamParser("test.xml")
	parser.walkMethods(initExprParser(tokens))
	params = parser.params
	return
}

func testCorrectParseExprExpression(t *testing.T, tests []testExpression) {
	for _, item := range tests {
		vars := make([]reflect.Value, 0)
		for _, v := range item.In {
			vars = append(vars, rv(v))
		}
		_expr := newExprParser(vars...)
		_expr.file = "expr.xml"
		
		params, err := testParseParams(_expr.file, item.Parameter)
		require.NoError(t, err, item)
		
		for ii, vv := range params {
			err = _expr.paramsStack.list.Front().Next().Value.(*exprParams).bind(vv, ii)
			require.NoError(t, err, item)
		}
		result, _, err := _expr.parseExpression(nil, item.Expr)
		if item.Err > 0 {
			require.Error(t, err, item)
			//writeError(t, fmt.Sprintf("test parse expression: %d", i), item, err)
		} else {
			require.NoError(t, err, item)
			dr, ok := result.(decimal.Decimal)
			if ok {
				require.Equal(t, item.Result, dr.String(), item)
			} else {
				require.Equal(t, item.Result, result, item)
			}
		}
	}
}

func testParseParams(file, tokens string) (params []*param, err error) {
	defer func() {
		e := recover()
		err = catch(file, e)
	}()
	params = parseParams(file, tokens)
	return
}
