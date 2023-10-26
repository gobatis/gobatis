package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestFetchQuery(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	m := getProductsMap()
	var products []*Product
	err = db.Debug().FetchQuery(batis.FetchQuery{
		SQL: "select * from products where price < #{price} order by price asc",
		Params: map[string]any{
			"price": 1000,
		},
		Batch: 2,
		Scan: func(scanner batis.Scanner) error {
			var items []*Product
			e := scanner.Scan(&items)
			if e != nil {
				return e
			}
			products = append(products, items...)
			return nil
		},
	}).Error
	require.NoError(t, err)
	expect := []*Product{
		m[Chair],
		m[BluetoothHeadphones],
		m[Smartwatch],
		m[Smartphone],
	}
	compareProducts(t, expect, products)
}
