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
	
	dm := generator.NewDataManager()
	// adm := generator.NewDataManager()
	
	//for i:=0;i<10;i++{
	{
		id := dm.NextId()
		
		val := _mock.Int64()
		err := mapper.InsertParameterBigintInt64(val)
		require.NoError(t, err, val)
		
		r1, err := mapper.SelectParameterBigintInt64(id)
		require.NoError(t, err, id)
		t.Log(id, val, r1)
	}
	{
		id := dm.NextId()
		
		val := _mock.String()
		err := mapper.InsertParameterCharacterString(val)
		require.NoError(t, err, val)
		
		r1, err := mapper.SelectParameterCharacterString(id)
		require.NoError(t, err, id)
		require.Equal(t, val, r1)
		t.Log(id, val, r1)
	}
	{
		id := dm.NextId()
		
		val := _mock.String()
		err := mapper.InsertParameterCharacterVaryingString(val)
		require.NoError(t, err, val)
		
		r1, err := mapper.SelectParameterCharacterVaryingString(id)
		require.NoError(t, err, id)
		t.Log(id, val, r1)
	}
	
	// }
}
