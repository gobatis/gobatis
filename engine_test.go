package gobatis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/gobatis/gobatis/test/entity"
	"github.com/gobatis/gobatis/test/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
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
	engine.BindSQL(NewBundle("test"))
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
	testSelectInsertForeachSlicePointer(t, _testMapper)
	testSelectInsertForeachMap(t, _testMapper)
	testSelectInsertForeachMapPointer(t, _testMapper)
	testSelectInsertForeachStruct(t, _testMapper)
	testSelectInsertForeachStructPointer(t, _testMapper)
	testSelectInsertContextTx(t, engine, _testMapper)
	testInsert(t, _testMapper)
	testSelectRow(t, _testMapper)
	testSelectRowPointer(t, _testMapper)
	testSelectRows(t, _testMapper)
	testSelectRowsPointer(t, _testMapper)
	testSelectStruct(t, _testMapper)
	testSelectStructs(t, _testMapper)
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

func testSelectInsertContextTx(t *testing.T, engine *Engine, _testMapper *mapper.TestMapper) {
	ctx := context.WithValue(context.Background(), "name", "gobatis")
	tx, err := engine.Master().Begin()
	require.NoError(t, err)
	id, err := _testMapper.SelectInsertContextTx(ctx, tx, entity.TestEntity{
		Char: "hello",
	})
	require.NoError(t, err)
	err = tx.Commit()
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testInsert(t *testing.T, _testMapper *mapper.TestMapper) {
	rows, err := _testMapper.Insert("Insert", "red", "yellow", "blue")
	require.NoError(t, err)
	if rows != 1 {
		require.Error(t, fmt.Errorf("rows expected 1"))
	}
}

func testSelectRow(t *testing.T, _testMapper *mapper.TestMapper) {
	tChar, tText, err := _testMapper.SelectRow(47)
	require.NoError(t, err)
	require.Equal(t, tChar, "hello")
	require.Equal(t, tText, "world")
}

func testSelectRowPointer(t *testing.T, _testMapper *mapper.TestMapper) {
	tChar, tText, err := _testMapper.SelectRowPointer(pointer.ToInt(47))
	require.NoError(t, err)
	require.Equal(t, *tChar, "hello")
	require.Equal(t, *tText, "world")
}

func testSelectRows(t *testing.T, _testMapper *mapper.TestMapper) {
	tChar, tText, err := _testMapper.SelectRows(363, 364)
	require.NoError(t, err)
	for _, v := range tChar {
		//require.Equal(t, v, "hello")
		fmt.Println(v)
	}
	for _, v := range tText {
		//require.Equal(t, v, "world")
		fmt.Printf("scanner: %+v\n", v)
	}
}

func testSelectRowsPointer(t *testing.T, _testMapper *mapper.TestMapper) {
	tChar, tText, err := _testMapper.SelectRowsPointer(pointer.ToInt(47), pointer.ToInt(50))
	require.NoError(t, err)
	for _, v := range tChar {
		require.Equal(t, *v, "hello")
	}
	for _, v := range tText {
		require.Equal(t, *v, "world")
	}
}

func testSelectStruct(t *testing.T, _testMapper *mapper.TestMapper) {
	item, err := _testMapper.SelectStruct(47)
	require.NoError(t, err)
	d, err := json.MarshalIndent(item, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(d))
	
	item2, err := _testMapper.SelectStructPointer(47)
	require.NoError(t, err)
	d, err = json.MarshalIndent(item2, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(d))
}

func testSelectStructs(t *testing.T, _testMapper *mapper.TestMapper) {
	item, err := _testMapper.SelectStructs(47)
	require.NoError(t, err)
	d, err := json.MarshalIndent(item, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(d))
	
	item2, err := _testMapper.SelectStructsPointer(47)
	require.NoError(t, err)
	d, err = json.MarshalIndent(item2, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(d))
}
