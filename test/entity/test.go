package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type TestEntity struct {
	Id                       int64
	Int8                     int8            `sql:"t_int8"`
	BigInt                   int64           `sql:"t_bigint"`
	Int                      int             `sql:"t_int"`
	Decimal                  decimal.Decimal `sql:"t_decimal"`
	Numeric                  decimal.Decimal `sql:"t_numeric"`
	Real                     float64         `sql:"t_real"`
	DoublePrecision          float64         `sql:"t_double_precision"`
	SmallSerial              int             `sql:"t_small_serial"`
	Serial                   int             `sql:"t_serial"`
	BigSerial                int             `sql:"t_big_serial"`
	Money                    int             `sql:"t_money"`
	Char                     int             `sql:"t_char"`
	Text                     int             `sql:"t_text"`
	TimestampWithoutTimeZone time.Time       `sql:"t_timestamp_without_time_zone"`
	TimestampWithTimeZone    time.Time       `sql:"t_timestamp_with_time_zone"`
	Date                     time.Time       `sql:"t_date"`
	TimeWithoutTimeZone      time.Time       `sql:"t_time_without_time_zone"`
	TimeWithTimeZone         time.Time       `sql:"t_time_with_time_zone"`
	Interval                 time.Duration   `sql:"t_interval"`
	Boolean                  bool            `sql:"t_boolean"`
}
