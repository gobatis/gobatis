package postgresql

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsert(t *testing.T) {
	var err error
	mapper := new(MakeMapper)
	engine := postgresql.NewEngine("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	err = engine.Init(gobatis.NewBundle("./sql"))
	require.NoError(t, err)
	err = engine.BindMapper(mapper)
	require.NoError(t, err)
	
	defer func() {
		engine.Close()
	}()
	
	//for i := 0; i < 10; i++ {
	//
	//	err = mapper.InsertParameterBigintInt64(_mock.Int64())
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertParameterBigintInt64PointerOriginal(pointer.ToInt64(_mock.Int64()))
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterBigintInt64([]int64{_mock.Int64(), _mock.Int64(), _mock.Int64()})
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterBigintInt64PointerOriginal([]*int64{pointer.ToInt64(_mock.Int64()), pointer.ToInt64(_mock.Int64()), pointer.ToInt64(_mock.Int64())})
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertParameterCharacterString(_mock.String())
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertParameterCharacterStringPointerOriginal(pointer.ToString(_mock.String()))
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterCharacterString([]string{_mock.String(), _mock.String(), _mock.String()})
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterCharacterStringPointerOriginal([]*string{pointer.ToString(_mock.String()), pointer.ToString(_mock.String()), pointer.ToString(_mock.String())})
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertParameterCharacterVaryingString(_mock.String())
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertParameterCharacterVaryingStringPointerOriginal(pointer.ToString(_mock.String()))
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterCharacterVaryingString([]string{_mock.String(), _mock.String(), _mock.String()})
	//	require.NoError(t, err)
	//
	//	err = mapper.InsertArrayParameterCharacterVaryingStringPointerOriginal([]*string{pointer.ToString(_mock.String()), pointer.ToString(_mock.String()), pointer.ToString(_mock.String())})
	//	require.NoError(t, err)
	//
	//}
	
	res, err := mapper.SelectArrayParameterBigintInt64(1)
	require.NoError(t, err)
	t.Log(res)
	
	res2, err := mapper.SelectArrayParameterBigintInt64OriginalPointer(1)
	require.NoError(t, err)
	for _, v := range res2 {
		t.Log(*v)
	}
}
