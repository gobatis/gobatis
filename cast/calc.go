package cast

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

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

func AddInt64(a, b int64) (r int64, err error) {
	r = a + b
	if (a >= math.MaxInt32 && b >= math.MaxInt32) ||
		(a <= math.MinInt32 && b <= math.MinInt32) {
		br := BigAddInt64(a, b)
		if big.NewInt(r).Cmp(br) != 0 {
			err = resultOverFlowError("int64")
		}
	}
	return
}

func SubInt64(a, b int64) (r int64, err error) {
	r = a - b
	if a >= math.MaxUint32 && b >= math.MaxUint64 {
		br := BigSubInt64(a, b)
		if big.NewInt(r).Cmp(br) != 0 {
			err = resultOverFlowError("int64")
		}
	}
	return
}

func MulInt64(a, b int64) (r int64, err error) {
	r = a * b
	br := BigMulInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
	}
	return
}

func DivInt64(a, b int64) (r int64, err error) {
	r = a / b
	br := BigDivInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
	}
	return
}

func AddUint64(a, b uint64) (r uint64, err error) {
	r = a + b
	br := BigAddUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
	}
	return
}

func SubUint64(a, b uint64) (r uint64, err error) {
	r = a - b
	br := BigSubUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
	}
	return
}

func MulUint64(a, b uint64) (r uint64, err error) {
	r = a * b
	br := BigMulUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
	}
	return
}

func DivUint64(a, b uint64) (r uint64, err error) {
	r = a + b
	br := BigDivUint64(a, b)
	rr := &big.Int{}
	rr.SetUint64(r)
	if rr.Cmp(br) != 0 {
		err = resultOverFlowError("uint64")
	}
	return
}

func AddInt(a, b int) (r int, err error) {
	r1, err := AddInt64(int64(a), int64(b))
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

func SubInt(a, b int) (r int, err error) {
	r1, err := SubInt64(int64(a), int64(b))
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

func MulInt(a, b int) (r int, err error) {
	r1, err := MulInt64(int64(a), int64(b))
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

func DivInt(a, b int) (r int, err error) {
	r1, err := DivInt64(int64(a), int64(b))
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

func AddInt8(a, b int8) (r int8, err error) {
	r1, err := AddInt64(int64(a), int64(b))
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

func SubInt8(a, b int8) (r int8, err error) {
	r1, err := SubInt64(int64(a), int64(b))
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

func MulInt8(a, b int8) (r int8, err error) {
	r1, err := MulInt64(int64(a), int64(b))
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

func DivInt8(a, b int8) (r int8, err error) {
	r1, err := DivInt64(int64(a), int64(b))
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

func AddInt16(a, b int16) (r int16, err error) {
	r1, err := AddInt64(int64(a), int64(b))
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

func SubInt16(a, b int16) (r int16, err error) {
	r1, err := SubInt64(int64(a), int64(b))
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

func MulInt16(a, b int16) (r int16, err error) {
	r1, err := MulInt64(int64(a), int64(b))
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

func DivInt16(a, b int16) (r int16, err error) {
	r1, err := DivInt64(int64(a), int64(b))
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

func AddInt32(a, b int32) (r int32, err error) {
	r1, err := AddInt64(int64(a), int64(b))
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

func SubInt32(a, b int32) (r int32, err error) {
	r1, err := SubInt64(int64(a), int64(b))
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

func MulInt32(a, b int32) (r int32, err error) {
	r1, err := MulInt64(int64(a), int64(b))
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

func DivInt32(a, b int32) (r int32, err error) {
	r1, err := DivInt64(int64(a), int64(b))
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

func AddUint(a, b uint) (r uint, err error) {
	r1, err := AddUint64(uint64(a), uint64(b))
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

func SubUint(a, b uint) (r uint, err error) {
	r1, err := SubUint64(uint64(a), uint64(b))
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

func MulUint(a, b uint) (r uint, err error) {
	r1, err := MulUint64(uint64(a), uint64(b))
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

func DivUint(a, b uint) (r uint, err error) {
	r1, err := DivUint64(uint64(a), uint64(b))
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

func AddUint8(a, b uint8) (r uint8, err error) {
	r1, err := AddUint64(uint64(a), uint64(b))
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

func SubUint8(a, b uint8) (r uint8, err error) {
	r1, err := SubUint64(uint64(a), uint64(b))
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

func MulUint8(a, b uint8) (r uint8, err error) {
	r1, err := MulUint64(uint64(a), uint64(b))
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

func DivUint8(a, b uint8) (r uint8, err error) {
	r1, err := DivUint64(uint64(a), uint64(b))
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

func AddUint16(a, b uint16) (r uint16, err error) {
	r1, err := AddUint64(uint64(a), uint64(b))
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

func SubUint16(a, b uint16) (r uint16, err error) {
	r1, err := SubUint64(uint64(a), uint64(b))
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

func MulUint16(a, b uint16) (r uint16, err error) {
	r1, err := MulUint64(uint64(a), uint64(b))
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

func DivUint16(a, b uint16) (r uint16, err error) {
	r1, err := DivUint64(uint64(a), uint64(b))
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

func AddUint32(a, b uint32) (r uint32, err error) {
	r1, err := AddUint64(uint64(a), uint64(b))
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

func SubUint32(a, b uint32) (r uint32, err error) {
	r1, err := SubUint64(uint64(a), uint64(b))
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

func MulUint32(a, b uint32) (r uint32, err error) {
	r1, err := DivUint64(uint64(a), uint64(b))
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

func DivUint32(a, b uint32) (r uint32, err error) {
	r1, err := DivUint64(uint64(a), uint64(b))
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
