package gobatis

import (
	"github.com/gobatis/gobatis/bundle"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func rv(a interface{}) reflect.Value {
	return reflect.ValueOf(a)
}

func TestEngine(t *testing.T) {

	//type Test struct {
	//	Id       string          `sql:"id"`
	//	Name     string          `sql:"name"`
	//	Duration decimal.Decimal `sql:"duration"`
	//}
	//
	//type TestMapper struct {
	//	SelectTestById func(id int64) (*Test, error)
	//}

	//var testMapper TestMapper

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/angel?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	////engine.BindMapper(&testMapper)
	//
	err = engine.Init()
	if err != nil {
		t.Error("DB init error:", err)
		return
	}
	err = engine.master.Ping()
	if err != nil {
		t.Error("DB ping error:", err)
		return
	}
	defer func() {
		err = engine.master.Close()
		if err != nil {
			t.Error("DB close error:", err)
			return
		}
	}()

	var names []string
	err = engine.Call("SelectTestById", 29, "hi", 2).Scan(&names)
	require.NoError(t, err)
	t.Log("result is:", names)

	//require.True(t, ok)

	//res := f.call(rv(1), rv("hello"), rv(2))
	//for _, v := range res {
	//	fmt.Println(v.Interface())
	//}
	//fmt.Println("done", res)

	//var a driver.ExecerContext
	//fmt.Println("ok:", reflect.TypeOf(engine.master).Implements(reflect.TypeOf(a)))
	//tx, _ := engine.master.Begin()
	//tx.Query()
	//smat, _ := engine.master.Prepare()
	//engine.master.Query()
	//engine.master.Conn()
	//test, err := engine.Call("SelectTestById", 0, "gobatis", 10)
	//test, err := engine.Call("SelectTestByName", 0, "", 10)
	//test, err := engine.Call("SelectTestForeach", []string{"a", "b", "c"})
	//if err != nil {
	//	t.Error("Call SelectTestById error:", err)
	//	return
	//}
	//d, _ := json.Marshal(test)
	//t.Log("query result:", string(d))
}
