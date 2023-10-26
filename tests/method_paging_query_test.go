package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestPagingQuery(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	m := getProductsMap()
	c := map[int]string{
		0: Chair,
		1: BluetoothHeadphones,
		2: Smartwatch,
		3: Smartphone,
	}

	for i := 0; i <= 4; i++ {
		var products []*Product
		var count int64

		q := batis.PagingQuery{
			Select: "*",
			Count:  "1",
			From:   "products",
			Where:  "price <= #{price}",
			Order:  "price asc",
			Page:   int64(i),
			Limit:  1,
			Params: map[string]any{
				"price": decimal.NewFromInt(900),
			},
			Scan: func(s batis.PagingScanner) error {
				return s.Scan(&count, &products)
			},
		}

		err = db.Debug().PagingQuery(q).Error
		require.NoError(t, err)
		require.Equal(t, int64(4), count)

		if i < 4 {
			require.Equal(t, 1, len(products))
			for _, v := range products {
				vv := m[v.ProductName]
				vv.Id = v.Id
				vv.AddedDateTime = v.AddedDateTime
				require.Equal(t, c[i], v.ProductName)
				compareProduct(t, v, vv)
			}
		} else {
			require.Equal(t, 0, len(products))
		}
	}
}
