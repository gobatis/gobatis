package expr

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	//"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/parser/commons"
)

type Expr struct {
	vars map[string]any
}

func valueOf(source interface{}) reflect.Value {
	v := reflect.ValueOf(source)
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			return v
		}
	}
}

type FetchVar func(name string) (val any, ok bool)

func Parse(source string, fetchVar FetchVar) (reflect.Value, error) {

	errs := &commons.ErrorListener{}

	lexer := NewExprLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewExprParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.Expressions()
	if errs.Error() != nil {
		return reflect.Value{}, errs.Error()
	}
	v := &Visitor{
		ErrorListener: errs,
		//vars:          map[string]reflect.Value{},
		fetchVar: fetchVar,
	}
	//for kk, vv := range vars {
	//	v.vars[kk] = reflect.ValueOf(vv)
	//}
	return v.visitExpressions(tree.(*ExpressionsContext))
}

type Visitor struct {
	*commons.ErrorListener
	fetchVar FetchVar
}

func (v Visitor) visitExpressions(ctx *ExpressionsContext) (rv reflect.Value, err error) {
	if ctx.Expression() == nil {
		return reflect.Value{}, fmt.Errorf("expression is nil")
	}
	defer func() {
		err = v.Error()
	}()
	rv = v.visitExpression(ctx.Expression().(*ExpressionContext))
	if rv.Kind() != reflect.Invalid {
		err = fmt.Errorf("invalid reflect value")
	}
	return
}

