package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	affected, err := db.Debug().Affect(1).Delete("products",
		batis.Where("product_name = #{name}", batis.Param("name", Smartphone))).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
}
