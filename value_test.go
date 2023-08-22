package batis

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestIndirect(t *testing.T) {
	
	var a *int
	rv := reflect.ValueOf(&a)
	
	t.Log(rv.Kind())
	
	rv = indirect(rv, false)
	
	t.Log(rv.Kind())
	
	rv.Set(reflect.ValueOf(3))
	
	t.Log(*a)
}

func TestSetValueBasic(t *testing.T) {
	
	var a *int
	rv := reflect.ValueOf(&a)
	err := setValue(rv, 1)
	require.NoError(t, err)
	t.Log(*a)
	
	var b int8
	rv = reflect.ValueOf(&b)
	err = setValue(rv, 9)
	require.NoError(t, err)
	t.Log(b)
	
	var c float32
	rv = reflect.ValueOf(&c)
	err = setValue(rv, 3)
	require.NoError(t, err)
	t.Log(c)
	
	var d string
	rv = reflect.ValueOf(&d)
	err = setValue(rv, "hello world")
	require.NoError(t, err)
	t.Log(d)
}
