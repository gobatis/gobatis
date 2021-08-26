package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/parser/xml"
)

const (
	unknownErr = iota + 1
	parameterTypeErr
	parameterConflictWithBuiltInErr
	parameterNotFoundErr
	varToReflectKindErr
	varBindErr
	popValueStackErr
	popParamsStackErr
	popBinaryOperandsErr
	popTertiaryOperandsErr
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
	resultAttributeConflictErr
	castBoolErr
)

func throw(file string, ctx antlr.ParserRuleContext, code int) *_error {
	return &_error{file: file, ctx: ctx, code: code}
}

type _error struct {
	code    int
	file    string
	parent  antlr.ParserRuleContext
	ctx     antlr.ParserRuleContext
	message string
}

func (p *_error) setParent(ctx antlr.ParserRuleContext) *_error {
	p.parent = ctx
	return p
}

func (p *_error) format(format string, args ...interface{}) {
	p.message = fmt.Sprintf(format, args...)
	panic(p)
}

func (p *_error) with(err error) {
	p.message = err.Error()
	panic(p)
}

func (p *_error) Error() string {
	msg := fmt.Sprintf("[ERROR %d]: %s", p.code, p.message)
	line := 0
	column := 0
	ctx := ""
	if p.parent != nil {
		line = p.parent.GetStart().GetLine()
		column = p.parent.GetStart().GetColumn()
		ctx = getText(p.parent)
	} else if p.ctx != nil {
		line = p.ctx.GetStart().GetLine()
		column = p.ctx.GetStart().GetColumn()
		ctx = getText(p.ctx)
	}
	if p.ctx != nil {
		msg += fmt.Sprintf("\n[file]: %s near line %d column %d:\n[context]: %s", p.file, line, column+1, ctx)
	}
	
	return msg
}

func castRecoverError(file string, e interface{}) error {
	if e != nil {
		_e, ok := e.(*_error)
		if ok {
			if _e.file == "" {
				_e.file = file
			}
			return _e
		}
		return &_error{
			code:    unknownErr,
			file:    file,
			message: fmt.Sprintf("%v", e),
		}
	}
	return nil
}

func getText(ctx antlr.ParserRuleContext) string {
	
	if ctx.GetChildCount() == 0 {
		return ""
	}
	
	var s string
	for _, child := range ctx.GetChildren() {
		_, ok := child.(*xml.AttributeContext)
		if ok {
			s += " " + child.(antlr.ParseTree).GetText()
		} else {
			s += child.(antlr.ParseTree).GetText()
		}
	}
	
	return s
}

func newParserErrorStrategy() *parserErrorStrategy {
	return &parserErrorStrategy{BailErrorStrategy: antlr.NewBailErrorStrategy()}
}

type parserErrorStrategy struct {
	*antlr.BailErrorStrategy
}

func (p *parserErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	// TODO handle syntax error detail
	context := recognizer.GetParserRuleContext()
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
