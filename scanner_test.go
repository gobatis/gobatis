package batis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBidingPathReg(t *testing.T) {
	r := bindingPathReg.FindStringSubmatch("a => $.A1.A2, b=>$.B")
	require.Equal(t, "a", r[2])
	require.Equal(t, "$.A1.A2", r[3])
	require.Equal(t, "b", r[6])
	require.Equal(t, "$.B", r[7])
}

func TestAssociateScannerParseBindingPath(t *testing.T) {
	{
		a := &associateScanner{}
		err := a.parseBindingPath("a => $.A1.A2, b=>$.B")
		require.NoError(t, err)
		require.Equal(t, "a", a.bindingPaths[0].column)
		require.Equal(t, "A1.A2", a.bindingPaths[0].path)
		require.Equal(t, "b", a.bindingPaths[1].column)
		require.Equal(t, "B", a.bindingPaths[1].path)
	}
	{
		a := &associateScanner{}
		err := a.parseBindingPath("product_name => $.Name")
		require.NoError(t, err)

		//require.Equal(t, "a", a.bindingPaths[0].column)
		//require.Equal(t, "A1.A2", a.bindingPaths[0].path)
		//require.Equal(t, "b", a.bindingPaths[1].column)
		//require.Equal(t, "B", a.bindingPaths[1].path)
	}
}
