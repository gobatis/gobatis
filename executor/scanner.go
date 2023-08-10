package executor

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

type Scanner struct {
	Error  error
	ctx    context.Context
	rows   []*sql.Rows
	must   bool
	debug  bool
	result []*sql.Result
}

func (s Scanner) Scan(ptr ...any) (err error) {
	
	defer func() {
		if err != nil {
			s.printError()
		}
	}()
	
	if s.Error != nil {
		err = s.Error
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
		err = qr.scan(reflect.ValueOf(ptr[i]))
		if err != nil {
			return fmt.Errorf("scan rows error: %s", err)
		}
	}
	
	return nil
}

func (s Scanner) printError() {
	debugLog("****", s.Error)
}

func (s Scanner) AffectRows() (affectedRows int, err error) {
	defer func() {
		if err != nil {
			s.printError()
		}
	}()
	
	if s.Error != nil {
		err = s.Error
		return
	}
	
	return 0, nil
}
