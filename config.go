package batis

import (
	"database/sql"
	"time"

	"github.com/gobatis/gobatis/dialector"
)

type Option interface {
	Apply(*Config) error
	AfterInitialize(*DB) error
}

var _ Option = (*Config)(nil)

type Config struct {
	CreateBatchSize int
	Plugins         map[string]Plugin
	NowFunc         func() time.Time
	Dialector       dialector.Dialector
	Logger          Logger
	db              *sql.DB
}

func (c Config) Apply(config *Config) error {
	//TODO implement me
	panic("implement me")
}

func (c Config) AfterInitialize(db *DB) error {
	//TODO implement me
	panic("implement me")
}

type Plugin interface {
	Name() string
	Initialize(*DB) error
}

type SavePointerDialectorInterface interface {
	SavePoint(tx *DB, name string) error
	RollbackTo(tx *DB, name string) error
}

type NamingStrategy interface {
	ColumnName(table, column string) string
}
