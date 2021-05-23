package gobatis

import (
	"fmt"
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
	
	//testSelectInsert(t, _testMapper)
	//testSelectInsertPointer(t, _testMapper)
	//testSelectInsertForeachSlice(t, _testMapper)
	//testSelectInsertForeachSlicePointer(t, _testMapper)
	//testSelectInsertForeachMap(t, _testMapper)
	//testSelectInsertForeachMapPointer(t, _testMapper)
	//testSelectInsertForeachStruct(t, _testMapper)
	testSelectInsertForeachStructPointer(t, _testMapper)
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
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
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
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachSlice(t *testing.T, _testMapper *mapper.TestMapper) {
	id, err := _testMapper.SelectInsertForeachSlice(entity.TestEntity{
		Int8: 1,
	}, []string{"tom", "alice"})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachSlicePointer(t *testing.T, _testMapper *mapper.TestMapper) {
	enums := [][]*string{
		{pointer.ToString("tom1"), pointer.ToString("alice1")},
		{pointer.ToString("tom2"), pointer.ToString("alice2")},
	}
	id, err := _testMapper.SelectInsertForeachSlicePointer(&entity.TestEntityPointer{
		Int8: pointer.ToInt8(1),
	}, &enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachMap(t *testing.T, _testMapper *mapper.TestMapper) {
	enums := map[string][]string{
		"first":  {"f1", "f2"},
		"second": {"fs", "s2"},
	}
	id, err := _testMapper.SelectInsertForeachMap(entity.TestEntity{
		Int8: 1,
	}, enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachMapPointer(t *testing.T, _testMapper *mapper.TestMapper) {
	enums := map[string][]*string{
		"first":  {pointer.ToString("f1"), pointer.ToString("f2")},
		"second": {pointer.ToString("fs"), pointer.ToString("s2")},
	}
	id, err := _testMapper.SelectInsertForeachMapPointer(&entity.TestEntityPointer{
		Int8: pointer.ToInt8(1),
	}, &enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachStruct(t *testing.T, _testMapper *mapper.TestMapper) {
	id, err := _testMapper.SelectInsertForeachStruct(entity.TestEntity{
		Int8: 1,
		Char: "Hello",
	})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachStructPointer(t *testing.T, _testMapper *mapper.TestMapper) {
	id, err := _testMapper.SelectInsertForeachStructPointer(&entity.TestEntityPointer{
		Char: pointer.ToString("Hello"),
	})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertContextTx(t *testing.T, _testMapper *mapper.TestMapper) {
	//id, err := _testMapper.SelectInsertContextTx(&entity.TestEntityPointer{
	//	Char: pointer.ToString("Hello"),
	//})
	//require.NoError(t, err)
	//if id <= 0 {
	//	require.Error(t, fmt.Errorf("returning id should greater 0"))
	//}
}
