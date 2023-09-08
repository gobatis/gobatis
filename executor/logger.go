package executor

import (
	"fmt"
	syslog "log"
	"runtime"
	"strings"
	"time"

	"github.com/gozelle/color"
)

type Logger interface {
	Debugf(format string, a ...any)
	Infof(format string, a ...any)
	Errorf(format string, a ...any)
	Warnf(format string, a ...any)
	Trace(trace, debug bool, beginAt time.Time, traceId string, tx bool, sql string, rowsAffected int64, err error)
}

func DefaultLogger() Logger {
	return &logger{}
}

var _ Logger = (*logger)(nil)

type logger struct {
}

func (l logger) brand() string {
	return ""
}

func (l logger) Trace(trace, debug bool, beginAt time.Time, traceId string, tx bool, sql string, rowsAffected int64, err error) {
	if !trace && !debug && err == nil {
		return
	}
	cost := time.Since(beginAt)
	info := &strings.Builder{}
	var out func(format string, a ...any)
	if err != nil {
		out = l.Errorf
	} else {
		out = l.Debugf
	}

	if traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(traceId))
	}
	var t string
	tx = true
	if tx {
		t = fmt.Sprintf("[%s]", color.MagentaString("transaction"))
	}
	info.WriteString(fmt.Sprintf("%s%s %s", traceId, t, color.RedString(runFuncPos(4))))
	err = fmt.Errorf("some error")
	if err != nil {
		info.WriteString(color.RedString(fmt.Sprintf(" ERROR: %s", err.Error())))
	}
	info.WriteString(fmt.Sprintf("\n%s %s %s",
		color.YellowString(fmt.Sprintf("[%s]", cost)),
		color.BlueString(fmt.Sprintf("[rows: %d]", rowsAffected)),
		sql))
	out(info.String())
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

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func runFuncPos(skip int) string {
	i := skip
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok || i > 10 {
			break
		}
		if (!strings.Contains(file, "gobatis/executor/") &&
			!strings.Contains(file, "gobatis/go")) ||
			strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		i++
	}
	return ""
}
