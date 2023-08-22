package batis

import (
	"fmt"
	"github.com/gozelle/color"
	"runtime"
	"strings"
	"time"
)

type tracer struct {
	now     time.Time
	traceId string
	err     error
	logger  Logger
	debug   bool
	sql     string
	raw     string
	exprs   []string
	vars    []interface{}
	tx      bool
	dynamic bool
	append  bool
}

func (t tracer) log() {
	if !t.debug && t.err == nil {
		return
	}
	cost := time.Since(t.now)
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if t.err != nil {
		status = color.RedString("Error")
		out = t.logger.Errorf
	} else {
		status = color.GreenString("Success")
		out = t.logger.Debugf
	}
	var traceId string
	if t.traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(t.traceId))
	}
	var tx string
	if t.tx {
		tx = fmt.Sprintf("[%s]", color.CyanString("Tx"))
	}
	info.WriteString(fmt.Sprintf("%s %s\n[%s][%s]%s%s", color.MagentaString("[gobatis]"), color.RedString(t.runFuncPos(4)), status, cost, traceId, tx))
	info.WriteString(fmt.Sprintf("\n%s", color.YellowString(t.sql)))
	if t.err != nil {
		info.WriteString(fmt.Sprintf("\n%s", color.RedString(t.err.Error())))
	}
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func (t tracer) runFuncPos(skip int) string {
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
