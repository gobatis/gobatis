package testdb

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/dialector"
	_ "github.com/proullon/ramsql/driver"
	"regexp"
	"strings"
)

var _ dialector.Dialector = (*Dialector)(nil)

func Open(dsn string) *Dialector {
	d := &Dialector{}
	d.db, d.err = sql.Open("ramsql", "TestLoadUserAddresses")
	return d
}

type Dialector struct {
	db  *sql.DB
	err error
}

func (d Dialector) DB() (*sql.DB, error) {
	return d.db, d.err
}

func (d Dialector) Namer() dialector.Namer {
	return Namer{}
}

var _ dialector.Namer = (*Namer)(nil)

type Namer struct {
}

func (n Namer) TableName(name string) string {
	items := strings.Split(name, ".")
	for i := range items {
		items[i] = n.ReservedName(items[i])
	}
	return strings.Join(items, ".")
}

func (n Namer) ReservedName(name string) string {
	if strings.HasPrefix(name, "\"") {
		return name
	}
	return fmt.Sprintf("\"%s\"", name)
}

func (n Namer) ColumnName(name string) string {
	var re = regexp.MustCompile(`([^A-Z_])([A-Z])`)
	snakeStr := re.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(snakeStr)
}
