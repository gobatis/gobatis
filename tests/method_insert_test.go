package tests

import (
	"testing"
	"time"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// Test common insertion scenarios, including ordinary insertion,
// returning auto-incremented ID, handling conflict insertion,
// and handling conflict insertion while returning row fields
// 1. insert into ... returning ...
// 2. insert into ... on conflict ...
// 3. insert into ... on conflict ... returning ...
func TestInsert(t *testing.T) {
	// perform ordinary insertion operation and
	// return the auto-increment primary key
	m := getProductsMap()
	affected, err := db.Debug().Insert("products", m[Smartwatch], batis.Returning("id")).Scan(&m[Smartwatch].Id).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
	require.True(t, m[Smartwatch].Id != nil && *m[Smartwatch].Id > 0)

	// test insertion conflict
	m[Smartwatch].ManufactureDate = time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC)
	m[Smartwatch].AddedDateTime = time.Now()
	affected, err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			ManufactureDate: m[Smartwatch].ManufactureDate,
			AddedDateTime:   m[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set manufacture_date = excluded.manufacture_date`),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	// test insertion conflict update and
	// return the specified field
	var productName string
	m[Smartwatch].Price = decimal.NewFromFloat(300.00)
	m[Smartwatch].ManufactureDate = time.Now()
	m[Smartwatch].AddedDateTime = time.Now()
	err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			Price:           m[Smartwatch].Price,
			ManufactureDate: m[Smartwatch].ManufactureDate,
			AddedDateTime:   m[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set price = excluded.price`),
		batis.Returning("product_name")).Scan(&productName).Error
	require.NoError(t, err)
	require.Equal(t, "Smartwatch", productName)

	// test query operation and
	// compare the data after changes
	var product *Product
	err = db.Query(`select * from products where id = #{id}`, batis.Param("id", *m[Smartwatch].Id)).Scan(&product).Error
	require.NoError(t, err)
	require.True(t, product.Id != nil && *product.Id > 0)
	require.Equal(t, "Smartwatch", product.ProductName)
	require.Equal(t, "Advanced health and fitness tracking smartwatch", product.Description)
	require.Equal(t, "300", product.Price.String())
	require.Equal(t, float32(0.05), product.Weight)
	require.Equal(t, int64(5), product.StockQuantity)
	require.Equal(t, true, product.IsAvailable)
	require.Equal(t, "2023-04-12", product.ManufactureDate.Format("2006-01-02"))
	require.Equal(t, true, product.AddedDateTime.Unix() > 0)

	err = db.Debug().Delete("products", batis.Where("1=1")).Error
	require.NoError(t, err)
}
