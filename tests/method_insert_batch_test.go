package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Testing batch insertion of data, including checking the number of affected rows,
// verifying the inserted data, returning the last insert ID, and scanning all auto-incremented IDs.
func TestInsertBatch(t *testing.T) {

	products := getProductsList()

	// TODO affect should move to default mo to
	err := db.Debug().Affect(5).InsertBatch("products", 2, products).Error
	require.Error(t, err)

	affected, err := db.Debug().Affect(6).InsertBatch("products", 2, products).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(6), affected)

	var result []*Product
	err = db.Debug().Query(`select * from products`).Scan(&result).Error
	require.NoError(t, err)

	compareProducts(t, result, products)

	err = db.Exec("delete from products").Error
	require.NoError(t, err)

	//affected, err = db.Debug().Affect(5).
	//	Delete("products", batis.Where("stock_quantity >= #{ v }", batis.Param("v", 10))).RowsAffected()
	//require.NoError(t, err)
	//
	//products = []*Product{}
	//affected, err = db.Debug().Affect(5).InsertBatch("products", 2, extractMemProducts(Smartwatch),
	//	batis.Returning("*")).Scan(&products).RowsAffected()
	//require.NoError(t, err)
	//require.Equal(t, int64(5), affected)
	//
	//compareProducts(t, extractMemProducts(Smartwatch), products)
}
