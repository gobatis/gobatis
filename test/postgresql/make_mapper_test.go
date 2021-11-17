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

}
