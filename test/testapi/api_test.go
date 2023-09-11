package testapi

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/test/testdb"
	"github.com/stretchr/testify/require"
)

var db *batis.DB

func init() {
	var err error
	db, err = batis.Open(testdb.Open(""))
	if err != nil {
		panic(err)
	}

	err = db.Exec(`
		create table if not exists users
		(
			id     serial
				primary key,
			name   varchar
		);
    `).Error
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	fmt.Println(3)
	//	err = db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
}

type User struct {
	Id   int64
	Name string
}

func TestInsert(t *testing.T) {
	user := User{
		Id:   0,
		Name: "jack",
	}
	err := db.Insert("users", user).Error
	require.NoError(t, err)
}

func TestExec(t *testing.T) {
	err := db.Exec(`INSERT INTO users (id, name) VALUES (1, 'tom');`).Error
	require.NoError(t, err)
}

func TestQuery(t *testing.T) {

	db.Query(``)
}
