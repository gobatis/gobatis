package tests

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestXMLIf(t *testing.T) {
	var products []*Product
	q := `
		select * from products 
		<where>
		    <if test="price > 0"> price > #{price} </if>
		    <if test="isAvailable">and is_available is true</if>
		</where>		                            		                            
	`
	err := db.Debug().Query(q, batis.Param("price", 1), batis.Param("isAvailable", true)).Scan(&products).Error
	require.NoError(t, err)

	spew.Json(products)

}
