package batis

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type entity struct {
	Id   *int64
	Name string
	Age  int
}

func namer(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

func TestInset(t *testing.T) {
	item := &entity{
		Name: "tom",
		Age:  18,
	}
	i := insert{
		table:      "users",
		data:       item,
		onConflict: OnConflict("name,age", "do nothing"),
	}
	sql, err := i.SQL(namer)
	require.NoError(t, err)
	t.Log(sql)
}
