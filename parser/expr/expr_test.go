package expr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {

	type A struct {
		V int64
	}

	v, err := Parse(`d[1]`, map[string]any{
		"a": A{
			V: 10,
		},
		"b": "ok",
		"c": 1,
		"d": []string{"d1", "d2"},
	})
	require.NoError(t, err)
	t.Log(v)
}
