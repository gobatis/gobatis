package postgres

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamer(t *testing.T) {
	n := &Namer{}
	a := n.ReservedName(`"insert"`)
	require.Equal(t, a, `"insert"`)
	
	b := n.TableName("public.products")
	require.Equal(t, b, `"public"."products"`)
}
