package batis

import (
	"context"
	"database/sql"
)

func NewTx(tx *sql.Tx, traceID string) *Tx {
	return &Tx{Tx: tx, traceId: traceID}
}

var _ conn = (*Tx)(nil)

type Tx struct {
	traceId string
	*sql.Tx
}

func (t *Tx) TraceId() string {
	return t.traceId
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
