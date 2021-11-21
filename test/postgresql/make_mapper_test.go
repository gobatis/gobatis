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
		vs := make([]int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBigintInt64(sid, "InsertArrayParameterBigintInt64", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBigintInt64(sid, "UpdateArrayParameterBigintInt64", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBigintInt64(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBigintInt64OriginalPointer(sid, "InsertArrayParameterBigintInt64OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBigintInt64OriginalPointer(sid, "UpdateArrayParameterBigintInt64OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterBigintInt64OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterBigintInt64PointerOriginal(sid, "InsertArrayParameterBigintInt64PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBigintInt64PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int64, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int64()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterBigintInt64PointerOriginal(sid, "UpdateArrayParameterBigintInt64PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBigintInt64PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBigintInt64PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt8Int8(sid, "InsertArrayParameterInt8Int8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt8Int8(sid, "UpdateArrayParameterInt8Int8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt8Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt8Int8OriginalPointer(sid, "InsertArrayParameterInt8Int8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt8Int8OriginalPointer(sid, "UpdateArrayParameterInt8Int8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt8Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterInt8Int8PointerOriginal(sid, "InsertArrayParameterInt8Int8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt8Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterInt8Int8PointerOriginal(sid, "UpdateArrayParameterInt8Int8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt8Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt8Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBooleanBool(sid, "InsertArrayParameterBooleanBool", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBooleanBool(sid, "UpdateArrayParameterBooleanBool", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBooleanBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBooleanBoolOriginalPointer(sid, "InsertArrayParameterBooleanBoolOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBooleanBoolOriginalPointer(sid, "UpdateArrayParameterBooleanBoolOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterBooleanBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterBooleanBoolPointerOriginal(sid, "InsertArrayParameterBooleanBoolPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBooleanBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterBooleanBoolPointerOriginal(sid, "UpdateArrayParameterBooleanBoolPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBooleanBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBooleanBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBoolBool(sid, "InsertArrayParameterBoolBool", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBoolBool(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBoolBool(sid, "UpdateArrayParameterBoolBool", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBoolBool(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBoolBool(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterBoolBoolOriginalPointer(sid, "InsertArrayParameterBoolBoolOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterBoolBoolOriginalPointer(sid, "UpdateArrayParameterBoolBoolOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterBoolBoolOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterBoolBoolPointerOriginal(sid, "InsertArrayParameterBoolBoolPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterBoolBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*bool, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Bool()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterBoolBoolPointerOriginal(sid, "UpdateArrayParameterBoolBoolPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterBoolBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterBoolBoolPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharacterString(sid, "InsertArrayParameterCharacterString", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharacterString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterString(sid, "UpdateArrayParameterCharacterString", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharacterString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharacterStringOriginalPointer(sid, "InsertArrayParameterCharacterStringOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterStringOriginalPointer(sid, "UpdateArrayParameterCharacterStringOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterCharacterStringPointerOriginal(sid, "InsertArrayParameterCharacterStringPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharacterStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterStringPointerOriginal(sid, "UpdateArrayParameterCharacterStringPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharacterStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharString(sid, "InsertArrayParameterCharString", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharString(sid, "UpdateArrayParameterCharString", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharStringOriginalPointer(sid, "InsertArrayParameterCharStringOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharStringOriginalPointer(sid, "UpdateArrayParameterCharStringOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterCharStringPointerOriginal(sid, "InsertArrayParameterCharStringPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterCharStringPointerOriginal(sid, "UpdateArrayParameterCharStringPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharacterVaryingString(sid, "InsertArrayParameterCharacterVaryingString", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterVaryingString(sid, "UpdateArrayParameterCharacterVaryingString", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterVaryingString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterCharacterVaryingStringOriginalPointer(sid, "InsertArrayParameterCharacterVaryingStringOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterVaryingStringOriginalPointer(sid, "UpdateArrayParameterCharacterVaryingStringOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterVaryingStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterCharacterVaryingStringPointerOriginal(sid, "InsertArrayParameterCharacterVaryingStringPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterCharacterVaryingStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterCharacterVaryingStringPointerOriginal(sid, "UpdateArrayParameterCharacterVaryingStringPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterCharacterVaryingStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterCharacterVaryingStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterVarcharString(sid, "InsertArrayParameterVarcharString", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterVarcharString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterVarcharString(sid, "UpdateArrayParameterVarcharString", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterVarcharString(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterVarcharString(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterVarcharStringOriginalPointer(sid, "InsertArrayParameterVarcharStringOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterVarcharStringOriginalPointer(sid, "UpdateArrayParameterVarcharStringOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterVarcharStringOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterVarcharStringPointerOriginal(sid, "InsertArrayParameterVarcharStringPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterVarcharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*string, 0)
		for i := 0; i < 3; i++ {
			v := _mock.String()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterVarcharStringPointerOriginal(sid, "UpdateArrayParameterVarcharStringPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterVarcharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterVarcharStringPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterFloat8Float32(sid, "InsertArrayParameterFloat8Float32", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterFloat8Float32(sid, "UpdateArrayParameterFloat8Float32", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat8Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterFloat8Float32OriginalPointer(sid, "InsertArrayParameterFloat8Float32OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterFloat8Float32OriginalPointer(sid, "UpdateArrayParameterFloat8Float32OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat8Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterFloat8Float32PointerOriginal(sid, "InsertArrayParameterFloat8Float32PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterFloat8Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterFloat8Float32PointerOriginal(sid, "UpdateArrayParameterFloat8Float32PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterFloat8Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat8Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterIntegerInt8(sid, "InsertArrayParameterIntegerInt8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterIntegerInt8(sid, "UpdateArrayParameterIntegerInt8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntegerInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterIntegerInt8OriginalPointer(sid, "InsertArrayParameterIntegerInt8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterIntegerInt8OriginalPointer(sid, "UpdateArrayParameterIntegerInt8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntegerInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterIntegerInt8PointerOriginal(sid, "InsertArrayParameterIntegerInt8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterIntegerInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterIntegerInt8PointerOriginal(sid, "UpdateArrayParameterIntegerInt8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterIntegerInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntegerInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterIntInt8(sid, "InsertArrayParameterIntInt8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterIntInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterIntInt8(sid, "UpdateArrayParameterIntInt8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterIntInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterIntInt8OriginalPointer(sid, "InsertArrayParameterIntInt8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterIntInt8OriginalPointer(sid, "UpdateArrayParameterIntInt8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterIntInt8PointerOriginal(sid, "InsertArrayParameterIntInt8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterIntInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterIntInt8PointerOriginal(sid, "UpdateArrayParameterIntInt8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterIntInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterIntInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt4Int8(sid, "InsertArrayParameterInt4Int8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt4Int8(sid, "UpdateArrayParameterInt4Int8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt4Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt4Int8OriginalPointer(sid, "InsertArrayParameterInt4Int8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt4Int8OriginalPointer(sid, "UpdateArrayParameterInt4Int8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt4Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterInt4Int8PointerOriginal(sid, "InsertArrayParameterInt4Int8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt4Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterInt4Int8PointerOriginal(sid, "UpdateArrayParameterInt4Int8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt4Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt4Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterNumericDecimal(sid, "InsertArrayParameterNumericDecimal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterNumericDecimal(sid, "UpdateArrayParameterNumericDecimal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterNumericDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterNumericDecimalOriginalPointer(sid, "InsertArrayParameterNumericDecimalOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterNumericDecimalOriginalPointer(sid, "UpdateArrayParameterNumericDecimalOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterNumericDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterNumericDecimalPointerOriginal(sid, "InsertArrayParameterNumericDecimalPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterNumericDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterNumericDecimalPointerOriginal(sid, "UpdateArrayParameterNumericDecimalPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterNumericDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterNumericDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterDecimalDecimal(sid, "InsertArrayParameterDecimalDecimal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterDecimalDecimal(sid, "UpdateArrayParameterDecimalDecimal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterDecimalDecimal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterDecimalDecimalOriginalPointer(sid, "InsertArrayParameterDecimalDecimalOriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterDecimalDecimalOriginalPointer(sid, "UpdateArrayParameterDecimalDecimalOriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterDecimalDecimalOriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterDecimalDecimalPointerOriginal(sid, "InsertArrayParameterDecimalDecimalPointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterDecimalDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*decimal, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Decimal()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterDecimalDecimalPointerOriginal(sid, "UpdateArrayParameterDecimalDecimalPointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterDecimalDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterDecimalDecimalPointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterFloat4Float32(sid, "InsertArrayParameterFloat4Float32", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterFloat4Float32(sid, "UpdateArrayParameterFloat4Float32", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat4Float32(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterFloat4Float32OriginalPointer(sid, "InsertArrayParameterFloat4Float32OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterFloat4Float32OriginalPointer(sid, "UpdateArrayParameterFloat4Float32OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat4Float32OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterFloat4Float32PointerOriginal(sid, "InsertArrayParameterFloat4Float32PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterFloat4Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*float32, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Float32()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterFloat4Float32PointerOriginal(sid, "UpdateArrayParameterFloat4Float32PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterFloat4Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterFloat4Float32PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterSmallintInt8(sid, "InsertArrayParameterSmallintInt8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterSmallintInt8(sid, "UpdateArrayParameterSmallintInt8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterSmallintInt8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterSmallintInt8OriginalPointer(sid, "InsertArrayParameterSmallintInt8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterSmallintInt8OriginalPointer(sid, "UpdateArrayParameterSmallintInt8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterSmallintInt8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterSmallintInt8PointerOriginal(sid, "InsertArrayParameterSmallintInt8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterSmallintInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterSmallintInt8PointerOriginal(sid, "UpdateArrayParameterSmallintInt8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterSmallintInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterSmallintInt8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt2Int8(sid, "InsertArrayParameterInt2Int8", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt2Int8(sid, "UpdateArrayParameterInt2Int8", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt2Int8(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

	}
	{

		sid := manager.NextId()
		vs := make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err := mapper.InsertArrayParameterInt2Int8OriginalPointer(sid, "InsertArrayParameterInt2Int8OriginalPointer", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, *rows)

		r, err := mapper.SelectArrayParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		vs = make([]int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, v)
		}
		rows, err = mapper.UpdateArrayParameterInt2Int8OriginalPointer(sid, "UpdateArrayParameterInt2Int8OriginalPointer", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, *rows)

		r, err = mapper.SelectArrayParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, vs[i], *r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt2Int8OriginalPointer(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, *rows)

	}
	{

		sid := manager.NextId()
		vs := make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err := mapper.InsertArrayParameterInt2Int8PointerOriginal(sid, "InsertArrayParameterInt2Int8PointerOriginal", vs)
		require.NoError(t, err, vs)
		require.Equal(t, 1, rows)

		r, err := mapper.SelectArrayParameterInt2Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		vs = make([]*int8, 0)
		for i := 0; i < 3; i++ {
			v := _mock.Int8()
			vs = append(vs, &v)
		}
		rows, err = mapper.UpdateArrayParameterInt2Int8PointerOriginal(sid, "UpdateArrayParameterInt2Int8PointerOriginal", vs)
		require.NoError(t, err, sid, vs)
		require.Equal(t, 1, rows)

		r, err = mapper.SelectArrayParameterInt2Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		for i := 0; i < 3; i++ {
			require.Equal(t, *vs[i], r[i])
		}

		rows, err = mapper.DeleteArrayParameterInt2Int8PointerOriginal(sid)
		require.NoError(t, err, sid)
		require.Equal(t, 1, rows)

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
