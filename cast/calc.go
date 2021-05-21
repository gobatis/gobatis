package cast

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

var differentOperandTypeErr = fmt.Errorf("operand types are different")

func resultOverFlowError(t string) error {
	return fmt.Errorf("result overflow type %s", t)
}

func BigAddInt64(a, b int64) *big.Int {
	r := big.NewInt(0)
	return r.Add(big.NewInt(a), big.NewInt(b))
}

func BigSubInt64(a, b int64) *big.Int {
	r := big.NewInt(0)
	return r.Sub(big.NewInt(a), big.NewInt(b))
}

func BigMulInt64(a, b int64) *big.Int {
	r := big.NewInt(0)
	return r.Mul(big.NewInt(a), big.NewInt(b))
}

func BigDivInt64(a, b int64) *big.Int {
	r := big.NewInt(0)
	return r.Div(big.NewInt(a), big.NewInt(b))
}

func BigAddUint64(a, b uint64) *big.Int {
	r := big.NewInt(0)
	ba := &big.Int{}
	ba.SetUint64(a)
	bb := &big.Int{}
	bb.SetUint64(b)
	return r.Add(ba, bb)
}

func BigSubUint64(a, b uint64) *big.Int {
	r := big.NewInt(0)
	ba := &big.Int{}
	ba.SetUint64(a)
	bb := &big.Int{}
	bb.SetUint64(b)
	return r.Sub(ba, bb)
}

func BigMulUint64(a, b uint64) *big.Int {
	r := big.NewInt(0)
	ba := &big.Int{}
	ba.SetUint64(a)
	bb := &big.Int{}
	bb.SetUint64(b)
	return r.Mul(ba, bb)
}

func BigDivUint64(a, b uint64) *big.Int {
	r := big.NewInt(0)
	ba := &big.Int{}
	ba.SetUint64(a)
	bb := &big.Int{}
	bb.SetUint64(b)
	return r.Div(ba, bb)
}

func AddFloat64(a, b float64) (r decimal.Decimal) {
	return decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b))
}

func SubFloat64(a, b float64) (r decimal.Decimal) {
	return decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b))
}

func MulFloat64(a, b float64) (r decimal.Decimal) {
	return decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b))
}

func DivFloat64(a, b float64) (r decimal.Decimal) {
	return decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b))
}

func AddFloat32(a, b float32) (r decimal.Decimal) {
	return decimal.NewFromFloat32(a).Add(decimal.NewFromFloat32(b))
}

func SubFloat32(a, b float32) (r decimal.Decimal) {
	return decimal.NewFromFloat32(a).Sub(decimal.NewFromFloat32(b))
}

func MulFloat32(a, b float32) (r decimal.Decimal) {
	return decimal.NewFromFloat32(a).Mul(decimal.NewFromFloat32(b))
}

func DivFloat32(a, b float32) (r decimal.Decimal) {
	return decimal.NewFromFloat32(a).Div(decimal.NewFromFloat32(b))
}

func AddInt64E(a, b int64) (r int64, err error) {
	r = a + b
	br := BigAddInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
		return
	}
	return
}

func SubInt64E(a, b int64) (r int64, err error) {
	r = a - b
	br := BigSubInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
		return
	}
	return
}

func MulInt64E(a, b int64) (r int64, err error) {
	r = a * b
	br := BigMulInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
		return
	}
	return
}

func DivInt64E(a, b int64) (r int64, err error) {
	if b == 0 {
		err = fmt.Errorf("divide 0")
		return
	}
	//br := BigDivInt64(a, b)
	//if big.NewInt(r).Cmp(br) != 0 {
	//	err = resultOverFlowError("int64")
	//	return
	//}
	return a / b, nil
}

func AddUint64E(a, b uint64) (r uint64, err error) {
	r = a + b
	br := BigAddUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
		return
	}
	return
}

