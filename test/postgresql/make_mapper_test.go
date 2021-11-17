package postgresql

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/test/generator"
	"github.com/gozelle/_mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsert(t *testing.T) {
	mapper := &Mapper{
		MakeMapper: &MakeMapper{},
	}
	engine := postgresql.NewEngine("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	err := engine.Init(gobatis.NewBundle("./sql"))
	require.NoError(t, err)
	err = engine.BindMapper(mapper)
	require.NoError(t, err)

	defer func() {
		engine.Close()
	}()

	err = mapper.ResetTable()
	require.NoError(t, err)

	sid := &generator.SID{}

	//for i:=0;i<10;i++{
	{

		id1 := sid.NextId()
		v1 := _mock.Int64()
		err = mapper.InsertParameterBigintInt64(id1, "InsertParameterBigintInt64", v1)
		require.NoError(t, err, v1)

		r1, err := mapper.SelectParameterBigintInt64(id1)
		require.NoError(t, err, id1)
		require.Equal(t, v1, r1)

		id2 := sid.NextId()
		v2 := _mock.Int64()
		err = mapper.InsertParameterBigintInt64OriginalPointer(id2, "InsertParameterBigintInt64OriginalPointer", v2)
		require.NoError(t, err, v2)

		r2, err := mapper.SelectParameterBigintInt64OriginalPointer(id2)
		require.NoError(t, err, id2)
		require.Equal(t, v2, *r2)

	}
	{

		id1 := sid.NextId()
		v1 := _mock.String()
		err = mapper.InsertParameterCharacterString(id1, "InsertParameterCharacterString", v1)
		require.NoError(t, err, v1)

		r1, err := mapper.SelectParameterCharacterString(id1)
		require.NoError(t, err, id1)
		require.Equal(t, v1, r1)

		id2 := sid.NextId()
		v2 := _mock.String()
		err = mapper.InsertParameterCharacterStringOriginalPointer(id2, "InsertParameterCharacterStringOriginalPointer", v2)
		require.NoError(t, err, v2)

		r2, err := mapper.SelectParameterCharacterStringOriginalPointer(id2)
		require.NoError(t, err, id2)
		require.Equal(t, v2, *r2)

	}
	{

		id1 := sid.NextId()
		v1 := _mock.String()
		err = mapper.InsertParameterCharacterVaryingString(id1, "InsertParameterCharacterVaryingString", v1)
		require.NoError(t, err, v1)

		r1, err := mapper.SelectParameterCharacterVaryingString(id1)
		require.NoError(t, err, id1)
		require.Equal(t, v1, r1)

		id2 := sid.NextId()
		v2 := _mock.String()
		err = mapper.InsertParameterCharacterVaryingStringOriginalPointer(id2, "InsertParameterCharacterVaryingStringOriginalPointer", v2)
		require.NoError(t, err, v2)

		r2, err := mapper.SelectParameterCharacterVaryingStringOriginalPointer(id2)
		require.NoError(t, err, id2)
		require.Equal(t, v2, *r2)

	}

	// }
}
