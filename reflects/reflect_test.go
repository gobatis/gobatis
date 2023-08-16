package reflects

import (
	"github.com/stretchr/testify/require"
	"testing"
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
	
	rs, err := ReflectRows(item)
	require.NoError(t, err)
	
	for _, v := range rs {
		t.Log("-")
		for _, vv := range v {
			t.Log(vv.column, vv.value)
		}
	}
	
	ReflectRows([]*entity{item})
	for _, v := range rs {
		t.Log("-")
		for _, vv := range v {
			t.Log(vv.column, vv.value)
		}
	}
}
