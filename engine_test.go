package gobatis

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"
)

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

	//pwd, err := os.Getwd()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/angel?connect_timeout=10&sslmode=disable")
	//engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	////engine.BindMapper(&testMapper)
	//
	//err = engine.Init()
	//if err != nil {
	//	t.Error("DB init error:", err)
	//	return
	//}
	//err = engine.master.Ping()
	//if err != nil {
	//	t.Error("DB ping error:", err)
	//	return
	//}
	//defer func() {
	//	err = engine.master.Close()
	//	if err != nil {
	//		t.Error("DB close error:", err)
	//		return
	//	}
	//}()
	//var a driver.ExecerContext
	b := new(sql.DB)

	itf := reflect.TypeOf((*driver.QueryerContext)(nil)).Elem()
	fmt.Println(itf)
	fmt.Println(reflect.TypeOf(b).Implements(itf))
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
