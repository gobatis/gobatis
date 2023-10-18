package testpostgres

import (
	"net"
	"testing"
	"time"

	batis "github.com/gobatis/gobatis"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var db *batis.DB

func prepareDatabase() {

}

func TestDB(t *testing.T) {

}

func TestInsert(t *testing.T) {

}

func TestInsertBatch(t *testing.T) {

}

func TestQuery(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestExec(t *testing.T) {

}

func TestPaging(t *testing.T) {

}

func TestParallelQuery(t *testing.T) {

}

func TestFetchQuery(t *testing.T) {

}

var data = []*DataType{
	{
		IntCol:       1,
		SmallintCol:  1,
		BigintCol:    10000000001,
		DecimalCol:   10.01,
		NumericCol:   10.01,
		RealCol:      1.1,
		DoubleCol:    1.11,
		SerialCol:    1,
		BigserialCol: 10000000001,

		CharCol:    "char1",
		VarcharCol: "varchar1",
		TextCol:    "text1",
		ByteaCol:   []byte("DEADBEEF"),

		DateCol:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TimeCol:      time.Date(0, 0, 0, 10, 0, 0, 0, time.UTC),
		TimestampCol: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),

		InetCol: net.ParseIP("192.168.1.1"),

		CidrCol:    net.IPNet{IP: net.ParseIP("192.168.1.0"), Mask: net.CIDRMask(24, 32)},
		MacaddrCol: net.HardwareAddr{0x08, 0x00, 0x2B, 0x01, 0x02, 0x03},

		BitCol:        []byte("1010101010"),
		BitVaryingCol: []byte("10101"),

		UuidCol: uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),

		JsonCol:  map[string]interface{}{"key1": "value1"},
		JsonbCol: map[string]interface{}{"key2": "value2"},

		IntArrayCol:  pq.Int64Array{1, 2, 3},
		TextArrayCol: pq.StringArray{"one", "two", "three"},
	},
}

type DataType struct {
	IntCol       int32
	SmallintCol  int16
	BigintCol    int64
	DecimalCol   float64
	NumericCol   float64
	RealCol      float32
	DoubleCol    float64
	SerialCol    int32
	BigserialCol int64

	CharCol    string
	VarcharCol string
	TextCol    string
	ByteaCol   []byte

	DateCol      time.Time
	TimeCol      time.Time
	TimestampCol time.Time
	IntervalCol  pq.NullTime // pq doesn't have a direct Interval type, so using NullTime as an example

	InetCol    net.IP
	CidrCol    net.IPNet
	MacaddrCol net.HardwareAddr

	BitCol        []byte
	BitVaryingCol []byte
	TsvectorCol   string // This is a simplification, real usage might differ
	TsqueryCol    string // This is a simplification, real usage might differ

	UuidCol uuid.UUID

	JsonCol  map[string]interface{}
	JsonbCol map[string]interface{}

	IntArrayCol  pq.Int64Array
	TextArrayCol pq.StringArray

	PointCol pq.NullTime // pq doesn't have a direct Point type, so using NullTime as an example
	LineCol  string      // This is a simplification, real usage might differ

}

const createDataTypesTable = `
CREATE TABLE IF NOT EXISTS data_types
(
    id              SERIAL PRIMARY KEY,
    -- Basic
    int_col         INT,
    smallint_col    SMALLINT,
    bigint_col      BIGINT,
    decimal_col     DECIMAL(10, 2),
    numeric_col     NUMERIC(10, 2),
    real_col        REAL,
    double_col      DOUBLE PRECISION,
    serial_col      SERIAL,
    bigserial_col   BIGSERIAL,

    -- Text and Binary"
    char_col        CHAR(10),
    varchar_col     VARCHAR(100),
    text_col        TEXT,
    bytea_col       BYTEA,

    -- Time and Date
    date_col        DATE,
    time_col        TIME,
    timestamp_col   TIMESTAMP,
    interval_col    INTERVAL,

    -- Network
    inet_col        INET,
    cidr_col        CIDR,
    macaddr_col     MACADDR,

    -- Bit
    bit_col         BIT(10),
    bit_varying_col BIT VARYING(10),

    -- Vector
    tsvector_col    TSVECTOR,
    tsquery_col     TSQUERY,

    -- UUID
    uuid_col        UUID,

    -- JSON
    json_col        JSON,
    jsonb_col       JSONB,

    -- Array
    int_array_col   INT[],
    text_array_col  TEXT[],

    -- Geometric
    point_col       POINT,
    line_col        LINE
);
`

const insertDataTypesData = `
INSERT INTO data_types (
    id, int_col, smallint_col, bigint_col, decimal_col, numeric_col, real_col, double_col,
    serial_col, bigserial_col, char_col, varchar_col, text_col, bytea_col, date_col,
    time_col, timestamp_col, interval_col, inet_col, cidr_col, macaddr_col, bit_col,
    bit_varying_col, tsvector_col, tsquery_col, uuid_col, json_col, jsonb_col,
    int_array_col, text_array_col, point_col, line_col
) VALUES
(1,1, 1, 10000000001, 10.01, 10.01, 1.1, 1.11, 1, 10000000001, 'char1', 'varchar1', 'text1', E'\\xDEADBEEF', '2023-01-01', '10:00:00', '2023-01-01 10:00:00', '1 year', '192.168.1.1', '192.168.1.0/24', '08:00:2B:01:02:03', B'1010101010', B'10101', 'apple', 'apple', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '{"key1": "value1"}', '{"key2": "value2"}', ARRAY[1,2,3], ARRAY['one', 'two', 'three'], '(1,1)', '{1,-1,0}'),
(2,2, 2, 10000000002, 20.02, 20.02, 2.2, 2.22, 2, 10000000002, 'char2', 'varchar2', 'text2', E'\\xDEADBEAF', '2023-02-02', '11:00:00', '2023-02-02 11:00:00', '2 years', '192.168.2.2', '192.168.2.0/24', '08:00:2B:01:02:04', B'0101010101', B'01010', 'banana', 'banana', 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '{"key2": "value2"}', '{"key3": "value3"}', ARRAY[4,5,6], ARRAY['four', 'five', 'six'], '(2,2)', '{2,-2,0}')
on conflict do nothing;
`
