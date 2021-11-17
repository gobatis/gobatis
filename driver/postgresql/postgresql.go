package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"reflect"
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
	var assigner pgtype.Value
	
	switch ct.DatabaseTypeName() {
	case "INT4":
		assigner = new(pgtype.Int4)
	case "INT8":
		assigner = new(pgtype.Int8)
	case "VARCHAR":
		assigner = new(pgtype.Varchar)
	case "BPCHAR":
		assigner = new(pgtype.BPChar)
	case "BOOL":
		assigner = new(pgtype.Bool)
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
		err = fmt.Errorf("scan %s err: %s", ct.Name(), err)
		return
	}
	if assigner.Get() != nil {
		err = assigner.AssignTo(value.Interface())
		if err != nil {
			err = fmt.Errorf("assign %s err: %s", ct.Name(), err)
			return
		}
	}
	return
}
