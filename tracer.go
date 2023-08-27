package batis

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gozelle/color"
)

type tracer struct {
	now     time.Time
	logger  Logger
	debug   bool
	must    bool
	tx      bool
	sql     string
	raw     string
	exprs   []string
	vars    []any
	traceId string
}

func (e *tracer) log(err error) {
	if !e.debug && err == nil {
		return
	}
	cost := time.Since(e.now)
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if err != nil {
		status = color.RedString("Error")
		out = e.logger.Errorf
	} else {
		status = color.GreenString("Success")
		out = e.logger.Debugf
	}
	var traceId string
	if e.traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(e.traceId))
	}
	var t string
	if e.tx {
		t = fmt.Sprintf("[%s]", color.CyanString("Tx"))
	}
	info.WriteString(fmt.Sprintf("%s %s", color.MagentaString("[gobatis]"), color.RedString(e.runFuncPos(4))))
	info.WriteString(fmt.Sprintf("\n[%s][%s]%s%s %s", status, cost, traceId, t, color.YellowString(e.sql)))
	if err != nil {
		info.WriteString(fmt.Sprintf("\n%s", color.RedString(err.Error())))
	}
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func (e *tracer) runFuncPos(skip int) string {
	i := skip
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok || i > 10 {
			break
		}
		if (!strings.Contains(file, "gobatis/executor/") &&
			!strings.Contains(file, "gobatis/db.go")) ||
			strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		i++
	}
	return ""
}
