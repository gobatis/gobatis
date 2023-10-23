package batis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {

	type Order struct {
		No int64
	}

	type User struct {
		Name   string
		Orders []Order
	}

	r, err := Extract(
		[]User{
			{Name: "a", Orders: []Order{{No: 1}}},
			{Name: "b", Orders: []Order{{No: 2}}},
		},
		"$.Orders")
	require.NoError(t, err)

	t.Log(r)
}
