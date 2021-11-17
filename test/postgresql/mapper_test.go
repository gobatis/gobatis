package postgresql

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/test/generator"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapper(t *testing.T) {
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
	err = engine.Master().Migrate(mapper)
	require.NoError(t, err)
	
	err = mapper.ResetTable()
	require.NoError(t, err)
	
	testScanTypes(t, mapper, generator.NewDataManager())
}
