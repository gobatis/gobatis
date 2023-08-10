package batis

import "database/sql"

type Dialector interface {
	DB() (*sql.DB, error)
	WrapName(name string) string
}
