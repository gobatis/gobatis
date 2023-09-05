package executor

import (
	"context"
	"database/sql"
)

func NewTx(tx *sql.Tx) *Tx {
	return &Tx{Tx: tx}
}

var _ Conn = (*Tx)(nil)

type Tx struct {
	*sql.Tx
}

func (t *Tx) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return t.Tx, nil
}

func (t *Tx) IsTx() bool {
	return true
}

func (t *Tx) Close() error {
	return nil
}
