package hook

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

type BeforeExecute func(ctx BeforeContext)

type AfterExecute func(ctx AfterContext)
