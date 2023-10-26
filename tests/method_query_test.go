package tests

import (
	"testing"

	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()
	var products []*Product
	err = db.Query(`select * from products`).Scan(&products).Error
	require.NoError(t, err)

	spew.Json(products)
}
