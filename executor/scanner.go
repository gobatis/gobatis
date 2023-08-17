package executor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gozelle/logger"
)

func NewErrorScanner(err error) Scanner {
	return Scanner{err: err}
}

type Scanner struct {
	err    error
	ctx    context.Context
	logger logger.Logger
	rows   []*sql.Rows
	must   bool
	debug  bool
	result []*sql.Result
}

func (s Scanner) Error() error {
	return s.err
}

func (s Scanner) Scan(ptr ...any) (err error) {

	defer func() {
		if err != nil {
			//s.printError()
		}
	}()

	if s.err != nil {
		err = s.err
		return
	}

	l1 := len(ptr)
	l2 := len(s.rows)
	if l1 > l2 {
		return fmt.Errorf("the receiving result ptrs length: %d > result length: %d", l1, l2)
	}

	for i := 0; i < l2; i++ {
		qr := queryResult{
			rows: s.rows[i],
		}
		err = qr.scan(ptr[i])
		if err != nil {
			return fmt.Errorf("scan rows error: %s", err)
		}
	}

	return nil
}

//func (s Scanner) printError() {
//	log("****", s.err)
//}

func (s Scanner) AffectRows() (affectedRows int, err error) {
	defer func() {
		if err != nil {
			//s.printError()
		}
	}()

	if s.err != nil {
		err = s.err
		return
	}

	return 0, nil
}
