package reflects

import (
	"testing"

	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/stretchr/testify/require"
)

type entity struct {
	Id   *int64
	Name string
	Age  int64
}

func TestReflect(t *testing.T) {

	item := &entity{
		Name: "tom",
		Age:  18,
	}

	rs, err := ReflectRows(item, postgres.Namer{}, "")
	require.NoError(t, err)

	for _, v := range rs {
		t.Log("-")
		for _, vv := range v {
			t.Log(vv.column, vv.value)
		}
	}

	ReflectRows([]*entity{item}, postgres.Namer{}, "")
	for _, v := range rs {
		t.Log("-")
		for _, vv := range v {
			t.Log(vv.column, vv.value)
		}
	}
}
