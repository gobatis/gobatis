package executor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAffecting_Check(t *testing.T) {

	r, err := newAffectingConstraint("10+")
	require.NoError(t, err)
	t.Log(r)
	err = r.Check(1)
	t.Log(err)
	require.True(t, errors.Is(err, RowsAffectedCheckErr))
	require.Error(t, err)
}
