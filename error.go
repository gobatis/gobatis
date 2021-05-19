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
	syntaxErr
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

func (p *_error) Error() string {
	msg := fmt.Sprintf("ERROR %d: %s", p.code, p.message)
	if p.ctx != nil {
		msg += fmt.Sprintf("\n%s line %d column %d:\n%s",
			p.file, p.ctx.GetStart().GetLine(), p.ctx.GetStart().GetColumn(), p.ctx.GetText())
	}
	return msg
}

func newParserErrorStrategy() *parserErrorStrategy {
	return &parserErrorStrategy{BailErrorStrategy: antlr.NewBailErrorStrategy()}
}

type parserErrorStrategy struct {
	*antlr.BailErrorStrategy
}

func (p *parserErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	context := recognizer.GetParserRuleContext()
	//for context != nil {
	//	context.SetException(e)
	//	context = context.GetParent().(antlr.ParserRuleContext)
	//}
	//panic(NewParseCancellationException()) // TODO we don't emit e properly
	throw("", context, syntaxErr).format("syntax error")
}

func (p *parserErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	p.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

func (p *parserErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}

func (p *parserErrorStrategy) ReportError(antlr.Parser, antlr.RecognitionException) {
	// pass
}

func (p *parserErrorStrategy) ReportMatch(antlr.Parser) {
	// pass
}

func newErrorListener() *errorListener {
	return new(errorListener)
}

type errorListener struct {
}

func (p *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	throw("", nil, syntaxErr).format("词法分析错误")
}

func (p *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// pass
}

func (p *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// pass
}

func (p *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// pass
}
