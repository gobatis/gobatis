package batis

import "context"

type Context struct {
	ctx   context.Context
	debug bool
	count uint64
	must  bool
	tx    *Tx
}

func WithContext(ctx context.Context) *Context {
	return &Context{}
}
