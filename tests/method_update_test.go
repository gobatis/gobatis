package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	m := getProductsMap()
	m[Smartphone].Price = decimal.NewFromFloat(900)
	m[Smartphone].StockQuantity = 30

	affected, err := db.Debug().Affect(1).Update("products",
		map[string]any{
			"price":          m[Smartphone].Price,
			"stock_quantity": m[Smartphone].StockQuantity,
		},
		batis.Where("product_name = #{name}", batis.Param("name", Smartphone)),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	var product *Product
	err = db.Query(`select * from products where product_name = #{name}`,
		batis.Param("name", Smartphone)).Scan(&product).Error
	require.NoError(t, err)

	m[Smartphone].Id = product.Id
	m[Smartphone].AddedDateTime = product.AddedDateTime

	compareProduct(t, m[Smartphone], product)
}
