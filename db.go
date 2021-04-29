package gobatis

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"time"
)

func NewDB(driver, dsn string) *DB {
	return &DB{
		driver: driver,
		dsn:    dsn,
	}
}

type DB struct {
	driver string
	dsn    string
	db     *sql.DB
}

func (p *DB) initDB() (err error) {
	switch p.driver {
	case postgresql.PGX:
		p.db, err = postgresql.InitDB(p.dsn)
	case mysql.MySQL:
		p.db, err = mysql.InitDB(p.dsn)
	default:
		p.db, err = sql.Open(p.driver, p.dsn)
		if err != nil {
			err = fmt.Errorf("%s connnet error: %s", p.driver, err)
			return
		}
	}
	p.dsn = ""
	return
}

func (p *DB) PingContext(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *DB) Ping() error {
	return p.db.Ping()
}

func (p *DB) Close() error {
	return p.db.Close()
}

func (p *DB) SetMaxIdleConns(n int) {
	p.db.SetMaxIdleConns(n)
}
func (p *DB) SetMaxOpenConns(n int) {
	p.db.SetMaxOpenConns(n)
}
func (p *DB) SetConnMaxLifetime(d time.Duration) {
	p.db.SetConnMaxLifetime(d)
}
func (p *DB) SetConnMaxIdleTime(d time.Duration) {
	p.db.SetConnMaxIdleTime(d)
}
func (p *DB) Stats() sql.DBStats {
	return p.db.Stats()
}

func (p *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return p.db.PrepareContext(ctx, query)
}

func (p *DB) Prepare(query string) (*sql.Stmt, error) {
	return p.db.Prepare(query)
}

func (p *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(query, args...)
}

func (p *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

func (p *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args...)
}

func (p *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

func (p *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.db.QueryRow(query, args...)
}

func (p *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return p.db.BeginTx(ctx, opts)
}

func (p *DB) Begin() (*sql.Tx, error) {
	return p.db.Begin()
}

func (p *DB) Driver() driver.Driver {
	return p.db.Driver()
}

func (p *DB) Conn(ctx context.Context) (conn *sql.Conn, err error) {
	return p.db.Conn(ctx)
}
