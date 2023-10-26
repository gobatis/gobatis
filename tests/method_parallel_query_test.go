package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestParallelQuery(t *testing.T) {
	m := getProductsMap()
	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()
	var products []*Product
	var count int64
	err = db.Debug().ParallelQuery(
		batis.ParallelQuery{
			SQL: `select * from products where price <= #{ price }`,
			Params: map[string]any{
				"price": 300,
			},
			Scan: func(s batis.Scanner) error {
				return s.Scan(&products)
			},
		},
		batis.ParallelQuery{
			SQL: `select count(1) from products where price <= #{ price }`,
			Params: map[string]any{
				"price": 300,
			},
			Scan: func(s batis.Scanner) error {
				return s.Scan(&count)
			},
		},
	).Error
	require.NoError(t, err)
	require.Equal(t, int64(3), count)
	require.Equal(t, 3, len(products))

	for _, v := range products {
		vv := m[v.ProductName]
		vv.Id = v.Id
		vv.AddedDateTime = v.AddedDateTime
		compareProduct(t, v, vv)
	}
}