func (v Visitor) visitExpression(ctx *ExpressionContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if ctx.Primary() != nil {
		//fmt.Println("primary:", ctx.GetText())
		return v.visitPrimary(ctx.Primary().(*PrimaryContext))
	} else if ctx.GetUnary() != nil {
		rv := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		return v.visitUnary(ctx.GetUnary().GetText(), rv)
	} else if ctx.GetRel() != nil {
		//fmt.Println("rel:", ctx.GetText())
		a := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		b := v.visitExpression(ctx.Expression(1).(*ExpressionContext))
		return v.visitRel(ctx.GetRel().GetText(), a, b)
	} else if ctx.Logical() != nil {
		a := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		b := v.visitExpression(ctx.Expression(1).(*ExpressionContext))
		return v.visitLogical(ctx.Logical().GetText(), a, b)
	} else if ctx.GetTertiary() != nil {
		a := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		b := v.visitExpression(ctx.Expression(1).(*ExpressionContext))
		c := v.visitExpression(ctx.Expression(2).(*ExpressionContext))
		return v.visitTertiary(a, b, c)
	} else {
		v.AddError(fmt.Errorf("unsupported expression: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitPrimary(ctx *PrimaryContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if ctx.Primary() != nil {
		rv := v.visitPrimary(ctx.Primary().(*PrimaryContext))
		if ctx.Member() != nil {
			return v.visitMember(rv, ctx.Member().(*MemberContext))
		} else if ctx.Index() != nil {
			return v.visitIndex(rv, ctx.Index().(*IndexContext))
		} else if ctx.Slice() != nil {
			return v.visitSlice(rv, ctx.Slice().(*SliceContext))
		} else if ctx.Call() != nil {
			return v.visitCall(rv, ctx.Call().(*CallContext))
		} else {
			v.AddError(fmt.Errorf("unsupported primary expression: %s", ctx.GetText()))
			return reflect.Value{}
		}
	} else if ctx.Operand() != nil {
		return v.visitOperand(ctx.Operand().(*OperandContext))
	} else {
		v.AddError(fmt.Errorf("unsupported primary expression: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitOperand(ctx *OperandContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if ctx.Literal() != nil {
		return v.visitLiteral(ctx.Literal().(*LiteralContext))
	} else if ctx.Var_() != nil {
		return v.visitVar(ctx.Var_().(*VarContext))
	} else if ctx.Expression() != nil {
		return v.visitExpression(ctx.Expression().(*ExpressionContext))
	} else {
		v.AddError(fmt.Errorf("unsupported operand: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitLiteral(ctx *LiteralContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if ctx.Nil_() != nil {
		return reflect.ValueOf(nil)
	} else if ctx.Integer() != nil {
		return v.visitInteger(ctx.Integer().(*IntegerContext))
	} else if ctx.String_() != nil {
		return v.visitString(ctx.String_().(*StringContext))
	} else if ctx.Float() != nil {
		return v.visitFloat(ctx.Float().(*FloatContext))
	} else {
		v.AddError(fmt.Errorf("unsupport literal: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitVar(ctx *VarContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	// TODO check nil 关键字
	vv, ok := v.fetchVar(ctx.GetText())
	if !ok {
		if _builtin.get(ctx.GetText()) != nil {
			return reflect.ValueOf(_builtin.get(ctx.GetText()))
		}
		v.AddError(fmt.Errorf("variable: %s is not defined", ctx.GetText()))
		return reflect.Value{}
	}
	return reflect.ValueOf(vv)
}

func (v Visitor) visitMember(rv reflect.Value, ctx *MemberContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if rv.Kind() != reflect.Struct {
		v.AddError(fmt.Errorf("variable: %s expect struct, got: %s", ctx.GetText(), rv.Kind()))
		return reflect.Value{}
	}
	fv := rv.FieldByName(ctx.IDENTIFIER().GetText())
	if fv.Kind() == reflect.Invalid {
		v.AddError(fmt.Errorf("visit member: %s not exist", ctx.IDENTIFIER().GetText()))
		return reflect.Value{}
	}
	if startsWithLower(ctx.IDENTIFIER().GetText()) {
		v.AddError(fmt.Errorf("visit memeber: %s is not exportable", ctx.IDENTIFIER().GetText()))
		return reflect.Value{}
	}
	return fv
}

func (v Visitor) visitIndex(rv reflect.Value, ctx *IndexContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	iv := v.visitExpression(ctx.Expression().(*ExpressionContext))

	if !iv.CanInt() && !iv.CanUint() {
		v.AddError(fmt.Errorf("invalid index type: %s", iv.Kind()))
		return reflect.Value{}
	}

	var i int
	if iv.CanInt() {
		i = int(iv.Int())
	} else {
		i = int(iv.Uint())
	}

	rv = rv.Index(i)
	if rv.Kind() == reflect.Invalid {
		v.AddError(fmt.Errorf("invalid index: %d", i))
		return reflect.Value{}
	}

	return rv
}

func (v Visitor) visitSlice(rv reflect.Value, ctx *SliceContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	if rv.Kind() != reflect.Slice {
		v.AddError(fmt.Errorf("unsupported call slice on type: %s", rv.Kind()))
		return reflect.Value{}
	}

	var sea, seb, sec reflect.Value
	if ctx.GetSea() != nil {
		sea = v.visitExpression(ctx.GetSea().(*ExpressionContext))
		if v.Error() != nil {
			return reflect.Value{}
		}
	}
	if !sea.IsValid() {
		sea = reflect.ValueOf(0)
	}

	if ctx.GetSeb() != nil {
		seb = v.visitExpression(ctx.GetSeb().(*ExpressionContext))
		if v.Error() != nil {
			return reflect.Value{}
		}
	}

	if ctx.GetSec() != nil {
		sec = v.visitExpression(ctx.GetSec().(*ExpressionContext))
		if v.Error() != nil {
			return reflect.Value{}
		}
	} else if !seb.IsValid() {
		seb = reflect.ValueOf(rv.Len())
	}

	if !sea.CanInt() && !sea.CanUint() {
		v.AddError(fmt.Errorf("invliad slice index type: %s", sea.Kind()))
		return reflect.Value{}
	}

	if !seb.CanInt() && seb.CanUint() {
		v.AddError(fmt.Errorf("invliad slice index type: %s", seb.Kind()))
		return reflect.Value{}
	}

	if sec.IsValid() && !sec.CanInt() && seb.CanUint() {
		v.AddError(fmt.Errorf("invliad slice index type: %s", sec.Kind()))
		return reflect.Value{}
	}

	var i, j, k int
	if sea.CanInt() {
		i = int(sea.Int())
	} else {
		i = int(sea.Uint())
	}

	if seb.CanInt() {
		j = int(seb.Int())
	} else {
		j = int(seb.Uint())
	}

	if sec.IsValid() {
		if sec.CanInt() {
			k = int(sec.Int())
		} else {
			k = int(sec.Uint())
		}
	}

	if sec.IsValid() {
		return rv.Slice3(i, j, k)
	}

	return rv.Slice(i, j)
}

func (v Visitor) visitCall(rv reflect.Value, ctx *CallContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	var params []reflect.Value
	for i := 0; i < len(ctx.ExpressionList().AllExpression()); i++ {
		vv := v.visitExpression(ctx.ExpressionList().Expression(i).(*ExpressionContext))
		if v.Error() != nil {
			return reflect.Value{}
		}
		params = append(params, vv)
	}
	// TODO check result error
	r := rv.Call(params)
	if len(r) > 0 {
		return r[0]
	}
	return reflect.Value{}
}

func (v Visitor) visitUnary(op string, rv reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	switch op {
	case "+":
		if rv.CanInt() || rv.CanUint() || rv.CanFloat() {
			return rv
		}
	case "-":
		if rv.CanInt() || rv.CanFloat() {
			if rv.CanInt() {
				return reflect.ValueOf(-rv.Int())
			} else {
				return reflect.ValueOf(-rv.Float())
			}
		}
	}

	v.AddError(fmt.Errorf("unsupported unary operation: %s for type: %s", op, rv.Kind()))
	return reflect.Value{}
}

func (v Visitor) visitRel(op string, a, b reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	if !a.IsValid() || !b.IsValid() {
		return v.compareNil(op, a, b)
	}

	if !a.Comparable() || !b.Comparable() {
		v.AddError(fmt.Errorf("mismatched compare types: %s and %s", a.Kind(), b.Kind()))
		return reflect.Value{}
	}

	if v.isNumber(a) || v.isNumber(b) {
		if v.isNumber(a) {
			if !v.isNumber(b) {
				v.AddError(fmt.Errorf("expact number, got: %s", b.Kind()))
				return reflect.Value{}
			}
		} else {
			if !v.isNumber(a) {
				v.AddError(fmt.Errorf("expact number, got: %s", b.Kind()))
				return reflect.Value{}
			}
		}
		return v.compareNumber(op, a, b)
	}

	if a.Type().String() != b.Type().String() || a.Type().PkgPath() != b.Type().PkgPath() {
		v.AddError(fmt.Errorf("mismatched number compare types %s and %s", a.Type(), b.Type()))
		return reflect.Value{}
	}

	switch a.Kind() {
	case reflect.Bool:
		return v.compareBool(op, a, b)
	}

	return reflect.Value{}
}

func (v Visitor) compareNil(op string, a, b reflect.Value) reflect.Value {
	switch op {
	case "==":
		return reflect.ValueOf(a.Equal(b))
		//return reflect.ValueOf((!a.IsValid() || a.IsNil()) && (!b.IsValid() || b.IsNil()))
	case "!=":
		return reflect.ValueOf(!a.Equal(b))
		//return reflect.ValueOf(!(!a.IsValid() || a.IsNil()) && (!b.IsValid() || b.IsNil()))
	default:
		v.AddError(fmt.Errorf("unsupport nil compare operation: %s", op))
		return reflect.Value{}
	}
}

func (v Visitor) compareNumber(op string, a, b reflect.Value) reflect.Value {
	switch op {
	case "==":
		return reflect.ValueOf(v.toFloat(a) == v.toFloat(b))
	case "!=":
		return reflect.ValueOf(v.toFloat(a) != v.toFloat(b))
	case "<":
		return reflect.ValueOf(v.toFloat(a) < v.toFloat(b))
	case "<=":
		return reflect.ValueOf(v.toFloat(a) <= v.toFloat(b))
	case ">":
		return reflect.ValueOf(v.toFloat(a) > v.toFloat(b))
	case ">=":
		return reflect.ValueOf(v.toFloat(a) >= v.toFloat(b))
	default:
		v.AddError(fmt.Errorf("unsupport number compare operation: %s", op))
		return reflect.Value{}
	}
}

func (v Visitor) compareBool(op string, a, b reflect.Value) reflect.Value {
	switch op {
	case "==":
		return reflect.ValueOf(a.Bool() == b.Bool())
	case "!=":
		return reflect.ValueOf(a.Bool() != b.Bool())
	default:
		v.AddError(fmt.Errorf("the operator %s is not defined on bool", op))
	}
	return reflect.Value{}
}

func (v Visitor) visitLogical(op string, a, b reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if a.Kind() != reflect.Bool || b.Kind() != reflect.Bool {
		v.AddError(fmt.Errorf("expect bool used as a condition"))
		return reflect.Value{}
	}
	switch op {
	case "&&":
		return reflect.ValueOf(a.Bool() && b.Bool())
	case "||":
		return reflect.ValueOf(a.Bool() || b.Bool())
	default:
		v.AddError(fmt.Errorf("unsupport logical operation: %s", op))
		return reflect.Value{}
	}
}

func (v Visitor) visitTertiary(a, b, c reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	if a.Kind() != reflect.Bool {
		v.AddError(fmt.Errorf("expect bool used as a condition"))
		return reflect.Value{}
	}
	if a.Bool() {
		return b
	}
	return c
}

func (v Visitor) visitInteger(ctx *IntegerContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	r, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		v.AddError(fmt.Errorf("parsee inter: %s error: %w", ctx.GetText(), err))
		return reflect.Value{}
	}

	return reflect.ValueOf(r)
}

func (v Visitor) visitString(ctx *StringContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(ctx.GetText()[1 : len(ctx.GetText())-1])
}

func (v Visitor) visitFloat(ctx *FloatContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	r, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		v.AddError(fmt.Errorf("parse float: %s error: %w", ctx.GetText(), err))
		return reflect.Value{}
	}

	return reflect.ValueOf(r)
}

func (v Visitor) isNumber(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func (v Visitor) toFloat(rv reflect.Value) float64 {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	default:
		v.AddError(fmt.Errorf("covert: %s to number faild", rv.Kind()))
		return 0
	}
}

func startsWithLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsLower(rune(s[0]))
}
