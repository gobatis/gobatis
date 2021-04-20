package gobatis

import (
	"go.uber.org/atomic"
	"sync"
)

type Stat struct {
	Count atomic.Uint64
	Sent  atomic.Uint64
	Recv  atomic.Uint64
}

type StatSQL struct {
	id      atomic.Uint64
	Alias   sync.Map
	Cost    sync.Map
	Timeout sync.Map
	Error   sync.Map
}

func (p *StatSQL) NewID() uint64 {
	p.id.Inc()
	return p.id.Load()
}
