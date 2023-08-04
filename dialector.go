package gobatis

import "database/sql"

type Dialector interface {
	DB() (*sql.DB, error)
	//WrapColumn(name string) string
}
