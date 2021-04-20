package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener struct {
	errors []error
}

func (p *ErrorListener) HasError() bool {
	return len(p.errors) > 0
}

func (p *ErrorListener) Errors() []error {
	return p.errors
}

func (p *ErrorListener) addError(err error) {
	p.errors = append(p.errors, err)
}

func (p *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p.addError(fmt.Errorf("行 %d:%d 语法错误: %s", line, column, msg))
}

func (p *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//p.addError(fmt.Errorf("行 %d:%d Ambiguity错误: %s", stopIndex, startIndex, ambigAlts.String()))
	panic("Call ReportAmbiguity")
}

func (p *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("Call ReportAttemptingFullContext")
}

func (p *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("Call ReportContextSensitivity")
}
