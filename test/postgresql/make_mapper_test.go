package postgresql

import (
	"github.com/AlekSi/pointer"
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
	
	dm := generator.NewDataManager()
	// adm := generator.NewDataManager()
	
	//for i := 0; i < 10; i++ {
	{
		err = mapper.InsertParameterBigintInt64(_mock.Int64())
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterBigintInt64(id)
		require.NoError(t, err, id)
		t.Log(id, res, err)
	}
	{
		err = mapper.InsertParameterBigintInt64PointerOriginal(pointer.ToInt64(_mock.Int64()))
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterBigintInt64OriginalPointer(id)
		require.NoError(t, err, id)
		t.Log(id, *res, err)
	}
	{
		err = mapper.InsertParameterCharacterString(_mock.String())
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterCharacterString(id)
		require.NoError(t, err, id)
		t.Log(id, res, err)
	}
	{
		err = mapper.InsertParameterCharacterStringPointerOriginal(pointer.ToString(_mock.String()))
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterCharacterStringOriginalPointer(id)
		require.NoError(t, err, id)
		t.Log(id, res)
	}
	{
		err = mapper.InsertParameterCharacterVaryingString(_mock.String())
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterCharacterVaryingString(id)
		require.NoError(t, err, id)
		t.Log(id, res)
	}
	{
		err = mapper.InsertParameterCharacterVaryingStringPointerOriginal(pointer.ToString(_mock.String()))
		require.NoError(t, err)
	}
	{
		id := dm.NextId()
		res, err := mapper.SelectParameterCharacterVaryingStringOriginalPointer(id)
		require.NoError(t, err, id)
		t.Log(id, res)
	}
	
	//}
}
