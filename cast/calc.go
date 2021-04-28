package cast

import (
	"fmt"
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

func BigModInt64(a, b int64) *big.Int {
	r := big.NewInt(0)
	return r.Mod(big.NewInt(a), big.NewInt(b))
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

func BigModUint64(a, b uint64) *big.Int {
	r := big.NewInt(0)
	ba := &big.Int{}
	ba.SetUint64(a)
	bb := &big.Int{}
	bb.SetUint64(b)
	return r.Div(ba, bb)
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

func ModInt64(a, b int64) (r int64, err error) {
	r = a % b
	br := BigModInt64(a, b)
	if big.NewInt(r).Cmp(br) != 0 {
		err = resultOverFlowError("int64")
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
