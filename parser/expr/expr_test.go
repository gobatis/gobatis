package expr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {

	type A struct {
		V int64
	}

	v, err := Parse(`d`, map[string]any{
		"a": A{
			V: 10,
		},
		"b": true,
		"c": false,
		"d": []string{"d1", "d2"},
	})
	require.NoError(t, err)
	t.Log(v.Interface())
}

func TestA(t *testing.T) {
	//fmt.Println()
}
