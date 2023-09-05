package batis

import (
	"testing"

	"github.com/gobatis/gobatis/driver/postgres"
	executor2 "github.com/gobatis/gobatis/executor"
	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestInset(t *testing.T) {
	item := &executor2.entity{
		Name: "tom",
		Age:  18,
	}
	i := insert{
		table: "public.users",
		data:  item,
		elems: []Element{OnConflict("name,age", "do nothing")},
	}
	sql, params, err := i.Raw(postgres.Namer{}, "")
	require.NoError(t, err)
	t.Log(sql)
	spew.Json(params)
}

func TestInsertBatch(t *testing.T) {
	items := []*executor2.entity{
		{Name: "tom", Age: 18},
		{Name: "jack", Age: 19},
	}
	i := insertBatch{
		table: "public.users",
		data:  items,
		batch: 2,
		elems: []Element{OnConflict("name,age", "do nothing")},
	}
	sql, params, err := i.Raw(postgres.Namer{}, "")
	require.NoError(t, err)
	t.Log(sql)
	spew.Json(params)
}

func TestUpdate(t *testing.T) {
	i := update{
		table: "public.users",
		data: map[string]any{
			"name": "123",
			"age":  21,
		},
		elems: []Element{Where("age = #{id}", Param("id", 10))},
	}
	sql, params, err := i.Raw(postgres.Namer{}, "")
	require.NoError(t, err)
	t.Log(sql)
	spew.Json(params)
}

func TestDelete(t *testing.T) {
	i := del{
		table: "public.users",
		elems: []Element{Where("age = #{id}", Param("id", 10))},
	}
	sql, params, err := i.Raw(postgres.Namer{}, "")
	require.NoError(t, err)
	t.Log(sql)
	spew.Json(params)
}

func TestPage(t *testing.T) {

}

func TestQuery(t *testing.T) {

}

func TestExec(t *testing.T) {

}

func TestBuild(t *testing.T) {

}
