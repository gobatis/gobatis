package testapi

import (
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

var db *batis.DB

func TestInsert(t *testing.T) {
	defer func() {
		require.NoError(t, db.Close())
	}()
	initDB(t)
	insertJack(t)
	queryJack(t)
	insertTom(t)
}

func insertJack(t *testing.T) {
	user := User{
		Id:   1,
		Name: "jack",
	}
	err := db.Insert("users", user).Error
	require.NoError(t, err)
}

func queryJack(t *testing.T) {
	var user *User
	err := db.Query(`select * from users where id = #{id}`, batis.Param("id", 1)).Scan(&user).Error
	require.NoError(t, err)
	require.Equal(t, "jack", user.Name)
}

func insertTom(t *testing.T) {
	err := db.Exec(`INSERT INTO users (id, name) VALUES (2, 'tom');`).Error
	require.NoError(t, err)
}

func initDB(t *testing.T) {
	var err error
	db, err = batis.Open(Open(""))
	if err != nil {
		panic(err)
	}

	err = db.Exec(`
		create table if not exists users
		(
			id     serial primary key,
			name   varchar
		);
    `).Error
	require.NoError(t, err)

}

type User struct {
	Id   int64
	Name string
}
