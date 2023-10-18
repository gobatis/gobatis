package executor

import (
	"errors"
)

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrInvalidTransaction = errors.New("invalid transaction")
	NoScanDestErr         = errors.New("expect 1 scan dest, got nil")
)

const (
	unknownErr = iota + 1
	syntaxErr
)
