// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go
package cast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"html/template"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"time"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")

func castOverFlowError(from, to string) error {
	return fmt.Errorf("result overflow type from %s to %s", from, to)
}

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}) (tim time.Time, err error) {
	i = indirect(i)
	
	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDate(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// ToBoolE casts an interface to a bool type.
func ToBoolE(i interface{}) (bool, error) {
	i = indirect(i)
	
	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}

// ToInt64E casts an interface to an int64 type.
func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		a := big.Int{}
		a.SetUint64(s)
		b := big.NewInt(math.MaxInt64)
		if a.Cmp(b) > 1 {
			return 0, castOverFlowError("uint64", "int64")
		}
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
			
		}
		a, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		if !a.Equal(decimal.NewFromInt(v)) {
			return 0, castOverFlowError("string", "int64")
		}
		return v, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// ToInt32E casts an interface to an int32 type.
func ToInt32E(i interface{}) (int32, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case int:
		r := int32(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "int32")
		}
		return r, nil
	case int64:
		r := int32(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "int32")
		}
		return r, nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		r := int32(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "int32")
		}
		return r, nil
	case uint64:
		r := int32(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "int32")
		}
		return r, nil
	case uint32:
		r := int32(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "int32")
		}
		return r, nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		return int32(s), nil
	case float32:
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
		r := int32(v)
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		if !d.Equal(decimal.NewFromInt32(r)) {
			return 0, castOverFlowError("string", "32")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// ToInt16E casts an interface to an int16 type.
func ToInt16E(i interface{}) (int16, error) {
	i = indirect(i)
	switch s := i.(type) {
	case int:
		r := int16(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "int16")
		}
		return r, nil
	case int64:
		r := int16(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "int16")
		}
		return r, nil
	case int32:
		r := int16(s)
		if int32(r) != s {
			return 0, castOverFlowError("int32", "int16")
		}
		return r, nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		r := int16(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "int16")
		}
		return int16(s), nil
	case uint64:
		r := int16(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "int16")
		}
		return r, nil
	case uint32:
		r := int16(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "int16")
		}
		return r, nil
	case uint16:
		r := int16(s)
		if uint16(r) != s {
			return 0, castOverFlowError("uint64", "int16")
		}
		return r, nil
	case uint8:
		return int16(s), nil
	case float64:
		u := int64(s)
		r := int16(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "int16")
		}
		return r, nil
	case float32:
		u := int32(s)
		r := int16(s)
		if int32(r) != u {
			return 0, castOverFlowError("float32", "int16")
		}
		return r, nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
		r := int16(v)
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		if !d.Equal(decimal.NewFromInt32(int32(r))) {
			return 0, castOverFlowError("string", "int16")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// ToInt8E casts an interface to an int8 type.
func ToInt8E(i interface{}) (int8, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case int:
		r := int8(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "int8")
		}
		return r, nil
	case int64:
		r := int8(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "int8")
		}
		return r, nil
	case int32:
		r := int8(s)
		if int32(r) != s {
			return 0, castOverFlowError("int32", "int8")
		}
		return r, nil
	case int16:
		r := int8(s)
		if int16(r) != s {
			return 0, castOverFlowError("int16", "int8")
		}
		return r, nil
	case int8:
		return s, nil
	case uint:
		r := int8(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "int8")
		}
		return r, nil
	case uint64:
		r := int8(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "int8")
		}
		return r, nil
	case uint32:
		r := int8(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "int8")
		}
		return r, nil
	case uint16:
		r := int8(s)
		if uint16(r) != s {
			return 0, castOverFlowError("uint16", "int8")
		}
		return r, nil
	case uint8:
		r := int8(s)
		if uint8(r) != s {
			return 0, castOverFlowError("uint8", "int8")
		}
		return r, nil
	case float64:
		u := int64(s)
		r := int8(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "int8")
		}
		return r, nil
	case float32:
		u := int32(s)
		r := int8(s)
		if int32(r) != u {
			return 0, castOverFlowError("float32", "int8")
		}
		return r, nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
		r := int8(v)
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		if !d.Equal(decimal.NewFromInt32(int32(r))) {
			return 0, castOverFlowError("string", "int8")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// ToIntE casts an interface to an int type.
func ToIntE(i interface{}) (int, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		r := int(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "int")
		}
		return r, nil
	case int32:
		r := int(s)
		if int32(r) != s {
			return 0, castOverFlowError("int64", "int")
		}
		return r, nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		r := int(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "int")
		}
		return r, nil
	case uint64:
		r := int(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "int")
		}
		return r, nil
	case uint32:
		r := int(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "int")
		}
		return r, nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		u := int64(s)
		r := int(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "int")
		}
		return r, nil
	case float32:
		u := int32(s)
		r := int(s)
		if int32(r) != u {
			return 0, castOverFlowError("float32", "int8")
		}
		return r, nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
		r := int(v)
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		if !d.Equal(decimal.NewFromInt(int64(r))) {
			return 0, castOverFlowError("string", "int")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// ToUintE casts an interface to a uint type.
func ToUintE(i interface{}) (uint, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
		}
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		r := uint(v)
		q, _ := decimal.NewFromString(fmt.Sprintf("%d", r))
		if !d.Equal(q) {
			return 0, castOverFlowError("string", "uint")
		}
		return r, nil
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "uint")
		}
		return r, nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		r := uint(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "uint")
		}
		return r, nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int64(s)
		r := uint(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "uint")
		}
		return r, nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// ToUint64E casts an interface to a uint64 type.
