package executor

import (
	"reflect"
	"testing"

	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/stretchr/testify/require"
)

func TestReplaceIsolatedLessThan(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is a <test> string with a single < and another <tag> and one more <.", "This is a <test> string with a single &lt; and another <tag> and one more &lt;."},
		{"<Hello<world", "&lt;Hello&lt;world"},
		{"Hello<world>", "Hello<world>"},
		{"<<<<<", "&lt;&lt;&lt;&lt;&lt;"},
		{"<", "&lt;"},
		{"Hello<", "Hello&lt;"},
		{">Hello", ">Hello"},
		{"<10", "&lt;10"},
	}

	for _, test := range tests {
		result := replaceIsolatedLessThanWithEntity(test.input)
		require.Equal(t, test.expected, result)
	}
}

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

func TestExtract(t *testing.T) {
	type R struct {
		A int64
	}

	items := []R{
		{A: 1},
		{A: 2},
		{A: 3},
	}

	r := Extract[R, int64](items, func(item R) int64 {
		return item.A
	})
	t.Log(r)
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
