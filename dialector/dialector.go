package dialector

import "database/sql"

type Dialector interface {
	DB() (*sql.DB, error)
	Namer() Namer
	Explain(sql string, vars ...interface{}) string
	//DataTypeOf(field *schema.Field) string
	//BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{})
}

type Namer interface {
	ReservedName(name string) string
	ColumnName(name string) string
	TableName(name string) string
}
