package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"{% endfor %}
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
    {% for testCase in Cases %}{
        {{ testCase.Code }}
    }
    {% endfor %}
    // }
}

