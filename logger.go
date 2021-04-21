package gobatis

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func newLogger() *logger {
	return &logger{
		logger: logrus.New(),
	}
}

type logger struct {
	logger *logrus.Logger
}

func (p logger) Debugf(format string, args ...interface{}) {
	p.logger.Debugf(format, args...)
}

func (p logger) Infof(format string, args ...interface{}) {
	p.logger.Infof(format, args...)
}

func (p logger) Warnf(format string, args ...interface{}) {
	p.logger.Warnf(format, args...)
}

func (p logger) Errorf(format string, args ...interface{}) {
	p.logger.Errorf(format, args...)
}
