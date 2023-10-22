package parser

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
	antlr.ConsoleErrorListener
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.AddError(fmt.Errorf("syntax error at line %d:%d - %s", line, column, msg))
}

func (d *ErrorListener) Error() error {
	return d.err
}

func (d *ErrorListener) AddError(err error) {
	d.err = AddError(d.err, err)
}

var _ antlr.ErrorStrategy = (*ErrorStrategy)(nil)

type ErrorStrategy struct {
	*antlr.DefaultErrorStrategy
}

func RecoverError(e any) error {
	err, ok := e.(error)
	if ok {
		return err
	} else {
		return fmt.Errorf("panic: %v", err)
	}
}
