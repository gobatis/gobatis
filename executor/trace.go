package executor

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gozelle/color"
)

type Trace struct {
	Error   error
	debug   bool
	traceId string
	tx      interface{}
	Logger  Logger
}

func (db *Trace) log() {
	if !db.debug && db.Error == nil {
		return
	}
	cost := time.Since(time.Now())
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if db.Error != nil {
		status = color.RedString("error")
		out = db.Logger.Errorf
	} else {
		status = color.GreenString("success")
		out = db.Logger.Debugf
	}
	var traceId string
	if db.traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(db.traceId))
	}
	var t string
	if db.tx != nil {
		t = fmt.Sprintf("[%s]", color.CyanString("Tx"))
	}
	info.WriteString(fmt.Sprintf("%s %s", color.MagentaString("[gobatis]"), color.RedString(db.runFuncPos(4))))
	info.WriteString(fmt.Sprintf("\n[%s][%s]%s%s %s", status, cost, traceId, t, color.YellowString("db.executor.raw")))
	if db.Error != nil {
		info.WriteString(fmt.Sprintf("\n%s", color.RedString(db.Error.Error())))
	}
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func (db *Trace) runFuncPos(skip int) string {
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
