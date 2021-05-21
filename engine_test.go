package gobatis

import (
	"github.com/AlekSi/pointer"
	"github.com/gobatis/gobatis/bundle"
	"github.com/gobatis/gobatis/test/entity"
	"github.com/gobatis/gobatis/test/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

var (
	pwd string
)

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func rv(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func TestEngine(t *testing.T) {
	
	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	err := engine.Init()
	require.NoError(t, err)
	
	err = engine.master.Ping()
	require.NoError(t, err)
	
	defer func() {
		err = engine.master.Close()
		require.NoError(t, err)
	}()
	
	_testMapper := new(mapper.TestMapper)
	err = engine.BindMapper(_testMapper)
	require.NoError(t, err)
	
	testSelectInsert(t, _testMapper)
	testSelectInsertPointer(t, _testMapper)
	testSelectInsertForeachSlice(t, _testMapper)
}

func testSelectInsert(t *testing.T, _testMapper *mapper.TestMapper) {
	id, err := _testMapper.SelectInsert(entity.TestEntity{
		Int8:                     1,
		BigInt:                   2,
		Int:                      3,
		Decimal:                  decimal.NewFromFloat(3.14),
		Numeric:                  decimal.NewFromFloat(3.14156),
		Real:                     4,
		DoublePrecision:          5.1,
		SmallSerial:              6,
		Serial:                   7,
		BigSerial:                8,
		Money:                    "1.1",
		Char:                     "hello",
		Text:                     "world",
		TimestampWithoutTimeZone: time.Now(),
		TimestampWithTimeZone:    time.Now(),
		Date:                     time.Now(),
		TimeWithoutTimeZone:      time.Now(),
		TimeWithTimeZone:         time.Now(),
		Interval:                 100 * time.Second,
		Boolean:                  true,
	})
	
	require.NoError(t, err)
	require.Greater(t, id, 0, "returning id should greater 0")
}

func testSelectInsertPointer(t *testing.T, _testMapper *mapper.TestMapper) {
	dec := decimal.NewFromFloat(3.14)
	now := time.Now()
	interval := 100 * time.Second
	
	id, err := _testMapper.SelectInsertPointer(&entity.TestEntityPointer{
		Int8:                     pointer.ToInt8(1),
		BigInt:                   pointer.ToInt64(2),
		Int:                      pointer.ToInt(3),
		Decimal:                  &dec,
		Numeric:                  &dec,
		Real:                     pointer.ToFloat64(4),
		DoublePrecision:          pointer.ToFloat64(5.1),
		SmallSerial:              pointer.ToInt(6),
		Serial:                   pointer.ToInt(7),
		BigSerial:                pointer.ToInt(8),
		Money:                    pointer.ToString("1.1"),
		Char:                     pointer.ToString("hello"),
		Text:                     pointer.ToString("world"),
		TimestampWithoutTimeZone: &now,
		TimestampWithTimeZone:    &now,
		Date:                     &now,
		TimeWithoutTimeZone:      &now,
		TimeWithTimeZone:         &now,
		Interval:                 &interval,
		Boolean:                  pointer.ToBool(true),
	})
	
	require.NoError(t, err)
	require.Greater(t, id, int32(0), "returning id should greater 0")
}

func testSelectInsertForeachSlice(t *testing.T, _testMapper *mapper.TestMapper) {
	id, err := _testMapper.SelectInsertForeachSlice(entity.TestEntity{
		Int8: 1,
	}, []string{"tom", "alice"})
	require.NoError(t, err)
	require.Greater(t, id, 0, "returning id should greater 0")
}
