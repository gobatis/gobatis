package expr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/shopspring/decimal"
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

func Parse(source string, vars map[string]any) (reflect.Value, error) {

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
		vars:          map[string]reflect.Value{},
	}
	for kk, vv := range vars {
		v.vars[kk] = reflect.ValueOf(vv)
	}
	return v.visitExpressions(tree.(*ExpressionsContext))
}

type Visitor struct {
	*commons.ErrorListener
	vars map[string]reflect.Value
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
		return v.visitLogical(ctx.GetRel().GetText(), a, b)
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
	vv, ok := v.vars[ctx.GetText()]
	if !ok {
		v.AddError(fmt.Errorf("var: %s is not defined", ctx.GetText()))
		return reflect.Value{}
	}
	return vv
}

func (v Visitor) visitMember(rv reflect.Value, ctx *MemberContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	if rv.Kind() != reflect.Struct {
		v.AddError(fmt.Errorf("var: %s expect struct, got: %s", ctx.GetText(), rv.Kind()))
		return reflect.Value{}
	}
	fv := rv.FieldByName(ctx.IDENTIFIER().GetText())
	if fv.Kind() == reflect.Invalid {
		v.AddError(fmt.Errorf("visit member: %s not exist", ctx.IDENTIFIER().GetText()))
		return reflect.Value{}
	}
	return fv
}

func (v Visitor) visitIndex(rv reflect.Value, ctx *IndexContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	iv := v.visitExpression(ctx.Expression().(*ExpressionContext))

	// TODO
	i, err := cast.ToIntE(iv.Interface().(decimal.Decimal).IntPart())
	if err != nil {
		v.AddError(fmt.Errorf("covert index error: %s", err))
		return reflect.Value{}
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
	return reflect.Value{}
}

func (v Visitor) visitCall(rv reflect.Value, ctx *CallContext) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitUnary(op string, rv reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitRel(op string, a, b reflect.Value) reflect.Value {
	if v.Error() != nil {
		return reflect.Value{}
	}

	if !a.Comparable() {
		v.AddError(fmt.Errorf("%s is not comparable", a.Kind()))
		return reflect.Value{}
	}

	if !b.Comparable() {
		v.AddError(fmt.Errorf("%s is not comparable", b.Kind()))
		return reflect.Value{}
	}

	if isNumber(a) {
		if !isNumber(b) {
			v.AddError(fmt.Errorf("expact number, got: %s", b.Kind()))
			return reflect.Value{}
		}
		return v.compareNumber(op, a, b)
	}

	if a.Type().String() != b.Type().String() || a.Type().PkgPath() != b.Type().PkgPath() {
		v.AddError(fmt.Errorf("mismatched types %s and %s", a.Type(), b.Type()))
		return reflect.Value{}
	}

	switch a.Kind() {
	case reflect.Bool:
		return v.compareBool(op, a, b)
	}

	return reflect.Value{}
}

func (v Visitor) compareNumber(op string, a, b reflect.Value) reflect.Value {
	switch op {
	case "==":
		return reflect.ValueOf(toFloat(a) == toFloat(b))
	case "!=":
		return reflect.ValueOf(toFloat(a) != toFloat(b))
	case "<":
		return reflect.ValueOf(toFloat(a) < toFloat(b))
	case "<=":
		return reflect.ValueOf(toFloat(a) <= toFloat(b))
	case ">":
		return reflect.ValueOf(toFloat(a) > toFloat(b))
	case ">=":
		return reflect.ValueOf(toFloat(a) >= toFloat(b))
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

	return reflect.Value{}
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

	if err != nil {
		v.AddError(fmt.Errorf("convert: %s to decimal error: %w", ctx.GetText(), err))
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

func isNumber(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func toFloat(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	default:
		return 0
	}
}
