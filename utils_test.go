package batis

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"HelloWorldExample", "hello_world_example"},
		{"anotherExampleWithID", "another_example_with_id"},
		{"simple", "simple"},
		{"with_SomeID", "with_some_id"},
		{"already_snake_case", "already_snake_case"},
		{"WithNumbers123And456", "with_numbers123_and456"},
	}

	for _, tt := range tests {
		result := toSnakeCase(tt.input)
		require.Equal(t, tt.expected, result)
	}
}

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

func TestIsStructSlice(t *testing.T) {

	type A struct {
		Name string
	}

	a := []A{{Name: "a"}}
	v := reflect.ValueOf(a)
	ok := isStructSlice(v.Type())
	require.True(t, ok)

	b := []*A{{Name: "b"}}
	v = reflect.ValueOf(b)
	ok = isStructSlice(v.Type())
	require.True(t, ok)
}

func TestSplitStructSlice(t *testing.T) {
	type Record struct {
		Num int
	}
	data := []Record{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
		{Num: 5},
		{Num: 6},
		{Num: 7},
		{Num: 8},
		{Num: 9},
	}
	chunks, err := SplitStructSlice(data, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	spew.Json(chunks)
}