func SubUint64E(a, b uint64) (r uint64, err error) {
	//r = a - b
	//br := BigSubUint64(a, b)
	//rr := &big.Int{}
	//rr.SetUint64(r)
	//if rr.Cmp(br) != 0 {
	//	err = resultOverFlowError("uint64")
	//	return
	//}
	return a - b, nil
}

func MulUint64E(a, b uint64) (r uint64, err error) {
	r = a * b
	br := BigMulUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
		return
	}
	return
}

func DivUint64E(a, b uint64) (r uint64, err error) {
	if b == 0 {
		err = fmt.Errorf("divide 0")
		return
	}
	//r = a / b
	//br := BigDivUint64(a, b)
	//rr := &big.Int{}
	//rr.SetUint64(r)
	//if rr.Cmp(br) != 0 {
	//	err = resultOverFlowError("uint64")
	//	return
	//}
	return a / b, nil
}

func AddIntE(a, b int) (r int, err error) {
	r1, err := AddInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int")
		return
	}
	return
}

func SubIntE(a, b int) (r int, err error) {
	r1, err := SubInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int")
		return
	}
	return
}

func MulIntE(a, b int) (r int, err error) {
	r1, err := MulInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int")
		return
	}
	return
}

func DivIntE(a, b int) (r int, err error) {
	r1, err := DivInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int")
		return
	}
	return
}

func AddInt8E(a, b int8) (r int8, err error) {
	r1, err := AddInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int8(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int8")
		return
	}
	return
}

func SubInt8E(a, b int8) (r int8, err error) {
	r1, err := SubInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int8(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int8")
		return
	}
	return
}

func MulInt8E(a, b int8) (r int8, err error) {
	r1, err := MulInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int8(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int8")
		return
	}
	return
}

func DivInt8E(a, b int8) (r int8, err error) {
	r1, err := DivInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int8(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int8")
		return
	}
	return
}

func AddInt16E(a, b int16) (r int16, err error) {
	r1, err := AddInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int16(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int16")
		return
	}
	return
}

func SubInt16E(a, b int16) (r int16, err error) {
	r1, err := SubInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int16(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int16")
		return
	}
	return
}

func MulInt16E(a, b int16) (r int16, err error) {
	r1, err := MulInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int16(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int16")
		return
	}
	return
}

func DivInt16E(a, b int16) (r int16, err error) {
	r1, err := DivInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int16(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int16")
		return
	}
	return
}

func AddInt32E(a, b int32) (r int32, err error) {
	r1, err := AddInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int32(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int32")
		return
	}
	return
}

func SubInt32E(a, b int32) (r int32, err error) {
	r1, err := SubInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int32(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int32")
		return
	}
	return
}

func MulInt32E(a, b int32) (r int32, err error) {
	r1, err := MulInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int32(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int32")
		return
	}
	return
}

func DivInt32E(a, b int32) (r int32, err error) {
	r1, err := DivInt64E(int64(a), int64(b))
	if err != nil {
		return
	}
	r = int32(r1)
	if int64(r) != r1 {
		err = resultOverFlowError("int32")
		return
	}
	return
}

func AddUintE(a, b uint) (r uint, err error) {
	r1, err := AddUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint")
		return
	}
	return
}

func SubUintE(a, b uint) (r uint, err error) {
	r1, err := SubUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint")
		return
	}
	return
}

func MulUintE(a, b uint) (r uint, err error) {
	r1, err := MulUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint")
		return
	}
	return
}

func DivUintE(a, b uint) (r uint, err error) {
	r1, err := DivUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint")
		return
	}
	return
}

func AddUint8E(a, b uint8) (r uint8, err error) {
	r1, err := AddUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint8(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint8")
		return
	}
	return
}

func SubUint8E(a, b uint8) (r uint8, err error) {
	r1, err := SubUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint8(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint8")
		return
	}
	return
}

func MulUint8E(a, b uint8) (r uint8, err error) {
	r1, err := MulUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint8(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint8")
		return
	}
	return
}

func DivUint8E(a, b uint8) (r uint8, err error) {
	r1, err := DivUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint8(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint8")
		return
	}
	return
}

