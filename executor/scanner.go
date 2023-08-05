package executor

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

func WithErrScanner(err error) Scanner {
	return Scanner{err: err}
}

type Scanner struct {
	err    error
	ctx    context.Context
	rows   []*sql.Rows
	must   bool
	debug  bool
	result []*sql.Result
}

func (s Scanner) Scan(ptr ...any) error {
	if s.err != nil {
		return s.err
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
		err := qr.scan(reflect.ValueOf(ptr[i]))
		if err != nil {
			return fmt.Errorf("scan rows error: %s", err)
		}
	}
	
	return nil
}

func (s Scanner) Error() error {
	return s.err
}

func (s Scanner) AffectRows() (int, error) {
	return 0, nil
}
