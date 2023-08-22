package batis

import (
	"github.com/stretchr/testify/require"
	"testing"
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
