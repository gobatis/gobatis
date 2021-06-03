package gobatis

import (
	"github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"
	"log"
)

var (
	log_level = Debug
)

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

func newLogger() *logger {
	return &logger{
		logger: logrus.New(),
	}
}

type logger struct {
	logger *logrus.Logger
}

func (p logger) Debugf(format string, args ...interface{}) {
	if log_level == Debug {
		log.Printf(chalk.Cyan.Color("[DEBUG]")+format, args...)
	}
}

func (p logger) Infof(format string, args ...interface{}) {
	if log_level <= Info {
		log.Printf(chalk.Green.Color("[INFO]")+format, args...)
	}
}

func (p logger) Warnf(format string, args ...interface{}) {
	if log_level <= Warn {
		log.Printf(chalk.Yellow.Color("[WARN]")+format, args...)
	}
}

func (p logger) Errorf(format string, args ...interface{}) {
	log.Printf(chalk.Red.Color("[ERROR]")+format, args...)
}