func AddUint16E(a, b uint16) (r uint16, err error) {
	r1, err := AddUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint16(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint16")
		return
	}
	return
}

func SubUint16E(a, b uint16) (r uint16, err error) {
	r1, err := SubUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint16(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint16")
		return
	}
	return
}

func MulUint16E(a, b uint16) (r uint16, err error) {
	r1, err := MulUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint16(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint16")
		return
	}
	return
}

func DivUint16E(a, b uint16) (r uint16, err error) {
	r1, err := DivUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint16(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint16")
		return
	}
	return
}

func AddUint32E(a, b uint32) (r uint32, err error) {
	r1, err := AddUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint32(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint32")
		return
	}
	return
}

func SubUint32E(a, b uint32) (r uint32, err error) {
	r1, err := SubUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint32(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint32")
		return
	}
	return
}

func MulUint32E(a, b uint32) (r uint32, err error) {
	r1, err := DivUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint32(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint32")
		return
	}
	return
}

func DivUint32E(a, b uint32) (r uint32, err error) {
	r1, err := DivUint64E(uint64(a), uint64(b))
	if err != nil {
		return
	}
	r = uint32(r1)
	if uint64(r) != r1 {
		err = resultOverFlowError("uint32")
		return
	}
	return
}

func AddAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return AddIntE(o1.(int), o2.(int))
	case int8:
		return AddInt8E(o1.(int8), o2.(int8))
	case int16:
		return AddInt16E(o1.(int16), o2.(int16))
	case int32:
		return AddInt32E(o1.(int32), o2.(int32))
	case int64:
		return AddInt64E(o1.(int64), o2.(int64))
	case uint:
		return AddUintE(o1.(uint), o2.(uint))
	case uint8:
		return AddUint8E(o1.(uint8), o2.(uint8))
	case uint16:
		return AddUint16E(o1.(uint16), o2.(uint16))
	case uint32:
		return AddUint32E(o1.(uint32), o2.(uint32))
	case uint64:
		return AddUint64E(o1.(uint64), o2.(uint64))
	case decimal.Decimal:
		return o1.(decimal.Decimal).Add(o2.(decimal.Decimal)), nil
	case string:
		return o1.(string) + o2.(string), nil
	default:
		return nil, fmt.Errorf("unsupport add type '%s'", s)
	}
}

func SubAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return SubIntE(o1.(int), o2.(int))
	case int8:
		return SubInt8E(o1.(int8), o2.(int8))
	case int16:
		return SubInt16E(o1.(int16), o2.(int16))
	case int32:
		return SubInt32E(o1.(int32), o2.(int32))
	case int64:
		return SubInt64E(o1.(int64), o2.(int64))
	case uint:
		return SubUintE(o1.(uint), o2.(uint))
	case uint8:
		return SubUint8E(o1.(uint8), o2.(uint8))
	case uint16:
		return SubUint16E(o1.(uint16), o2.(uint16))
	case uint32:
		return SubUint32E(o1.(uint32), o2.(uint32))
	case uint64:
		return SubUint64E(o1.(uint64), o2.(uint64))
	case decimal.Decimal:
		return o1.(decimal.Decimal).Sub(o2.(decimal.Decimal)), nil
	default:
		return nil, fmt.Errorf("unsupport sub type '%s'", s)
	}
}

func MulAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return MulIntE(o1.(int), o2.(int))
	case int8:
		return MulInt8E(o1.(int8), o2.(int8))
	case int16:
		return MulInt16E(o1.(int16), o2.(int16))
	case int32:
		return MulInt32E(o1.(int32), o2.(int32))
	case int64:
		return MulInt64E(o1.(int64), o2.(int64))
	case uint:
		return MulUintE(o1.(uint), o2.(uint))
	case uint8:
		return MulUint8E(o1.(uint8), o2.(uint8))
	case uint16:
		return MulUint16E(o1.(uint16), o2.(uint16))
	case uint32:
		return MulUint32E(o1.(uint32), o2.(uint32))
	case uint64:
		return MulUint64E(o1.(uint64), o2.(uint64))
	case decimal.Decimal:
		return o1.(decimal.Decimal).Mul(o2.(decimal.Decimal)), nil
	default:
		return nil, fmt.Errorf("unsupport mul type '%s'", s)
	}
}

func DivAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return DivIntE(o1.(int), o2.(int))
	case int8:
		return DivInt8E(o1.(int8), o2.(int8))
	case int16:
		return DivInt16E(o1.(int16), o2.(int16))
	case int32:
		return DivInt32E(o1.(int32), o2.(int32))
	case int64:
		return DivInt64E(o1.(int64), o2.(int64))
	case uint:
		return DivUintE(o1.(uint), o2.(uint))
	case uint8:
		return DivUint8E(o1.(uint8), o2.(uint8))
	case uint16:
		return DivUint16E(o1.(uint16), o2.(uint16))
	case uint32:
		return DivUint32E(o1.(uint32), o2.(uint32))
	case uint64:
		return DivUint64E(o1.(uint64), o2.(uint64))
	case decimal.Decimal:
		return o1.(decimal.Decimal).Div(o2.(decimal.Decimal)), nil
	default:
		return nil, fmt.Errorf("unsupport div type '%s'", s)
	}
}

func CaretAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) ^ o2.(int), nil
	case int8:
		return o1.(int8) ^ o2.(int8), nil
	case int16:
		return o1.(int16) ^ o2.(int16), nil
	case int32:
		return o1.(int32) ^ o2.(int32), nil
	case int64:
		return o1.(int64) ^ o2.(int64), nil
	case uint:
		return o1.(uint) ^ o2.(uint), nil
	case uint8:
		return o1.(uint8) ^ o2.(uint8), nil
	case uint16:
		return o1.(uint16) ^ o2.(uint16), nil
	case uint32:
		return o1.(uint32) ^ o2.(uint32), nil
	case uint64:
		return o1.(uint64) ^ o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport caret type '%s'", s)
	}
}

func OrAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) | o2.(int), nil
	case int8:
		return o1.(int8) | o2.(int8), nil
	case int16:
		return o1.(int16) | o2.(int16), nil
	case int32:
		return o1.(int32) | o2.(int32), nil
	case int64:
		return o1.(int64) | o2.(int64), nil
	case uint:
		return o1.(uint) | o2.(uint), nil
	case uint8:
		return o1.(uint8) | o2.(uint8), nil
	case uint16:
		return o1.(uint16) | o2.(uint16), nil
	case uint32:
		return o1.(uint32) | o2.(uint32), nil
	case uint64:
		return o1.(uint64) | o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport or type '%s'", s)
	}
}

func AndAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) & o2.(int), nil
	case int8:
		return o1.(int8) & o2.(int8), nil
	case int16:
		return o1.(int16) & o2.(int16), nil
	case int32:
		return o1.(int32) & o2.(int32), nil
	case int64:
		return o1.(int64) & o2.(int64), nil
	case uint:
		return o1.(uint) & o2.(uint), nil
	case uint8:
		return o1.(uint8) & o2.(uint8), nil
	case uint16:
		return o1.(uint16) & o2.(uint16), nil
	case uint32:
		return o1.(uint32) & o2.(uint32), nil
	case uint64:
		return o1.(uint64) & o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport and type '%s'", s)
	}
}

func ModAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) % o2.(int), nil
	case int8:
		return o1.(int8) % o2.(int8), nil
	case int16:
		return o1.(int16) % o2.(int16), nil
	case int32:
		return o1.(int32) % o2.(int32), nil
	case int64:
		return o1.(int64) % o2.(int64), nil
	case uint:
		return o1.(uint) % o2.(uint), nil
	case uint8:
		return o1.(uint8) % o2.(uint8), nil
	case uint16:
		return o1.(uint16) % o2.(uint16), nil
	case uint32:
		return o1.(uint32) % o2.(uint32), nil
	case uint64:
		return o1.(uint64) % o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport mod type '%s'", s)
	}
}

func LeftShiftAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) << o2.(int), nil
	case int8:
		return o1.(int8) << o2.(int8), nil
	case int16:
		return o1.(int16) << o2.(int16), nil
	case int32:
		return o1.(int32) << o2.(int32), nil
	case int64:
		return o1.(int64) << o2.(int64), nil
	case uint:
		return o1.(uint) << o2.(uint), nil
	case uint8:
		return o1.(uint8) << o2.(uint8), nil
	case uint16:
		return o1.(uint16) << o2.(uint16), nil
	case uint32:
		return o1.(uint32) << o2.(uint32), nil
	case uint64:
		return o1.(uint64) << o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport left shift type '%s'", s)
	}
}

func RightShiftAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) >> o2.(int), nil
	case int8:
		return o1.(int8) >> o2.(int8), nil
	case int16:
		return o1.(int16) >> o2.(int16), nil
	case int32:
		return o1.(int32) >> o2.(int32), nil
	case int64:
		return o1.(int64) >> o2.(int64), nil
	case uint:
		return o1.(uint) >> o2.(uint), nil
	case uint8:
		return o1.(uint8) >> o2.(uint8), nil
	case uint16:
		return o1.(uint16) >> o2.(uint16), nil
	case uint32:
		return o1.(uint32) >> o2.(uint32), nil
	case uint64:
		return o1.(uint64) >> o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport right shift type '%s'", s)
	}
}

func BitClearAnyE(left, right interface{}) (interface{}, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return nil, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) &^ o2.(int), nil
	case int8:
		return o1.(int8) &^ o2.(int8), nil
	case int16:
		return o1.(int16) &^ o2.(int16), nil
	case int32:
		return o1.(int32) &^ o2.(int32), nil
	case int64:
		return o1.(int64) &^ o2.(int64), nil
	case uint:
		return o1.(uint) &^ o2.(uint), nil
	case uint8:
		return o1.(uint8) &^ o2.(uint8), nil
	case uint16:
		return o1.(uint16) &^ o2.(uint16), nil
	case uint32:
		return o1.(uint32) &^ o2.(uint32), nil
	case uint64:
		return o1.(uint64) &^ o2.(uint64), nil
	default:
		return nil, fmt.Errorf("unsupport bit clear type '%s'", s)
	}
}

