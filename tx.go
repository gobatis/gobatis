package gobatis

import (
	"context"
	"database/sql"
	"sync"
)

func NewTx(tx *sql.Tx) *Tx {
	return &Tx{tx: tx}
}

type Tx struct {
	tx      *sql.Tx
	mu      sync.RWMutex
	stmtMap map[string]*Stmt
}

func (p *Tx) Tx() *sql.Tx {
	return p.tx
}

func (p *Tx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return p.tx.PrepareContext(ctx, query)
}

func (p *Tx) addStmt(stmt *Stmt) {
	p.mu.Lock()
	defer func() {
		p.mu.Unlock()
	}()
	if p.stmtMap == nil {
		p.stmtMap = map[string]*Stmt{}
	}
	_, ok := p.stmtMap[stmt.caller.fragment.id]
	if ok {
		return
	}
	p.stmtMap[stmt.caller.fragment.id] = stmt
}

func (p *Tx) getStmt(id string) *Stmt {
	p.mu.RLock()
	defer func() {
		p.mu.RUnlock()
	}()
	return p.stmtMap[id]
}

func (p *Tx) Commit() error {
	return p.tx.Commit()
}
func (p *Tx) Rollback() error {
	return p.tx.Rollback()
}

func (p *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.tx.ExecContext(ctx, query, args...)
}

func (p *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.tx.QueryContext(ctx, query, args...)
}
