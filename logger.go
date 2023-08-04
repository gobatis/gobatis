package gobatis

import (
	"github.com/gozelle/logging"
	"github.com/gozelle/zap"
)

type Level = int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

type Logger interface {
	SetModule(name string)
	SetLevel(level Level)
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Sync() error
}

func newLogger() Logger {
	return &logger{
		log: logging.Logger("[gobatis]").SugaredLogger,
	}
}

type logger struct {
	log zap.SugaredLogger
}

func (l logger) SetModule(name string) {
	//l.log.SugaredLogger
}

func (l logger) SetLevel(level Level) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l logger) Sync() error {
	return l.log.Sync()
}
