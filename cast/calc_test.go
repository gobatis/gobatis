package cast

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestAddInt64(t *testing.T) {
	tests := []struct {
		a     int64
		b     int64
		iserr bool
	}{
		{a: 0, b: 0, iserr: false},
		{a: math.MaxInt64, b: 1, iserr: true},
		{a: 1, b: math.MaxInt64, iserr: true},
	}
	for _, v := range tests {
		_, err := AddInt64(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		assert.NoError(t, err)
	}
}

func TestSubInt64(t *testing.T) {
	tests := []struct {
		a     int64
		b     int64
		iserr bool
	}{
		{a: 0, b: 0, iserr: false},
		{a: -1, b: math.MaxInt64, iserr: true},
		{a: 1, b: math.MinInt64, iserr: true},
	}
	for _, v := range tests {
		_, err := SubInt64(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		assert.NoError(t, err)
	}
}

func TestBigAddInt64(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected string
	}{
		{math.MaxInt64, 100, "9223372036854775907"},
		{math.MinInt64, -1, "-9223372036854775809"},
	}
	for _, v := range tests {
		r := BigAddInt64(v.a, v.b)
		assert.Equal(t, v.expected, r.String())
	}
}

func TestBigAddUint64(t *testing.T) {
	tests := []struct {
		a        uint64
		b        uint64
		expected string
	}{
		{math.MaxUint64, 100, "18446744073709551715"},
		{math.MaxUint64, 1, "18446744073709551616"},
	}
	for _, v := range tests {
		r := BigAddUint64(v.a, v.b)
		assert.Equal(t, v.expected, r.String())
	}
}
