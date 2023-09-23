package commons

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var _ antlr.ErrorListener = (*CustomErrorListener)(nil)

type CustomErrorListener struct {
	err error
}

func (d *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("ambiguity detected between tokens %d and %d. Ambiguous alternatives: %v", startIndex, stopIndex, ambigAlts))
}

func (d *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("attempting full context mode between tokens %d and %d", startIndex, stopIndex))
}

func (d *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	d.AddError(fmt.Errorf("context sensitivity detected between tokens %d and %d", startIndex, stopIndex))
}

func (d *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.AddError(fmt.Errorf("syntax error at line %d:%d - %s", line, column, msg))
}

func (d *CustomErrorListener) Error() error {
	return d.err
}

func (d *CustomErrorListener) AddError(err error) {
	d.err = AddError(d.err, err)
}

func AddError(err, added error) error {
	if err == nil {
		return added
	} else if added != nil {
		return fmt.Errorf("%v; %w", err, added)
	}
	return err
}
