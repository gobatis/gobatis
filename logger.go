package batis

import (
	syslog "log"
)

type Logger interface {
	Debugf(format string, a ...any)
	Infof(format string, a ...any)
	Errorf(format string, a ...any)
	Warnf(format string, a ...any)
	//Trace( begin time.Time, fc func() (sql string, rowsAffected int64), err error)
}

func DefaultLogger() Logger {
	return &logger{}
}

var _ Logger = (*logger)(nil)

type logger struct {
}

func (l logger) Debugf(format string, a ...any) {
	syslog.Printf(format, a...)
}

func (l logger) Infof(format string, a ...any) {
	syslog.Printf(format, a...)
}

func (l logger) Errorf(format string, a ...any) {
	syslog.Printf(format, a...)
}

func (l logger) Warnf(format string, a ...any) {
	syslog.Printf(format, a...)
}
