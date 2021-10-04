// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"html/template"
	"math"
	"testing"
	"time"
)

func TestToUintE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"18446744073709551616", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToUintE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToUint64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		// errors
		{"test", 0, true},
		{math.MaxInt64, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToUint32E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{math.MaxInt32, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToUint16E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{math.MaxUint8 + 1, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToUint8E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{"18446744073709551616", 0, true},
		{"-18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToIntE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int64
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{"18446744073709551616", 0, true},
		{"-18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToInt64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{math.MaxInt32 + 1, 0, true},
		{math.MinInt32 - 1, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToInt32E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{math.MaxInt16 + 1, 0, true},
		{math.MinInt16 - 1, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToInt16E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{math.MaxInt8 + 1, 0, true},
		{math.MinInt8 - 1, 0, true},
		{"18446744073709551616", 0, true},
		{testing.T{}, 0, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToInt8E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}
func TestToStringE(t *testing.T) {
	type Key struct {
		k string
	}
	key := &Key{"foo"}
	
	tests := []struct {
		input  interface{}
		expect string
		iserr  bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToStringE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

type foo struct {
	val string
}

func (x foo) String() string {
	return x.val
}

func TestStringerToString(t *testing.T) {
	var x foo
	x.val = "bar"
	v, err := ToStringE(x)
	require.NoError(t, err)
	assert.Equal(t, "bar", v)
}

type fu struct {
	val string
}

func (x fu) Error() string {
	return x.val
}

func TestErrorToString(t *testing.T) {
	var x fu
	x.val = "bar"
	v, err := ToStringE(x)
	require.NoError(t, err)
	assert.Equal(t, "bar", v)
}

func TestToTimeEE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
		iserr  bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), false},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		// errors
		{"2006", time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}
	
	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		
		v, err := ToTimeE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}
		
		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v.UTC(), errmsg)
	}
}
