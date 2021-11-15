package postgresql

import (
	"database/sql"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"time"
)

type Entity struct {
	Id                       int64           `sql:"id"`
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

type EntityPointer struct {
	Id                       *int64
	Int8                     *int8            `sql:"t_int8"`
	BigInt                   *int64           `sql:"t_bigint"`
	Int                      *int             `sql:"t_int"`
	Decimal                  *decimal.Decimal `sql:"t_decimal"`
	Numeric                  *decimal.Decimal `sql:"t_numeric"`
	Real                     *float64         `sql:"t_real"`
	DoublePrecision          *float64         `sql:"t_double_precision"`
	SmallSerial              *int             `sql:"t_small_serial"`
	Serial                   *int             `sql:"t_serial"`
	BigSerial                *int             `sql:"t_big_serial"`
	Money                    *string          `sql:"t_money"`
	Char                     *string          `sql:"t_char"`
	Text                     *string          `sql:"t_text"`
	TimestampWithoutTimeZone *time.Time       `sql:"t_timestamp_without_time_zone"`
	TimestampWithTimeZone    *time.Time       `sql:"t_timestamp_with_time_zone"`
	Date                     *time.Time       `sql:"t_date"`
	TimeWithoutTimeZone      *time.Time       `sql:"t_time_without_time_zone"`
	TimeWithTimeZone         *time.Time       `sql:"t_time_with_time_zone"`
	Interval                 *time.Duration   `sql:"t_interval"`
	Boolean                  *bool            `sql:"t_boolean"`
}

type User struct {
	Id   int64            `sql:"id"`
	Name string           `sql:"name"`
	Age  int              `sql:"age"`
	From string           `sql:"from"`
	Tags pgtype.TextArray `sql:"tags"`
}

type Member struct {
	Id        int64      `sql:"id"`
	Username  string     `sql:"username"`
	Email     string     `sql:"email"`
	Mobile    string     `sql:"mobile"`
	Password  string     `sql:"password"`
	Status    int        `sql:"status"`
	CreatedAt time.Time  `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoType struct {
	Id                     int64           `sql:"id"`
	TBigint                int64           `sql:"t_bigint" accept:"int*,uint*" reject:"float*,string"`
	TInt8                  int8            `sql:"t_int8"`
	TBigserial             int64           `sql:"t_bigserial"`
	TSerial8               int64           `sql:"t_serial8"`
	TBit                   byte            `sql:"t_bit"`
	TBitVarying            byte            `sql:"t_bit_varying"`
	TBoolean               bool            `sql:"t_boolean"`
	TBool                  bool            `sql:"t_bool"`
	TBox                   string          `sql:"t_box"`
	TBytea                 byte            `sql:"t_bytea"`
	TCharacter             string          `sql:"t_character"`
	TChar                  string          `sql:"t_char"`
	TCharacterVarying      string          `sql:"t_character_varying"`
	TVarchar               string          `sql:"t_varchar"`
	TCidr                  string          `sql:"t_cidr"`
	TCircle                string          `sql:"t_circle"`
	TDate                  time.Time       `sql:"t_date"`
	TDoublePrecision       float64         `sql:"t_double_precision"`
	TFloat8                float32         `sql:"t_float8"`
	TInet                  string          `sql:"t_inet"`
	TInteger               int             `sql:"t_integer"`
	TInt                   int             `sql:"t_int"`
	TInt4                  int8            `sql:"t_int4"`
	TInterval              time.Duration   `sql:"t_interval"`
	TJson                  string          `sql:"t_json"`
	TJsonb                 string          `sql:"t_jsonb"`
	TLine                  int             `sql:"t_line"`
	TLseg                  int             `sql:"t_lseg"`
	TMacaddr               string          `sql:"t_macaddr"`
	TMacaddr8              string          `sql:"t_macaddr8"`
	TMoney                 decimal.Decimal `sql:"t_money"`
	TNumeric               decimal.Decimal `sql:"t_numeric"`
	TDecimal               decimal.Decimal `sql:"t_decimal"`
	TPath                  int             `sql:"t_path"`
	TPgLsn                 int             `sql:"t_pg_lsn"`
	TPgSnapshot            int             `sql:"t_pg_snapshot"`
	TPoint                 int             `sql:"t_point"`
	TPolygon               int             `sql:"t_polygon"`
	TReal                  int             `sql:"t_real"`
	TFloat4                int             `sql:"t_float4"`
	TSmallint              int             `sql:"t_smallint"`
	TInt2                  int             `sql:"t_int2"`
	TSmallserial           int             `sql:"t_smallserial"`
	TSerial2               int             `sql:"t_serial2"`
	TSerial                int             `sql:"t_serial"`
	TSerial4               int             `sql:"t_serial4"`
	TText                  string          `sql:"t_text"`
	TTime                  time.Time       `sql:"t_time"`
	TTimeWithTimezone      time.Time       `sql:"t_time_with_timezone"`
	TTimez                 time.Time       `sql:"t_timez"`
	TTimestamp             time.Time       `sql:"t_timestamp"`
	TTimestampWithTimezone time.Time       `sql:"t_timestamp_with_timezone"`
	TTimestampz            time.Time       `sql:"t_timestampz"`
	TTsquery               int8            `sql:"t_tsquery"`
	TTsvector              int8            `sql:"t_tsvector"`
	TTxidSnapshot          int8            `sql:"t_txid_snapshot"`
	TUuid                  string          `sql:"t_uuid"`
	TXml                   string          `sql:"t_xml"`
}

type PgxType struct {
	Id                     int64              `sql:"id"`
	TBigint                int64              `sql:"t_bigint"`
	TInt8                  pgtype.Int8        `sql:"t_int8"`
	TBigserial             int64              `sql:"t_bigserial"`
	TSerial8               int64              `sql:"t_serial8"`
	TBit                   pgtype.Bit         `sql:"t_bit"`
	TBitVarying            pgtype.Varbit      `sql:"t_bit_varying"`
	TBoolean               bool               `sql:"t_boolean"`
	TBool                  pgtype.Bool        `sql:"t_bool"`
	TBox                   pgtype.Box         `sql:"t_box"`
	TBytea                 pgtype.Bytea       `sql:"t_bytea"`
	TCharacter             string             `sql:"t_character"`
	TChar                  string             `sql:"t_char"`
	TCharacterVarying      string             `sql:"t_character_varying"`
	TVarchar               string             `sql:"t_varchar"`
	TCidr                  pgtype.CIDR        `sql:"t_cidr"`
	TCircle                pgtype.Circle      `sql:"t_circle"`
	TDate                  pgtype.Date        `sql:"t_date"`
	TDoublePrecision       float64            `sql:"t_double_precision"`
	TFloat8                float32            `sql:"t_float8"`
	TInet                  pgtype.Inet        `sql:"t_inet"`
	TInteger               int                `sql:"t_integer"`
	TInt                   int                `sql:"t_int"`
	TInt4                  pgtype.Int4        `sql:"t_int4"`
	TInterval              pgtype.Interval    `sql:"t_interval"`
	TJson                  pgtype.JSON        `sql:"t_json"`
	TJsonb                 pgtype.JSONB       `sql:"t_jsonb"`
	TLine                  pgtype.Line        `sql:"t_line"`
	TLseg                  pgtype.Lseg        `sql:"t_lseg"`
	TMacaddr               pgtype.Macaddr     `sql:"t_macaddr"`
	TMacaddr8              pgtype.Macaddr     `sql:"t_macaddr8"`
	TMoney                 decimal.Decimal    `sql:"t_money"`
	TNumeric               decimal.Decimal    `sql:"t_numeric"`
	TDecimal               decimal.Decimal    `sql:"t_decimal"`
	TPath                  pgtype.Path        `sql:"t_path"`
	TPgLsn                 int                `sql:"t_pg_lsn"`
	TPgSnapshot            int                `sql:"t_pg_snapshot"`
	TPoint                 pgtype.Point       `sql:"t_point"`
	TPolygon               pgtype.Polygon     `sql:"t_polygon"`
	TReal                  int                `sql:"t_real"`
	TFloat4                pgtype.Float4      `sql:"t_float4"`
	TSmallint              int                `sql:"t_smallint"`
	TInt2                  int                `sql:"t_int2"`
	TSmallserial           int                `sql:"t_smallserial"`
	TSerial2               int                `sql:"t_serial2"`
	TSerial                int                `sql:"t_serial"`
	TSerial4               int                `sql:"t_serial4"`
	TText                  pgtype.Text        `sql:"t_text"`
	TTime                  pgtype.Time        `sql:"t_time"`
	TTimeWithTimezone      pgtype.Time        `sql:"t_time_with_timezone"`
	TTimez                 pgtype.Time        `sql:"t_timez"`
	TTimestamp             pgtype.Timestamp   `sql:"t_timestamp"`
	TTimestampWithTimezone pgtype.Timestamp   `sql:"t_timestamp_with_timezone"`
	TTimestampz            pgtype.Timestamptz `sql:"t_timestampz"`
	TTsquery               int8               `sql:"t_tsquery"`
	TTsvector              int8               `sql:"t_tsvector"`
	TTxidSnapshot          int8               `sql:"t_txid_snapshot"`
	TUuid                  pgtype.UUID        `sql:"t_uuid"`
	TXml                   string             `sql:"t_xml"`
}
