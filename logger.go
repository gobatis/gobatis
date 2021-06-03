package gobatis

import (
	"github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"
	"log"
)

var (
	_level         = Debug
	_logger Logger = newDefaultLogger()
)

func Debugf(format string, args ...interface{}) {
	if _level == Debug {
		_logger.Debugf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if _level <= Info {
		_logger.Infof(format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if _level <= Warn {
		_logger.Warnf(format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}

type LogLevel = int

const (
	Debug = iota
	Info
	Warn
	Error
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func newDefaultLogger() *defaultLogger {
	return &defaultLogger{
		logger: logrus.New(),
	}
}

type defaultLogger struct {
	logger *logrus.Logger
}

func (p defaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf(chalk.Cyan.Color("[DEBUG]")+format, args...)
}

func (p defaultLogger) Infof(format string, args ...interface{}) {
	log.Printf(chalk.Green.Color("[INFO]")+format, args...)
}

func (p defaultLogger) Warnf(format string, args ...interface{}) {
	log.Printf(chalk.Yellow.Color("[WARN]")+format, args...)
}

func (p defaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(chalk.Red.Color("[ERROR]")+format, args...)
}
