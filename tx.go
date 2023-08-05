package batis

import (
	"context"
	"database/sql"
)

func (d *DB) Tx() *sql.Tx {
	return d.tx
}

func (d *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return d.tx.PrepareContext(ctx, query)
}

//func (d *DB) addStmt(stmt *Stmt) {
//	d.mu.Lock()
//	defer func() {
//		d.mu.Unlock()
//	}()
//	if d.stmtMap == nil {
//		d.stmtMap = map[string]*Stmt{}
//	}
//	_, ok := d.stmtMap[stmt.caller.fragment.id]
//	if ok {
//		return
//	}
//	d.stmtMap[stmt.caller.fragment.id] = stmt
//}
//
//func (d *DB) getStmt(id string) *Stmt {
//	d.mu.RLock()
//	defer func() {
//		d.mu.RUnlock()
//	}()
//	return d.stmtMap[id]
//}

func (d *DB) Commit() error {
	return d.tx.Commit()
}
func (d *DB) Rollback() error {
	return d.tx.Rollback()
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.tx.ExecContext(ctx, query, args...)
}

func (d *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.tx.QueryContext(ctx, query, args...)
}
