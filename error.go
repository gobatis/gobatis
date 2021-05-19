package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	unknownErr = iota + 1
	parameterTypeErr
	parameterConflictWithBuiltInErr
	parameterNotFoundErr
	varToReflectKindErr
	varToAliasErr
	popStackErr
	popBinaryOperandsErr
	popResultErr
	unsupportedRelationCalcErr
	unsupportedUnaryCalc
	unsupportedNumericCalc
	numericCalcErr
	unaryCalcError
	relationCalcError
	logicCalcErr
	visitMemberErr
	visitMapErr
	visitArrayErr
	indexErr
	callErr
	parseIntegerErr
	parseDecimalErr
	parseCoveredErr
	checkParameterErr
	checkResultErr
	parseMapperErr
	validateXMLNodeErr
	parasFragmentErr
	callerErr
)

func throw(file string, ctx antlr.ParserRuleContext, code int) *_error {
	return &_error{file: file, ctx: ctx, code: code}
}

type _error struct {
	code    int
	file    string
	ctx     antlr.ParserRuleContext
	message string
}

func (p *_error) format(format string, args ...interface{}) {
	p.message = fmt.Sprintf(format, args...)
	panic(p)
}

func (p *_error) with(err error) {
	// TODO
	// if err is _error, contact ctx
	p.message = err.Error()
	panic(p)
}

func (p *_error) slient() {

}

func (p *_error) Error() string {
	msg := fmt.Sprintf("ERROR %d: %s", p.code, p.message)
	if p.ctx != nil {
		msg += fmt.Sprintf("\n%s line %d column %d:\n%s",
			p.file, p.ctx.GetStart().GetLine(), p.ctx.GetStart().GetColumn(), p.ctx.GetText())
	}
	return msg
}
