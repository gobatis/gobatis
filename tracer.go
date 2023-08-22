package batis

import "time"

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
	dynamic bool
	append  bool
}

func (t tracer) log() {
	if !t.debug && t.err == nil {
		return
	}
}