func EqualAnyE(left, right interface{}) (bool, error) {
	
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	
	if o1 == nil || o2 == nil {
		return o1 == o2, nil
	}
	
	switch s := o1.(type) {
	case int:
		return o1.(int) == o2.(int), nil
	case int8:
		return o1.(int8) == o2.(int8), nil
	case int16:
		return o1.(int16) == o2.(int16), nil
	case int32:
		return o1.(int32) == o2.(int32), nil
	case int64:
		return o1.(int64) == o2.(int64), nil
	case uint:
		return o1.(uint) == o2.(uint), nil
	case uint8:
		return o1.(uint8) == o2.(uint8), nil
	case uint16:
		return o1.(uint16) == o2.(uint16), nil
	case uint32:
		return o1.(uint32) == o2.(uint32), nil
	case uint64:
		return o1.(uint64) == o2.(uint64), nil
	case string:
		return o1.(string) == o2.(string), nil
	case decimal.Decimal:
		return o1.(decimal.Decimal).Equal(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport equal type '%s'", s)
	}
}

func NotEqualAnyE(left, right interface{}) (bool, error) {
	
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	
	if o1 == nil || o2 == nil {
		return o1 != o2, nil
	}
	
	switch s := o1.(type) {
	case int:
		return o1.(int) != o2.(int), nil
	case int8:
		return o1.(int8) != o2.(int8), nil
	case int16:
		return o1.(int16) != o2.(int16), nil
	case int32:
		return o1.(int32) != o2.(int32), nil
	case int64:
		return o1.(int64) != o2.(int64), nil
	case uint:
		return o1.(uint) != o2.(uint), nil
	case uint8:
		return o1.(uint8) != o2.(uint8), nil
	case uint16:
		return o1.(uint16) != o2.(uint16), nil
	case uint32:
		return o1.(uint32) != o2.(uint32), nil
	case uint64:
		return o1.(uint64) != o2.(uint64), nil
	case string:
		return o1.(string) != o2.(string), nil
	case decimal.Decimal:
		return !o1.(decimal.Decimal).Equal(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport not equal type '%s'", s)
	}
}

func LessAnyE(left, right interface{}) (bool, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) < o2.(int), nil
	case int8:
		return o1.(int8) < o2.(int8), nil
	case int16:
		return o1.(int16) < o2.(int16), nil
	case int32:
		return o1.(int32) < o2.(int32), nil
	case int64:
		return o1.(int64) < o2.(int64), nil
	case uint:
		return o1.(uint) < o2.(uint), nil
	case uint8:
		return o1.(uint8) < o2.(uint8), nil
	case uint16:
		return o1.(uint16) < o2.(uint16), nil
	case uint32:
		return o1.(uint32) < o2.(uint32), nil
	case uint64:
		return o1.(uint64) < o2.(uint64), nil
	case string:
		return o1.(string) < o2.(string), nil
	case decimal.Decimal:
		return o1.(decimal.Decimal).LessThan(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport less type '%s'", s)
	}
}

func LessOrEqualAnyE(left, right interface{}) (bool, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) <= o2.(int), nil
	case int8:
		return o1.(int8) <= o2.(int8), nil
	case int16:
		return o1.(int16) <= o2.(int16), nil
	case int32:
		return o1.(int32) <= o2.(int32), nil
	case int64:
		return o1.(int64) <= o2.(int64), nil
	case uint:
		return o1.(uint) <= o2.(uint), nil
	case uint8:
		return o1.(uint8) <= o2.(uint8), nil
	case uint16:
		return o1.(uint16) <= o2.(uint16), nil
	case uint32:
		return o1.(uint32) <= o2.(uint32), nil
	case uint64:
		return o1.(uint64) <= o2.(uint64), nil
	case string:
		return o1.(string) <= o2.(string), nil
	case decimal.Decimal:
		return o1.(decimal.Decimal).LessThanOrEqual(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport less or equal type '%s'", s)
	}
}

func GreaterAnyE(left, right interface{}) (bool, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) > o2.(int), nil
	case int8:
		return o1.(int8) > o2.(int8), nil
	case int16:
		return o1.(int16) > o2.(int16), nil
	case int32:
		return o1.(int32) > o2.(int32), nil
	case int64:
		return o1.(int64) > o2.(int64), nil
	case uint:
		return o1.(uint) > o2.(uint), nil
	case uint8:
		return o1.(uint8) > o2.(uint8), nil
	case uint16:
		return o1.(uint16) > o2.(uint16), nil
	case uint32:
		return o1.(uint32) > o2.(uint32), nil
	case uint64:
		return o1.(uint64) > o2.(uint64), nil
	case string:
		return o1.(string) > o2.(string), nil
	case decimal.Decimal:
		return o1.(decimal.Decimal).GreaterThanOrEqual(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport greater type '%s'", s)
	}
}

func GreaterOrEqualAnyE(left, right interface{}) (bool, error) {
	o1, o2, err := ToBinOperandE(left, right)
	if err != nil {
		return false, err
	}
	switch s := o1.(type) {
	case int:
		return o1.(int) >= o2.(int), nil
	case int8:
		return o1.(int8) >= o2.(int8), nil
	case int16:
		return o1.(int16) >= o2.(int16), nil
	case int32:
		return o1.(int32) >= o2.(int32), nil
	case int64:
		return o1.(int64) >= o2.(int64), nil
	case uint:
		return o1.(uint) >= o2.(uint), nil
	case uint8:
		return o1.(uint8) >= o2.(uint8), nil
	case uint16:
		return o1.(uint16) >= o2.(uint16), nil
	case uint32:
		return o1.(uint32) >= o2.(uint32), nil
	case uint64:
		return o1.(uint64) >= o2.(uint64), nil
	case string:
		return o1.(string) >= o2.(string), nil
	case decimal.Decimal:
		return o1.(decimal.Decimal).GreaterThanOrEqual(o2.(decimal.Decimal)), nil
	default:
		return false, fmt.Errorf("unsupport greater or equal type '%s'", s)
	}
}

