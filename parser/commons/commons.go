package commons

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func AddError(err, added error) error {
	if err == nil {
		return added
	} else if added != nil {
		return fmt.Errorf("%v; %w", err, added)
	}
	return err
}

var _ antlr.ErrorListener = (*ErrorListener)(nil)

type ErrorListener struct {
	err error
}

func (d *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("ambiguity detected between tokens %d and %d. Ambiguous alternatives: %v", startIndex, stopIndex, ambigAlts))
}

func (d *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("attempting full context mode between tokens %d and %d", startIndex, stopIndex))
}

func (d *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("context sensitivity detected between tokens %d and %d", startIndex, stopIndex))
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.AddError(fmt.Errorf("syntax error at line %d:%d - %s", line, column, msg))
}

func (d *ErrorListener) GetError() error {
	return d.err
}

func (d *ErrorListener) AddError(err error) {
	d.err = AddError(d.err, err)
}

var _ antlr.ErrorStrategy = (*ErrorStrategy)(nil)

type ErrorStrategy struct {
	*antlr.DefaultErrorStrategy
}
