package gobatis

import (
	"context"
	"database/sql"
	"go.uber.org/atomic"
)

type DB struct {
	count atomic.Uint64
	db    *sql.DB
}

func (p *DB) PingContext() {

}
func (p *DB) Ping() {

}
func (p *DB) Close() {

}
func (p *DB) SetMaxIdleConns() {

}
func (p *DB) SetMaxOpenConns() {

}
func (p *DB) SetConnMaxLifetime() {

}
func (p *DB) SetConnMaxIdleTime() {

}
func (p *DB) Stats() {

}
func (p *DB) PrepareContext() {

}
func (p *DB) Prepare() {

}
func (p *DB) ExecContext() {

}
func (p *DB) Exec() {

}
func (p *DB) QueryContext() {

}
func (p *DB) Query() {

}
func (p *DB) QueryRowContext() {

}
func (p *DB) QueryRow() {

}
func (p *DB) BeginTx() {

}
func (p *DB) Begin() {

}
func (p *DB) Driver() {

}

func (p *DB) Conn(ctx context.Context) (conn *sql.Conn, err error) {
	return p.db.Conn(ctx)
}
