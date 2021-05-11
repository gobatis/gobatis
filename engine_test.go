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

func rv(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func TestEngine(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	err = engine.Init()
	require.NoError(t, err)

	err = engine.master.Ping()
	require.NoError(t, err)

	defer func() {
		err = engine.master.Close()
		require.NoError(t, err)
	}()

	productMapper := new(mapper.ProductMapper)
	err = engine.BindMapper(productMapper)
	require.NoError(t, err)

	affected, err := productMapper.CreateProduct(&entity.Product{
		Name:   "gobatis manual",
		Age:    1,
		Height: 17.8,
		Price:  decimal.NewFromFloat(16.8),
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	item, err := productMapper.GetProductById(24)
	require.NoError(t, err)
	require.Equal(t, "gobatis manual", item.Name)
	require.Equal(t, "16.8", item.Price.String())
}