func ToUint64E(i interface{}) (uint64, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
		}
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		q, _ := decimal.NewFromString(s)
		if !d.Equal(q) {
			return 0, castOverFlowError("string", "uint64")
		}
		return v, nil
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

// ToUint32E casts an interface to a uint32 type.
func ToUint32E(i interface{}) (uint32, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
		}
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		r := uint32(v)
		q, _ := decimal.NewFromString(fmt.Sprintf("%d", r))
		if !d.Equal(q) {
			return 0, castOverFlowError("string", "uint32")
		}
		return r, nil
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint32(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "uint32")
		}
		return r, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint32(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "uint32")
		}
		return r, nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case uint:
		r := uint32(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "uint32")
		}
		return r, nil
	case uint64:
		r := uint32(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint", "uint32")
		}
		return r, nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int64(s)
		r := uint32(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "uint32")
		}
		return r, nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// ToUint16E casts an interface to a uint16 type.
func ToUint16E(i interface{}) (uint16, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 16)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
		}
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		r := uint16(v)
		q, _ := decimal.NewFromString(fmt.Sprintf("%d", r))
		if !d.Equal(q) {
			return 0, castOverFlowError("string", "uint16")
		}
		return r, nil
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint16(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "uint16")
		}
		return r, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint16(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "uint16")
		}
		return r, nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint16(s)
		if int32(r) != s {
			return 0, castOverFlowError("int32", "uint16")
		}
		return r, nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case uint:
		r := uint16(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "uint16")
		}
		return r, nil
	case uint64:
		r := uint16(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "uint16")
		}
		return r, nil
	case uint32:
		r := uint16(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "uint16")
		}
		return r, nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int64(s)
		r := uint16(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "uint16")
		}
		return r, nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int32(s)
		r := uint16(s)
		if int32(r) != u {
			return 0, castOverFlowError("float32", "uint16")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// ToUint8E casts an interface to a uint type.
func ToUint8E(i interface{}) (uint8, error) {
	i = indirect(i)
	
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
		}
		d, err := decimal.NewFromString(s)
		if err != nil {
			return 0, err
		}
		r := uint8(v)
		q, _ := decimal.NewFromString(fmt.Sprintf("%d", r))
		if !d.Equal(q) {
			return 0, castOverFlowError("string", "uint8")
		}
		return r, nil
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint8(s)
		if int(r) != s {
			return 0, castOverFlowError("int", "uint8")
		}
		return r, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint8(s)
		if int64(r) != s {
			return 0, castOverFlowError("int64", "uint8")
		}
		return r, nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint8(s)
		if int32(r) != s {
			return 0, castOverFlowError("int32", "uint8")
		}
		return r, nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		r := uint8(s)
		if int16(r) != s {
			return 0, castOverFlowError("int16", "uint8")
		}
		return r, nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case uint:
		r := uint8(s)
		if uint(r) != s {
			return 0, castOverFlowError("uint", "uint8")
		}
		return r, nil
	case uint64:
		r := uint8(s)
		if uint64(r) != s {
			return 0, castOverFlowError("uint64", "uint8")
		}
		return r, nil
	case uint32:
		r := uint8(s)
		if uint32(r) != s {
			return 0, castOverFlowError("uint32", "uint8")
		}
		return r, nil
	case uint16:
		r := uint8(s)
		if uint16(r) != s {
			return 0, castOverFlowError("uint16", "uint8")
		}
		return r, nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int64(s)
		r := uint8(s)
		if int64(r) != u {
			return 0, castOverFlowError("float64", "uint8")
		}
		return r, nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		u := int32(s)
		r := uint8(s)
		if int32(r) != u {
			return 0, castOverFlowError("float32", "uint8")
		}
		return r, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringE casts an interface to a string type.
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)
	
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// StringToDate attempts to parse a string into a time.Time type using a
// predefined list of formats.  If no suitable format is found, an error is
// returned.
func StringToDate(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
