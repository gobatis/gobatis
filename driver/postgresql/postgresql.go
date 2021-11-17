package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis"
	"github.com/gozelle/decimal"
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
	if value.Type().Implements(gobatis.ScannerType) {
		err = rows.Scan(value.Interface())
		if err != nil {
			err = fmt.Errorf("scan %s err: %s", ct.Name(), err)
			return
		}
	} else {
		var assigner pgtype.Value
		switch ct.DatabaseTypeName() {
		case "INT2":
			assigner = new(pgtype.Int2)
		case "INT4":
			assigner = new(pgtype.Int4)
		case "INT8":
			assigner = new(pgtype.Int8)
		case "FLOAT4":
			assigner = new(pgtype.Float4)
		case "FLOAT8":
			assigner = new(pgtype.Float8)
		case "NUMERIC":
			assigner = new(pgtype.Numeric)
		case "VARCHAR":
			assigner = new(pgtype.Varchar)
		case "BPCHAR":
			assigner = new(pgtype.BPChar)
		case "TEXT":
			assigner = new(pgtype.Text)
		case "BOOL":
			assigner = new(pgtype.Bool)
		case "TIME":
			assigner = new(pgtype.Time)
		case "TIMESTAMP":
			assigner = new(pgtype.Timestamp)
		case "_INT2":
			assigner = new(pgtype.Int2Array)
		case "_INT4":
			assigner = new(pgtype.Int4Array)
		case "_INT8":
			assigner = new(pgtype.Int8Array)
		case "_FLOAT4":
			assigner = new(pgtype.Float4Array)
		case "_FLOAT8":
			assigner = new(pgtype.Float8Array)
		case "_NUMERIC":
			assigner = new(pgtype.NumericArray)
		case "_VARCHAR":
			assigner = new(pgtype.VarcharArray)
		case "_BPCHAR":
			assigner = new(pgtype.BPCharArray)
		case "_TEXT":
			assigner = new(pgtype.TextArray)
		case "_BOOL":
			assigner = new(pgtype.BoolArray)
		default:
			err = fmt.Errorf("unsupport scan type: %s(%s)", ct.Name(), ct.DatabaseTypeName())
			return
		}
		err = rows.Scan(assigner)
		if err != nil {
			err = fmt.Errorf("scan %s(%s) err: %s", ct.Name(), ct.DatabaseTypeName(), err)
			return
		}
		if assigner.Get() != nil {
			if ct.DatabaseTypeName() == "_NUMERIC" && strings.HasSuffix(value.Type().String(), "decimal.Decimal") {
				for _, v := range assigner.(*pgtype.NumericArray).Elements {
					var r float64
					err = v.AssignTo(&r)
					if err != nil {
						err = fmt.Errorf("assign %s(%s) to decimal err: %s", ct.Name(), ct.DatabaseTypeName(), err)
						return
					}
					d := decimal.NewFromFloat(r)
					value.Elem().Set(reflect.Append(value.Elem(), reflect.ValueOf(d)))
				}
			} else {
				err = assigner.AssignTo(value.Interface())
			}
			if err != nil {
				err = fmt.Errorf("assign %s(%s) err: %s", ct.Name(), ct.DatabaseTypeName(), err)
				return
			}
		}
	}
	return
}