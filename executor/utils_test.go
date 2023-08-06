package executor

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
