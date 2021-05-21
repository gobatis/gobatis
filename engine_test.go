package gobatis

import (
	"github.com/gobatis/gobatis/bundle"
	"github.com/gobatis/gobatis/test/entity"
	"github.com/gobatis/gobatis/test/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"reflect"
	"testing"
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
	
	id, err := _testMapper.SelectInsert(entity.TestEntity{
		Int8:            1,
		BigInt:          2,
		Int:             3,
		Decimal:         decimal.NewFromFloat(3.14),
		Numeric:         decimal.NewFromFloat(3.14156),
		Real:            4,
		DoublePrecision: 5.1,
		SmallSerial:     6,
		Serial:          7,
		BigSerial:       8,
		Money:           "1.1",
		Char:            "hello",
		Text:            "world",
		//TimestampWithoutTimeZone: time.Now(),
		//TimestampWithTimeZone:    time.Now(),
		//Date:                     time.Now(),
		//TimeWithoutTimeZone:      time.Now(),
		//TimeWithTimeZone:         time.Now(),
		//Interval:                 100 * time.Second,
		Boolean: true,
	})
	
	require.NoError(t, err)
	require.Greaterf(t, id, 0, "id should greater 0")
	//
	//item, err := productMapper.GetProductById(11)
	//require.NoError(t, err)
	//require.Equal(t, "gobatis manual", item.Name)
	//require.Equal(t, "16.8", item.Price.String())
	//
	//items, err := productMapper.GetProductsById(11)
	//require.NoError(t, err)
	//require.Equal(t, 1, len(items))
	//require.Equal(t, item.Id, items[0].Id)
	//require.Equal(t, item.CreatedAt, items[0].CreatedAt)
	
	//item, err = productMapper.GetProductById(142)
	//require.NoError(t, err)
	//d, err := json.MarshalIndent(item, "", "\t")
	//require.NoError(t, err)
	//t.Log(string(d))
}

func TestExpression(t *testing.T) {
	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	err := engine.parseBundle()
	require.NoError(t, err)
}
