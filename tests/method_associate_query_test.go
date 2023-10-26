package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestAssociateQuery(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	type ProductWrap struct {
		Name    string
		Product *Product
		Age     int64
	}

	wraps := []ProductWrap{
		{Name: "Laptop"},
		{Name: "TV"},
	}

	err = db.Debug().AssociateQuery(batis.AssociateQuery{
		SQL: "select * from products where product_name in #{ids}",
		Params: map[string]any{
			//"ids": batis.Extract(wraps, "$.Name"),
			"ids": []string{"Laptop", "TV"},
		},
		Scan: func(scanner batis.AssociateScanner) error {
			return scanner.Scan(&wraps, "product_name => $.Name", "$.Product")
		},
	}).Error

	require.NoError(t, err)

	spew.Json(wraps)
}
