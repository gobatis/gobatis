package stmt

import (
	"fmt"
	"github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStmtMapper(t *testing.T) {
	engine := gobatis.NewEngine(gobatis.NewDB("", ""))
	engine.InitLogger()
	err := engine.RegisterBundle(gobatis.NewBundle("./sql"))
	require.NoError(t, err)
	
	mapper := new(InsertMapper)
	err = engine.BindMapper(mapper)
	require.NoError(t, err)
	
	stmt, err := mapper.InsertS001Stmt()
	require.NoError(t, err)
	fmt.Println(stmt.RealSQL())
	
	//stmt, err = mapper.InsertS002Stmt()
	//require.NoError(t, err)
	//fmt.Println(stmt.RealSQL())
	
	//stmt, err = mapper.InsertS003Stmt()
	//require.NoError(t, err)
	//fmt.Println(stmt.RealSQL())
	
	//stmt, err = mapper.InsertS004Stmt()
	//require.NoError(t, err)
	//fmt.Println(stmt.RealSQL())
	
	stmt, err = mapper.InsertS005Stmt([]Item{{B: "b1", C: "c1"}, {B: "b2", C: "c2"}})
	require.NoError(t, err)
	fmt.Println(stmt.RealSQL())
}
