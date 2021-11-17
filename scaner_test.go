package gobatis

import (
	"database/sql"
	"github.com/gozelle/decimal"
	"time"
)

// reference: http://www.postgres.cn/docs/12/datatype.html

type PGTYPES struct {
	Int8                     int8            `sql:"t_int8"`
	Int                      int             `sql:"t_int"`
	Decimal                  decimal.Decimal `sql:"t_decimal"`
	Numeric                  decimal.Decimal `sql:"t_numeric"`
	Real                     float64         `sql:"t_real"`
	DoublePrecision          float64         `sql:"t_double_precision"`
	SmallSerial              int             `sql:"t_small_serial"`
	Serial                   int             `sql:"t_serial"`
	BigSerial                int             `sql:"t_big_serial"`
	Money                    string          `sql:"t_money"`
	Char                     string          `sql:"t_char"`
	NullChar                 sql.NullString  `sql:"t_char"`
	Text                     string          `sql:"t_text"`
	TimestampWithoutTimeZone time.Time       `sql:"t_timestamp_without_time_zone"`
	TimestampWithTimeZone    time.Time       `sql:"t_timestamp_with_time_zone"`
	Date                     time.Time       `sql:"t_date"`
	TimeWithoutTimeZone      time.Time       `sql:"t_time_without_time_zone"`
	TimeWithTimeZone         time.Time       `sql:"t_time_with_time_zone"`
	Interval                 time.Duration   `sql:"t_interval"`
	Boolean                  bool            `sql:"t_boolean"`
}

type Number struct {
	Id      int             `sql:"id"`
	Int     int             `sql:"t_int"`
	Int8    int8            `sql:"t_int_8"`
	Int16   int16           `sql:"t_int_16"`
	Int32   int32           `sql:"t_int_32"`
	Int64   int64           `sql:"t_int_64"`
	Uint    uint            `sql:"t_uint"`
	Uint8   uint8           `sql:"t_uint_8"`
	Uint16  uint16          `sql:"t_uint_16"`
	Uint32  uint32          `sql:"t_uint_32"`
	Uint64  uint64          `sql:"t_uint_64"`
	Float32 float32         `sql:"t_float_32"`
	Float64 float64         `sql:"t_float_64"`
	Decimal decimal.Decimal `sql:"t_decimal"`
}

type NumberPointer struct {
	Id      *int             `sql:"id"`
	Int     *int             `sql:"t_int"`
	Int8    *int8            `sql:"t_int_8"`
	Int16   *int16           `sql:"t_int_16"`
	Int32   *int32           `sql:"t_int_32"`
	Int64   *int64           `sql:"t_int_64"`
	Uint    *uint            `sql:"t_uint"`
	Uint8   *uint8           `sql:"t_uint_8"`
	Uint16  *uint16          `sql:"t_uint_16"`
	Uint32  *uint32          `sql:"t_uint_32"`
	Uint64  *uint64          `sql:"t_uint_64"`
	Float32 *float32         `sql:"t_float_32"`
	Float64 *float64         `sql:"t_float_64"`
	Decimal *decimal.Decimal `sql:"t_decimal"`
}

type Numbers struct {
	Id      int               `sql:"id"`
	Int     []int             `sql:"t_int"`
	Int8    []int8            `sql:"t_int_8"`
	Int16   []int16           `sql:"t_int_16"`
	Int32   []int32           `sql:"t_int_32"`
	Int64   []int64           `sql:"t_int_64"`
	Uint    []uint            `sql:"t_uint"`
	Uint8   []uint8           `sql:"t_uint_8"`
	Uint16  []uint16          `sql:"t_uint_16"`
	Uint32  []uint32          `sql:"t_uint_32"`
	Uint64  []uint64          `sql:"t_uint_64"`
	Float32 []float32         `sql:"t_float_32"`
	Float64 []float64         `sql:"t_float_64"`
	Decimal []decimal.Decimal `sql:"t_decimal"`
}

type NumbersPointer struct {
	Id      *int               `sql:"id"`
	Int     []*int             `sql:"t_int"`
	Int8    []*int8            `sql:"t_int_8"`
	Int16   []*int16           `sql:"t_int_16"`
	Int32   []*int32           `sql:"t_int_32"`
	Int64   []*int64           `sql:"t_int_64"`
	Uint    []*uint            `sql:"t_uint"`
	Uint8   []*uint8           `sql:"t_uint_8"`
	Uint16  []*uint16          `sql:"t_uint_16"`
	Uint32  []*uint32          `sql:"t_uint_32"`
	Uint64  []*uint64          `sql:"t_uint_64"`
	Float32 []*float32         `sql:"t_float_32"`
	Float64 []*float64         `sql:"t_float_64"`
	Decimal []*decimal.Decimal `sql:"t_decimal"`
}

type String struct {
	Id     int    `sql:"id"`
	String string `sql:"string"`
}

type StringPointer struct {
	Id     *int    `sql:"id"`
	String *string `sql:"string"`
}

type Strings struct {
	Id     int      `sql:"id"`
	String []string `sql:"string"`
}

type StringsPointer struct {
	Id     *int      `sql:"id"`
	String []*string `sql:"string"`
}

type Time struct {
	Id   int       `sql:"id"`
	Time time.Time `sql:"time"`
}

type TimePointer struct {
	Id   *int       `sql:"id"`
	Time *time.Time `sql:"time"`
}

type Times struct {
	Id   int         `sql:"id"`
	Time []time.Time `sql:"time"`
}

type TimesPointer struct {
	Id   *int         `sql:"id"`
	Time []*time.Time `sql:"time"`
}

type Duration struct {
	Id       int           `sql:"id"`
	Duration time.Duration `sql:"duration"`
}

type DurationPointer struct {
	Id       *int           `sql:"id"`
	Duration *time.Duration `sql:"duration"`
}

type Durations struct {
	Id       int             `sql:"id"`
	Duration []time.Duration `sql:"duration"`
}

type DurationsPointer struct {
	Id       int              `sql:"id"`
	Duration []*time.Duration `sql:"duration"`
}

type Byte struct {
	Id   int
	Byte byte
}

type Boolean struct {
	Boolean bool
}

type Booleans struct {
	Boolean []bool
}

type Bytes struct {
	Bytes []byte
}

type PostgresMapper struct {
	Test_bigint func() (Number, error)
	Test_int8   func() (Number, error)
	Test_bit    func() (Number, error)
	Test_uuid   func() (Number, error)
}
