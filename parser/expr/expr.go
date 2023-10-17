package expr

import (
	"fmt"
	"reflect"

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

func Parse(source string, vars map[string]any) (any, error) {

	errs := &commons.ErrorListener{}

	//source = replaceIsolatedLessThanWithEntity(source)
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
	if errs.GetError() != nil {
		return nil, errs.GetError()
	}
	v := &Visitor{
		errs: errs,
		vars: map[string]reflect.Value{},
	}
	for kk, vv := range vars {
		v.vars[kk] = valueOf(vv)
	}
	return v.visitExpressions(tree.(*ExpressionsContext))
}

type Visitor struct {
	errs *commons.ErrorListener
	vars map[string]reflect.Value
}

func (v Visitor) visitExpressions(ctx *ExpressionsContext) (r any, err error) {
	if ctx.Expression() == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	defer func() {
		err = v.errs.GetError()
	}()
	rv := v.visitExpression(ctx.Expression().(*ExpressionContext))
	if rv.Kind() != reflect.Invalid {
		// TODO handle valuer
		r = rv.Interface()
	}
	return
}

func (v Visitor) visitExpression(ctx *ExpressionContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	if ctx.Primary() != nil {
		return v.visitPrimary(ctx.Primary().(*PrimaryContext))
	} else if ctx.GetUnary() != nil {
		rv := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		return v.visitUnary(ctx.GetUnary().GetText(), rv)
	} else if ctx.GetRel() != nil {
		a := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		b := v.visitExpression(ctx.Expression(1).(*ExpressionContext))
		return v.visitRel(ctx.GetRel().GetText(), a, b)
	} else if ctx.Logical() != nil {
		a := v.visitExpression(ctx.Expression(0).(*ExpressionContext))
		b := v.visitExpression(ctx.Expression(1).(*ExpressionContext))
		return v.visitLogical(ctx.GetRel().GetText(), a, b)
	} else {
		v.errs.AddError(fmt.Errorf("unsupported expression: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitPrimary(ctx *PrimaryContext) reflect.Value {
	if v.errs.GetError() != nil {
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
		} else {
			v.errs.AddError(fmt.Errorf("unsupported primary expression: %s", ctx.GetText()))
			return reflect.Value{}
		}
	} else if ctx.Operand() != nil {
		return v.visitOperand(ctx.Operand().(*OperandContext))
	} else {
		v.errs.AddError(fmt.Errorf("unsupported primary expression: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitOperand(ctx *OperandContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	if ctx.Literal() != nil {
		return v.visitLiteral(ctx.Literal().(*LiteralContext))
	} else if ctx.Var_() != nil {
		return v.visitVar(ctx.Var_().(*VarContext))
	} else if ctx.Expression() != nil {
		return v.visitExpression(ctx.Expression().(*ExpressionContext))
	} else {
		v.errs.AddError(fmt.Errorf("unsupported operand: %s", ctx.GetText()))
		return reflect.Value{}
	}
}

func (v Visitor) visitLiteral(ctx *LiteralContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}

	if ctx.Nil_() != nil {

	} else if ctx.Integer() != nil {
		vv, err := decimal.NewFromString(ctx.Integer().GetText())
		if err != nil {
			v.errs.AddError(fmt.Errorf("convert: %s to decimal error: %w", ctx.Integer().GetText(), err))
			return reflect.Value{}
		}
		return valueOf(vv)
	} else if ctx.String_() != nil {
		return valueOf(ctx.String_().GetText()[1 : len(ctx.String_().GetText())-1])
	} else if ctx.Float() != nil {
		// TODO convert number
		vv, err := decimal.NewFromString(ctx.Float().GetText())
		if err != nil {
			v.errs.AddError(fmt.Errorf("convert: %s to decimal error: %w", ctx.Float().GetText(), err))
			return reflect.Value{}
		}
		return valueOf(vv)
	}

	return reflect.Value{}
}

func (v Visitor) visitVar(ctx *VarContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	vv, ok := v.vars[ctx.GetText()]
	if !ok {
		v.errs.AddError(fmt.Errorf("var: %s is not defined", ctx.GetText()))
		return reflect.Value{}
	}
	return vv
}

func (v Visitor) visitMember(rv reflect.Value, ctx *MemberContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	if rv.Kind() != reflect.Struct {
		v.errs.AddError(fmt.Errorf("var: %s expect struct, got: %s", ctx.GetText(), rv.Kind()))
		return reflect.Value{}
	}
	fv := rv.FieldByName(ctx.IDENTIFIER().GetText())
	if fv.Kind() == reflect.Invalid {
		v.errs.AddError(fmt.Errorf("visit member: %s not exist", ctx.IDENTIFIER().GetText()))
		return reflect.Value{}
	}
	return fv
}

func (v Visitor) visitIndex(rv reflect.Value, ctx *IndexContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	iv := v.visitExpression(ctx.Expression().(*ExpressionContext))

	// TODO
	i, err := cast.ToIntE(iv.Interface().(decimal.Decimal).IntPart())
	if err != nil {
		v.errs.AddError(fmt.Errorf("covert index error: %s", err))
		return reflect.Value{}
	}

	rv = rv.Index(i)
	if rv.Kind() == reflect.Invalid {
		v.errs.AddError(fmt.Errorf("invalid index: %d", i))
		return reflect.Value{}
	}

	return rv
}

func (v Visitor) visitSlice(rv reflect.Value, ctx *SliceContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitUnary(op string, rv reflect.Value) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitRel(op string, a, b reflect.Value) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitLogical(op string, a, b reflect.Value) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitInteger(ctx *IntegerContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitString(ctx *StringContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func (v Visitor) visitFloat(ctx *FloatContext) reflect.Value {
	if v.errs.GetError() != nil {
		return reflect.Value{}
	}
	return reflect.Value{}
}
