package dialector

import "database/sql"

type Dialector interface {
	DB() (*sql.DB, error)
	Namer() Namer
}

type Namer interface {
	ReservedName(name string) string
	ColumnName(name string) string
	TableName(name string) string
}
