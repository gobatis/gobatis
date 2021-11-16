package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"reflect"
	"strings"
)

const PGX = "pgx"

func InitDB(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open(PGX, dsn)
	if err != nil {
		err = fmt.Errorf("postgresql connnet error: %s", err)
		return
	}
	return
}

func NewEngine(dsn string) *gobatis.Engine {
	db, err := InitDB(dsn)
	if err != nil {
		panic(err)
	}
	engine := gobatis.NewEngine(gobatis.NewDB(db))
	engine.SetScannerFactory(ScannerFactory)
	return engine
}

func ScannerFactory() gobatis.Scanner {
	return new(Scanner)
}

type Assigner interface {
	AssignTo(interface{}) error
}

type Scanner struct {
	fields map[string]*sql.ColumnType
}

func (s *Scanner) Scan(rows *sql.Rows, ct *sql.ColumnType, value reflect.Value) (err error) {
	if !strings.HasPrefix(ct.DatabaseTypeName(), "_") {
		err = rows.Scan(value.Interface())
		if err != nil {
			return
		}
	} else {
		var assigner Assigner
		switch ct.DatabaseTypeName() {
		case "_INT8":
			assigner = new(pgtype.Int8Array)
		case "_VARCHAR":
			assigner = new(pgtype.VarcharArray)
		case "_BPCHAR":
			assigner = new(pgtype.BPCharArray)
		default:
			err = fmt.Errorf("unsupport scan type: %s(%s)", ct.Name(), ct.DatabaseTypeName())
			return
		}
		err = rows.Scan(assigner)
		if err != nil {
			return
		}
		err = assigner.AssignTo(value.Interface())
		if err != nil {
			return
		}
	}
	return
}