func UnaryPlusAnyE(v interface{}) (interface{}, error) {
	v = Indirect(v)
	switch s := v.(type) {
	case int:
		return +v.(int), nil
	case int8:
		return +v.(int8), nil
	case int16:
		return +v.(int16), nil
	case int32:
		return +v.(int32), nil
	case int64:
		return +v.(int64), nil
	case uint:
		return +v.(uint), nil
	case uint8:
		return +v.(uint8), nil
	case uint16:
		return +v.(uint16), nil
	case uint32:
		return +v.(uint32), nil
	case uint64:
		return +v.(uint64), nil
	case decimal.Decimal:
		vv := v.(decimal.Decimal).String()
		if strings.HasPrefix(vv, "-") {
			r, err := decimal.NewFromString(strings.TrimPrefix(vv, "-"))
			return r, err
		}
		return v.(decimal.Decimal), nil
	default:
		return false, fmt.Errorf("unsupport unary plus type '%s'", s)
	}
}

func UnaryMinusAnyE(v interface{}) (interface{}, error) {
	v = Indirect(v)
	switch s := v.(type) {
	case int:
		return -v.(int), nil
	case int8:
		return -v.(int8), nil
	case int16:
		return -v.(int16), nil
	case int32:
		return -v.(int32), nil
	case int64:
		return -v.(int64), nil
	case uint:
		return -v.(uint), nil
	case uint8:
		return -v.(uint8), nil
	case uint16:
		return -v.(uint16), nil
	case uint32:
		return -v.(uint32), nil
	case uint64:
		return -v.(uint64), nil
	case decimal.Decimal:
		vv := v.(decimal.Decimal).String()
		if strings.HasPrefix(vv, "-") {
			r, err := decimal.NewFromString(strings.TrimPrefix(vv, "-"))
			return r, err
		} else {
			r, err := decimal.NewFromString("-" + vv)
			return r, err
		}
	default:
		return false, fmt.Errorf("unsupport unary minus type '%s'", s)
	}
}

func UnaryCaretAnyE(v interface{}) (interface{}, error) {
	v = Indirect(v)
	switch s := v.(type) {
	case int:
		return ^v.(int), nil
	case int8:
		return ^v.(int8), nil
	case int16:
		return ^v.(int16), nil
	case int32:
		return ^v.(int32), nil
	case int64:
		return ^v.(int64), nil
	case uint:
		return ^v.(uint), nil
	case uint8:
		return ^v.(uint8), nil
	case uint16:
		return ^v.(uint16), nil
	case uint32:
		return ^v.(uint32), nil
	case uint64:
		return ^v.(uint64), nil
	default:
		return false, fmt.Errorf("unsupport unary caret type '%s'", s)
	}
}

func UnaryNotAnyE(v interface{}) (interface{}, error) {
	v = Indirect(v)
	switch s := v.(type) {
	case bool:
		return !v.(bool), nil
	default:
		return false, fmt.Errorf("unsupport unary not type '%s'", s)
	}
}

func LogicAndAnyE(left, right interface{}) (bool, error) {
	a, err := ToBoolE(left)
	if err != nil {
		return false, err
	}
	b, err := ToBoolE(right)
	if err != nil {
		return false, err
	}
	return a && b, nil
}

func LogicOrAnyE(left, right interface{}) (bool, error) {
	a, err := ToBoolE(left)
	if err != nil {
		return false, err
	}
	b, err := ToBoolE(right)
	if err != nil {
		return false, err
	}
	return a || b, nil
}
