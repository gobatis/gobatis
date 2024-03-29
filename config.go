package batis

import (
	"database/sql"
	"time"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/logger"
)

type Option interface {
	Apply(*Config) error
	AfterInitialize(*DB) error
}

var _ Option = (*Config)(nil)

type Config struct {
	//CreateBatchSize int
	Plugins   map[string]Plugin
	NowFunc   func() time.Time
	Dialector dialector.Dialector
	Logger    logger.Logger
	Hooks     func(db *DB)
	ColumnTag string
	db        *sql.DB
}

func (c Config) clone() *Config {
	return &Config{
		Plugins:   c.Plugins,
		NowFunc:   c.NowFunc,
		Dialector: c.Dialector,
		Logger:    c.Logger,
		Hooks:     c.Hooks,
		ColumnTag: c.ColumnTag,
		db:        c.db,
	}
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
