package postgresql

import (
	"github.com/gobatis/gobatis/test/generator"
	"github.com/gozelle/_mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func testScanTypes(t *testing.T, mapper *Mapper, manager *generator.DataManager) {
	{

		sid := manager.NextId()
		v := _mock.Int64()
		rows, err := mapper.InsertParameterBigintInt64(sid, "InsertParameterBigintInt64", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int64()
		rows, err = mapper.UpdateParameterBigintInt64(sid, "UpdateParameterBigintInt64", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int64()
		rows, err := mapper.InsertParameterBigintInt64OriginalPointer(sid, "InsertParameterBigintInt64OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int64()
		rows, err = mapper.UpdateParameterBigintInt64OriginalPointer(sid, "UpdateParameterBigintInt64OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt8Int8(sid, "InsertParameterInt8Int8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt8Int8(sid, "UpdateParameterInt8Int8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt8Int8OriginalPointer(sid, "InsertParameterInt8Int8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt8Int8OriginalPointer(sid, "UpdateParameterInt8Int8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Bool()
		rows, err := mapper.InsertParameterBooleanBool(sid, "InsertParameterBooleanBool", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Bool()
		rows, err = mapper.UpdateParameterBooleanBool(sid, "UpdateParameterBooleanBool", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Bool()
		rows, err := mapper.InsertParameterBooleanBoolOriginalPointer(sid, "InsertParameterBooleanBoolOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Bool()
		rows, err = mapper.UpdateParameterBooleanBoolOriginalPointer(sid, "UpdateParameterBooleanBoolOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Bool()
		rows, err := mapper.InsertParameterBoolBool(sid, "InsertParameterBoolBool", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterBoolBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Bool()
		rows, err = mapper.UpdateParameterBoolBool(sid, "UpdateParameterBoolBool", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterBoolBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterBoolBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Bool()
		rows, err := mapper.InsertParameterBoolBoolOriginalPointer(sid, "InsertParameterBoolBoolOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Bool()
		rows, err = mapper.UpdateParameterBoolBoolOriginalPointer(sid, "UpdateParameterBoolBoolOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharacterString(sid, "InsertParameterCharacterString", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterCharacterString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharacterString(sid, "UpdateParameterCharacterString", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterCharacterString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterCharacterString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharacterStringOriginalPointer(sid, "InsertParameterCharacterStringOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharacterStringOriginalPointer(sid, "UpdateParameterCharacterStringOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharString(sid, "InsertParameterCharString", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterCharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharString(sid, "UpdateParameterCharString", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterCharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterCharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharStringOriginalPointer(sid, "InsertParameterCharStringOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharStringOriginalPointer(sid, "UpdateParameterCharStringOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharacterVaryingString(sid, "InsertParameterCharacterVaryingString", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharacterVaryingString(sid, "UpdateParameterCharacterVaryingString", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterCharacterVaryingStringOriginalPointer(sid, "InsertParameterCharacterVaryingStringOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterCharacterVaryingStringOriginalPointer(sid, "UpdateParameterCharacterVaryingStringOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterVarcharString(sid, "InsertParameterVarcharString", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterVarcharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterVarcharString(sid, "UpdateParameterVarcharString", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterVarcharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterVarcharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterVarcharStringOriginalPointer(sid, "InsertParameterVarcharStringOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterVarcharStringOriginalPointer(sid, "UpdateParameterVarcharStringOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Float32()
		rows, err := mapper.InsertParameterFloat8Float32(sid, "InsertParameterFloat8Float32", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Float32()
		rows, err = mapper.UpdateParameterFloat8Float32(sid, "UpdateParameterFloat8Float32", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Float32()
		rows, err := mapper.InsertParameterFloat8Float32OriginalPointer(sid, "InsertParameterFloat8Float32OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Float32()
		rows, err = mapper.UpdateParameterFloat8Float32OriginalPointer(sid, "UpdateParameterFloat8Float32OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterIntegerInt8(sid, "InsertParameterIntegerInt8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterIntegerInt8(sid, "UpdateParameterIntegerInt8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterIntegerInt8OriginalPointer(sid, "InsertParameterIntegerInt8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterIntegerInt8OriginalPointer(sid, "UpdateParameterIntegerInt8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterIntInt8(sid, "InsertParameterIntInt8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterIntInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterIntInt8(sid, "UpdateParameterIntInt8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterIntInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterIntInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterIntInt8OriginalPointer(sid, "InsertParameterIntInt8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterIntInt8OriginalPointer(sid, "UpdateParameterIntInt8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt4Int8(sid, "InsertParameterInt4Int8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt4Int8(sid, "UpdateParameterInt4Int8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt4Int8OriginalPointer(sid, "InsertParameterInt4Int8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt4Int8OriginalPointer(sid, "UpdateParameterInt4Int8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Decimal()
		rows, err := mapper.InsertParameterNumericDecimal(sid, "InsertParameterNumericDecimal", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Decimal()
		rows, err = mapper.UpdateParameterNumericDecimal(sid, "UpdateParameterNumericDecimal", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Decimal()
		rows, err := mapper.InsertParameterNumericDecimalOriginalPointer(sid, "InsertParameterNumericDecimalOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Decimal()
		rows, err = mapper.UpdateParameterNumericDecimalOriginalPointer(sid, "UpdateParameterNumericDecimalOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Decimal()
		rows, err := mapper.InsertParameterDecimalDecimal(sid, "InsertParameterDecimalDecimal", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Decimal()
		rows, err = mapper.UpdateParameterDecimalDecimal(sid, "UpdateParameterDecimalDecimal", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Decimal()
		rows, err := mapper.InsertParameterDecimalDecimalOriginalPointer(sid, "InsertParameterDecimalDecimalOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Decimal()
		rows, err = mapper.UpdateParameterDecimalDecimalOriginalPointer(sid, "UpdateParameterDecimalDecimalOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Float32()
		rows, err := mapper.InsertParameterFloat4Float32(sid, "InsertParameterFloat4Float32", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Float32()
		rows, err = mapper.UpdateParameterFloat4Float32(sid, "UpdateParameterFloat4Float32", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Float32()
		rows, err := mapper.InsertParameterFloat4Float32OriginalPointer(sid, "InsertParameterFloat4Float32OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Float32()
		rows, err = mapper.UpdateParameterFloat4Float32OriginalPointer(sid, "UpdateParameterFloat4Float32OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterSmallintInt8(sid, "InsertParameterSmallintInt8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterSmallintInt8(sid, "UpdateParameterSmallintInt8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterSmallintInt8OriginalPointer(sid, "InsertParameterSmallintInt8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterSmallintInt8OriginalPointer(sid, "UpdateParameterSmallintInt8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt2Int8(sid, "InsertParameterInt2Int8", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt2Int8(sid, "UpdateParameterInt2Int8", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Int8()
		rows, err := mapper.InsertParameterInt2Int8OriginalPointer(sid, "InsertParameterInt2Int8OriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.Int8()
		rows, err = mapper.UpdateParameterInt2Int8OriginalPointer(sid, "UpdateParameterInt2Int8OriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterTextString(sid, "InsertParameterTextString", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectParameterTextString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterTextString(sid, "UpdateParameterTextString", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectParameterTextString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, r)

		rows, err = mapper.DeleteParameterTextString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.String()
		rows, err := mapper.InsertParameterTextStringOriginalPointer(sid, "InsertParameterTextStringOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectParameterTextStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		v = _mock.String()
		rows, err = mapper.UpdateParameterTextStringOriginalPointer(sid, "UpdateParameterTextStringOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectParameterTextStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, v, *r)

		rows, err = mapper.DeleteParameterTextStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimeTime(sid, "InsertParameterTimeTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimeTime(sid, "UpdateParameterTimeTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimeTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimeTimeOriginalPointer(sid, "InsertParameterTimeTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimeTimeOriginalPointer(sid, "UpdateParameterTimeTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimeTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimeWithTimezoneTime(sid, "InsertParameterTimeWithTimezoneTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimeWithTimezoneTime(sid, "UpdateParameterTimeWithTimezoneTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimeWithTimezoneTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimeWithTimezoneTimeOriginalPointer(sid, "InsertParameterTimeWithTimezoneTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimeWithTimezoneTimeOriginalPointer(sid, "UpdateParameterTimeWithTimezoneTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimeWithTimezoneTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimetzTime(sid, "InsertParameterTimetzTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimetzTime(sid, "UpdateParameterTimetzTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimetzTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimetzTimeOriginalPointer(sid, "InsertParameterTimetzTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimetzTimeOriginalPointer(sid, "UpdateParameterTimetzTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimetzTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestampTime(sid, "InsertParameterTimestampTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestampTime(sid, "UpdateParameterTimestampTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimestampTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestampTimeOriginalPointer(sid, "InsertParameterTimestampTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestampTimeOriginalPointer(sid, "UpdateParameterTimestampTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimestampTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestampWithTimezoneTime(sid, "InsertParameterTimestampWithTimezoneTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestampWithTimezoneTime(sid, "UpdateParameterTimestampWithTimezoneTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimestampWithTimezoneTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestampWithTimezoneTimeOriginalPointer(sid, "InsertParameterTimestampWithTimezoneTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestampWithTimezoneTimeOriginalPointer(sid, "UpdateParameterTimestampWithTimezoneTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimestampWithTimezoneTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestamptzTime(sid, "InsertParameterTimestamptzTime", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestamptzTime(sid, "UpdateParameterTimestamptzTime", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, rows)

		rows, err = mapper.DeleteParameterTimestamptzTime(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		v := _mock.Time()
		rows, err := mapper.InsertParameterTimestamptzTimeOriginalPointer(sid, "InsertParameterTimestamptzTimeOriginalPointer", v)
		require.NoError(t, err, v)
		require.Equal(t, 1, *rows)

		v = _mock.Time()
		rows, err = mapper.UpdateParameterTimestamptzTimeOriginalPointer(sid, "UpdateParameterTimestamptzTimeOriginalPointer", v)
		require.NoError(t, err, sid, v)
		require.Equal(t, 1, *rows)

		rows, err = mapper.DeleteParameterTimestamptzTimeOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}

}
