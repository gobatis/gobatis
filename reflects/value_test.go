package reflects

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
	err := SetValue(&a, 1)
	require.NoError(t, err)
	t.Log(*a)
	
	var b int8
	err = SetValue(&b, 9)
	require.NoError(t, err)
	t.Log(b)
	
	var c float32
	err = SetValue(&c, 3)
	require.NoError(t, err)
	t.Log(c)
	
	var d string
	err = SetValue(&d, "hello world")
	require.NoError(t, err)
	t.Log(d)
}
