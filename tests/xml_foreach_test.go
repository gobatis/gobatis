package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestXMLForeach(t *testing.T) {
	var products []*Product
	err := db.Debug().Query(`
   select * from products where id in 
	<foreach item="item" index="index" collection="ids" open="(" separator="," close=")">
		#{item}
	</foreach>
`, batis.Param("ids", []int64{1, 2, 3})).Scan(&products).Error
	require.NoError(t, err)
}
