package cast

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestAddInt64(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
		iserr    bool
	}{
		{a: 0, b: 0, iserr: false},
		{a: 1, b: 1, expected: 2, iserr: false},
		{a: 1, b: -1, expected: 0, iserr: false},
		{a: -1, b: -1, expected: -2, iserr: false},
		{a: math.MaxInt64, b: 1, iserr: true},
		{a: 1, b: math.MaxInt64, iserr: true},
	}
	for _, v := range tests {
		r, err := AddInt64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestSubInt64(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: false},
		{a: -1, b: math.MaxInt64, expected: math.MinInt64, iserr: false},
		{a: -2, b: math.MaxInt64, iserr: true},
		{a: 1, b: math.MinInt64, iserr: true},
	}
	for _, v := range tests {
		r, err := SubInt64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestMulInt64(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: false},
		{a: 2, b: 3, expected: 6, iserr: false},
		{a: 2, b: -3, expected: -6, iserr: false},
		{a: -2, b: 3, expected: -6, iserr: false},
		{a: -1, b: math.MaxInt64, expected: -math.MaxInt64, iserr: false},
		{a: 1, b: math.MinInt64, expected: math.MinInt64, iserr: false},
		{a: -2, b: math.MaxInt64, iserr: true},
		{a: 2, b: math.MaxInt64, iserr: true},
	}
	for _, v := range tests {
		r, err := MulInt64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestDivInt64(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: true},
		{a: 4, b: 0, expected: 0, iserr: true},
		{a: 4, b: 2, expected: 2, iserr: false},
		{a: -4, b: 2, expected: -2, iserr: false},
		{a: 4, b: -2, expected: -2, iserr: false},
	}
	for _, v := range tests {
		r, err := DivInt64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestAddUint64(t *testing.T) {
	tests := []struct {
		a        uint64
		b        uint64
		expected uint64
		iserr    bool
	}{
		{a: 0, b: 0, iserr: false},
		{a: 1, b: 1, expected: 2, iserr: false},
		{a: math.MaxUint64, b: 1, iserr: true},
	}
	for _, v := range tests {
		r, err := AddUint64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestSubUint64(t *testing.T) {
	tests := []struct {
		a        uint64
		b        uint64
		expected uint64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: false},
		{a: 2, b: 1, expected: 1, iserr: false},
	}
	for _, v := range tests {
		r, err := SubUint64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestMulUint64(t *testing.T) {
	tests := []struct {
		a        uint64
		b        uint64
		expected uint64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: false},
		{a: 2, b: 3, expected: 6, iserr: false},
		{a: 2, b: math.MaxUint64, iserr: true},
	}
	for _, v := range tests {
		r, err := MulUint64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
	}
}

func TestDivUint64(t *testing.T) {
	tests := []struct {
		a        uint64
		b        uint64
		expected uint64
		iserr    bool
	}{
		{a: 0, b: 0, expected: 0, iserr: true},
		{a: 4, b: 0, expected: 0, iserr: true},
		{a: 4, b: 2, expected: 2, iserr: false},
	}
	for _, v := range tests {
		r, err := DivUint64E(v.a, v.b)
		if v.iserr {
			require.Error(t, err)
			continue
		}
		require.Equal(t, v.expected, r)
		require.NoError(t, err)
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
