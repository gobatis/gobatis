package postgresql

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/gobatis/gobatis/test/generator"
	"github.com/stretchr/testify/require"
	"sync"
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
	//err = engine.Master().Migrate(mapper)
	err = mapper.Migrate()
	require.NoError(t, err)
	
	err = mapper.ResetTable()
	require.NoError(t, err)
	
	manager := generator.NewDataManager()
	n := 1
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			testScanTypes(t, mapper, manager)
		}()
	}
	wg.Wait()
	
}
