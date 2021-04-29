package cast

import (
	"fmt"
	"github.com/shopspring/decimal"
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
