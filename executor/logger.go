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
	Trace(pos, id string, tx bool, err error, st *SQLTrace)
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

func (l logger) Trace(pos, traceId string, tx bool, err error, tr *SQLTrace) {
	if !tr.Trace && !tr.Debug && err == nil {
		return
	}
	info := &strings.Builder{}
	var f func(format string, a ...any)
	if err != nil {
		f = l.Errorf
	} else {
		f = l.Debugf
	}
	if traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(traceId))
	}
	var t string
	if tx {
		t = fmt.Sprintf("[%s]", color.MagentaString("tx"))
	}
	if pos == "" {
		pos = CallFuncPos(5)
	}
	info.WriteString(fmt.Sprintf("%s%s %s", traceId, t, color.RedString(pos)))
	if err != nil {
		info.WriteString(color.RedString(fmt.Sprintf(" ERROR: %s", err.Error())))
	}
	
	if tr != nil {
		cost := time.Since(tr.BeginAt)
		info.WriteString(fmt.Sprintf("\n%s %s %s",
			color.YellowString(fmt.Sprintf("[%s]", cost)),
			color.BlueString(fmt.Sprintf("[rows:%d]", tr.RowsAffected)),
			tr.RawSQL,
		))
	}
	f(info.String())
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

type SQLTrace struct {
	Trace        bool
	Debug        bool
	BeginAt      time.Time
	RawSQL       string
	PlainSQL     string
	RowsAffected int64
}

type TraceSQL struct {
}

// CallFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func CallFuncPos(skip int) string {
	i := skip
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok || i > 20 {
			break
		}
		if !strings.Contains(file, "/gobatis/gobatis") || strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		i++
	}
	return ""
}
