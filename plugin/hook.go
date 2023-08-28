package plugin

import "time"

type BeforeContext struct {
	ID    int64
	Raw   string
	Vars  string
	Query bool
	Tx    bool
}

type AfterContext struct {
	BeforeContext
	Error      error
	Cost       time.Time
	AffectRows int64
}

type Plugin interface {
	BeforeExecute(ctx BeforeContext)
	AfterExecute(ctx AfterContext)
}
