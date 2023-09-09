package testmysql

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gozelle/spew"
)

func TestLastInsertId(t *testing.T) {
	db := &batis.DB{}
	id, err := db.Insert("users", nil).LastInsertId()
	if err != nil {
		return
	}
	spew.Json(id)
}
