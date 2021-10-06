package gobatis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSnackString(t *testing.T) {
	require.Equal(t, snake("HelloWorld"), "hello_world")
	require.Equal(t, snake("API"), "a_p_i")
	require.Equal(t, snake("Mapper"), "mapper")
}
