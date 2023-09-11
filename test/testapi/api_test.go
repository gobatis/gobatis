package testapi

import (
	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/test/testdb"
	"github.com/stretchr/testify/require"
	"testing"
)

var db *batis.DB

func init() {
	var err error
	db, err = batis.Open(testdb.Open(""))
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}()
}

func TestAPI(t *testing.T) {
	err := db.Exec(`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 1);`).Error
	require.NoError(t, err)
}
