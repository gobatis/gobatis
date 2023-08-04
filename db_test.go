package gobatis

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
	
	"github.com/AlekSi/pointer"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gobatis/gobatis/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	pwd string
)

type StmtMapper struct {
	TestInsertStmtTx  func(tx *DB, user *test.User) error
	TestInsertStmt2Tx func(tx *DB, user *test.User) error
	TestQueryStmtTx   func(tx *DB, name string, age int64) ([]*test.User, error)
	TestQueryStmt2Tx  func(tx *DB, name string, age int64) ([]*test.User, error)
}

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func rv(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func (User) TableName() string {
	return "users"
}

func TestDBQuery(t *testing.T) {
	db, err := Open(postgres.Open(""))
	require.NoError(t, err)
	
	var user User
	//err = db.Query(Background(), `
	//	select * from users where id = ${ id }
	//`, Param("id", 21)).Scan(&user)
	
	err = db.Execute(Background(), `
		insert into users(name,age) values(
		<foreach collection="enums" separator="," open="'" close="'">
            ${item.Name}, ${item.Age}
        </foreach>		             
		)
	`).Scan()
	require.NoError(t, err)
	
	var users []User
	err = db.Execute(Background(), `
		select * from messages where cid in ${ids} 
	`, Param("ids", []int64{1, 2, 34})).Scan(&users)
	require.NoError(t, err)
	
	err = db.Insert(Background(), "users", user, nil).Error()
	if err != nil {
		return
	}
	
	err = db.Update(Background(), "tests", map[string]any{
		"name": "name",
		"age":  18,
	}, Where("age > ${ age }", Param("age", 18))).Error()
	if err != nil {
		return
	}
	
	err = db.InsertBatch(Background(), "users", user, 10, nil).Error()
	if err != nil {
		return
	}
	
	err = db.Delete(Background(), user.TableName(), Where("id = ${user.Id}", Param("user", user))).Error()
	if err != nil {
		return
	}
	
	require.NoError(t, err)
}

func TestDBInsert(t *testing.T) {
	db := &DB{}
	
	var users []User
	var total int64
	err := db.Build(
		Background(), Select("*").
			From("users").
			Where("agg > 0").
			Page(0, 10).
			Count("*"),
	).Scan(&users, &total)
	require.NoError(t, err)
	
	err = db.Build(
		Background(), Raw(`
				select * from a left join b on a.name = b.name where a.age > 18
            `, Param("a", map[string]any{})).
			Scroll(10, And("a.ab > #{ age }", Param("age", 18))).
			Page(0, 10).
			Count("*"),
	).Scan(&users, &total)
	require.NoError(t, err)
	
	err = db.Update(Background(), "", map[string]any{}, Where("")).Error()
	if err != nil {
		return
	}
	
	return
}

func TestDBQuery45(t *testing.T) {
	db := &DB{}
	
	err := db.Execute(Background(),
		`update public.users where id = #{a} and name = #{b}`,
		Param("a", 10),
		Param("b", 10),
	).Error()
	if err != nil {
		return
	}
	
	require.NoError(t, err)
}

func TestDBQuery2(t *testing.T) {
	db := &DB{}
	
	var user User
	err := db.Query(
		Background(),
		`select * from users where <if test="age >= 10"> id = #{ age  }</if>`,
		Param("age", 10),
	).Scan(&user)
	
	require.NoError(t, err)
}

//func TestEngine(t *testing.T) {
//
//	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
//	err := engine.Init(NewBundle("test"))
//	require.NoError(t, err)
//
//	err = engine.master.Ping()
//	require.NoError(t, err)
//
//	defer func() {
//		err = engine.master.Close()
//		require.NoError(t, err)
//	}()
//
//	_testMapper := new(test.TestMapper)
//	err = engine.BindMapper(_testMapper)
//	require.NoError(t, err)
//
//	testSelectInsert(t, _testMapper)
//	testSelectInsertPointer(t, _testMapper)
//	testSelectInsertForeachSlice(t, _testMapper)
//	testSelectInsertForeachSlicePointer(t, _testMapper)
//	testSelectInsertForeachMap(t, _testMapper)
//	testSelectInsertForeachMapPointer(t, _testMapper)
//	testSelectInsertForeachStruct(t, _testMapper)
//	testSelectInsertForeachStructPointer(t, _testMapper)
//	//testSelectInsertContextTx(t, engine, _testMapper)
//	testInsert(t, _testMapper)
//	testSelectRow(t, _testMapper)
//	testSelectRowPointer(t, _testMapper)
//	testSelectRows(t, _testMapper)
//	testSelectRowsPointer(t, _testMapper)
//	testSelectStruct(t, _testMapper)
//	testSelectStructs(t, _testMapper)
//}

//func TestStmt(t *testing.T) {
//	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
//	err := engine.Init(NewBundle("test/sql"))
//	require.NoError(t, err)
//	err = engine.master.Ping()
//	require.NoError(t, err)
//	engine.SetLogLevel(DebugLevel)
//	defer func() {
//		engine.Close()
//	}()
//
//	stmtMapper := new(test.StmtMapper)
//	err = engine.BindMapper(stmtMapper)
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt(&test.User{
//		Name: "tom",
//		Age:  18,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt(&test.User{
//		Name: "michael",
//		Age:  8,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2(&test.User{
//		Name: "jack",
//		Age:  2,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2(&test.User{
//		Name: "jack",
//		Age:  3,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2(&test.User{
//		Name: "jack",
//		Age:  4,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2(&test.User{
//		Name: "default",
//		Age:  8,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2(&test.User{
//		Name: "default",
//		Age:  9,
//		From: "usa",
//	})
//	require.NoError(t, err)
//
//	users, err := stmtMapper.TestQueryStmt("tom", 18)
//	require.NoError(t, err)
//	require.True(t, len(users) > 0)
//	t.Log(users[0].Name, users[0].Age)
//
//	users, err = stmtMapper.TestQueryStmt("michael", 8)
//	require.NoError(t, err)
//	require.True(t, len(users) > 0)
//	t.Log(users[0].Name, users[0].Age)
//}

//func TestStmtTx(t *testing.T) {
//	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
//	err := engine.Init(NewBundle("test/sql"))
//	require.NoError(t, err)
//	err = engine.master.Ping()
//	require.NoError(t, err)
//	engine.SetLogLevel(DebugLevel)
//	defer func() {
//		engine.Close()
//	}()
//
//	stmtMapper := new(StmtMapper)
//	err = engine.BindMapper(stmtMapper)
//	require.NoError(t, err)
//
//	tx, err := engine.Master().Begin()
//	require.NoError(t, err)
//	defer func() {
//		if err != nil {
//			err = tx.Rollback()
//			require.NoError(t, err)
//		} else {
//			err = tx.Commit()
//			require.NoError(t, err)
//		}
//	}()
//
//	err = stmtMapper.TestInsertStmtTx(tx, &test.User{
//		Name: "tom_tx",
//		Age:  18,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmtTx(tx, &test.User{
//		Name: "michael_tx",
//		Age:  8,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2Tx(tx, &test.User{
//		Name: "default_tx",
//		Age:  8,
//	})
//	require.NoError(t, err)
//
//	err = stmtMapper.TestInsertStmt2Tx(tx, &test.User{
//		Name: "default_tx",
//		Age:  9,
//	})
//	require.NoError(t, err)
//
//	users, err := stmtMapper.TestQueryStmtTx(tx, "tom_tx", 18)
//	require.NoError(t, err)
//	require.True(t, len(users) > 0, len(users))
//	t.Log(users[0].Name, users[0].Age)
//
//	users, err = stmtMapper.TestQueryStmtTx(tx, "default_tx", 8)
//	require.NoError(t, err)
//	require.True(t, len(users) > 0, len(users))
//	t.Log(users[0].Name, users[0].Age)
//}

//func TestStringArray(t *testing.T) {
//	engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
//	err := engine.Init(NewBundle("test"))
//	err = engine.master.Ping()
//	require.NoError(t, err)
//	engine.SetLogLevel(DebugLevel)
//	defer func() {
//		engine.Close()
//	}()
//	m := new(test.StmtMapper)
//	err = engine.BindMapper(m)
//	require.NoError(t, err)
//
//	tags := pgtype.TextArray{}
//	err = tags.Set([]string{"a", "b"})
//	require.NoError(t, err)
//	err = m.InsertStringArray(&test.User{
//		Id:   0,
//		Name: "tags",
//		Tags: tags,
//	})
//	require.NoError(t, err)
//
//	//r, err := m.GetStringArray("tags")
//	//require.NoError(t, err)
//	//for _, v := range r.Tags.Elements {
//	//	fmt.Println(v.String)
//	//}
//}

func testSelectInsert(t *testing.T, _testMapper *test.TestMapper) {
	id, err := _testMapper.SelectInsert(test.Entity{
		Int8:                     1,
		BigInt:                   2,
		Int:                      3,
		Decimal:                  decimal.NewFromFloat(3.14),
		Numeric:                  decimal.NewFromFloat(3.14156),
		Real:                     4,
		DoublePrecision:          5.1,
		SmallSerial:              6,
		Serial:                   7,
		BigSerial:                8,
		Money:                    "1.1",
		Char:                     "hello",
		Text:                     "world",
		TimestampWithoutTimeZone: time.Now(),
		TimestampWithTimeZone:    time.Now(),
		Date:                     time.Now(),
		TimeWithoutTimeZone:      time.Now(),
		TimeWithTimeZone:         time.Now(),
		Interval:                 100 * time.Second,
		Boolean:                  true,
	})
	
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertPointer(t *testing.T, _testMapper *test.TestMapper) {
	dec := decimal.NewFromFloat(3.14)
	now := time.Now()
	interval := 100 * time.Second
	
	id, err := _testMapper.SelectInsertPointer(&test.EntityPointer{
		Int8:                     pointer.ToInt8(1),
		BigInt:                   pointer.ToInt64(2),
		Int:                      pointer.ToInt(3),
		Decimal:                  &dec,
		Numeric:                  &dec,
		Real:                     pointer.ToFloat64(4),
		DoublePrecision:          pointer.ToFloat64(5.1),
		SmallSerial:              pointer.ToInt(6),
		Serial:                   pointer.ToInt(7),
		BigSerial:                pointer.ToInt(8),
		Money:                    pointer.ToString("1.1"),
		Char:                     pointer.ToString("hello"),
		Text:                     pointer.ToString("world"),
		TimestampWithoutTimeZone: &now,
		TimestampWithTimeZone:    &now,
		Date:                     &now,
		TimeWithoutTimeZone:      &now,
		TimeWithTimeZone:         &now,
		Interval:                 &interval,
		Boolean:                  pointer.ToBool(true),
	})
	
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachSlice(t *testing.T, _testMapper *test.TestMapper) {
	id, err := _testMapper.SelectInsertForeachSlice(test.Entity{
		Int8: 1,
	}, []string{"tom", "alice"})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachSlicePointer(t *testing.T, _testMapper *test.TestMapper) {
	enums := [][]*string{
		{pointer.ToString("tom1"), pointer.ToString("alice1")},
		{pointer.ToString("tom2"), pointer.ToString("alice2")},
	}
	id, err := _testMapper.SelectInsertForeachSlicePointer(&test.EntityPointer{
		Int8: pointer.ToInt8(1),
	}, &enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachMap(t *testing.T, _testMapper *test.TestMapper) {
	enums := map[string][]string{
		"first":  {"f1", "f2"},
		"second": {"fs", "s2"},
	}
	id, err := _testMapper.SelectInsertForeachMap(test.Entity{
		Int8: 1,
	}, enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachMapPointer(t *testing.T, _testMapper *test.TestMapper) {
	enums := map[string][]*string{
		"first":  {pointer.ToString("f1"), pointer.ToString("f2")},
		"second": {pointer.ToString("fs"), pointer.ToString("s2")},
	}
	id, err := _testMapper.SelectInsertForeachMapPointer(&test.EntityPointer{
		Int8: pointer.ToInt8(1),
	}, &enums)
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachStruct(t *testing.T, _testMapper *test.TestMapper) {
	id, err := _testMapper.SelectInsertForeachStruct(test.Entity{
		Int8: 1,
		Char: "Hello",
	})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

func testSelectInsertForeachStructPointer(t *testing.T, _testMapper *test.TestMapper) {
	id, err := _testMapper.SelectInsertForeachStructPointer(&test.EntityPointer{
		Char: pointer.ToString("Hello"),
	})
	require.NoError(t, err)
	if id <= 0 {
		require.Error(t, fmt.Errorf("returning id should greater 0"))
	}
}

//func testSelectInsertContextTx(t *testing.T, engine *Engine, _testMapper *test.TestMapper) {
//	ctx := context.WithValue(context.Background(), "name", "gobatis")
//	tx, err := engine.Master().Begin()
//	require.NoError(t, err)
//	t.Log(_testMapper.SelectInsertContextTx == nil)
//	id, err := _testMapper.SelectInsertContextTx(ctx, tx.Tx(), test.Entity{
//		Char: "hello",
//	})
//	require.NoError(t, err)
//	err = tx.Commit()
//	require.NoError(t, err)
//	if id <= 0 {
//		require.Error(t, fmt.Errorf("returning id should greater 0"))
//	}
//}

func testInsert(t *testing.T, _testMapper *test.TestMapper) {
	rows, err := _testMapper.Insert("Insert", "red", "yellow", "blue")
	require.NoError(t, err)
	if rows != 1 {
		require.Error(t, fmt.Errorf("rows expected 1"))
	}
}

func testSelectRow(t *testing.T, _testMapper *test.TestMapper) {
	tChar, tText, err := _testMapper.SelectRow(950)
	require.NoError(t, err)
	require.Equal(t, tChar, "hello")
	require.Equal(t, tText, "world")
}

func testSelectRowPointer(t *testing.T, _testMapper *test.TestMapper) {
	tChar, tText, err := _testMapper.SelectRowPointer(pointer.ToInt(950))
	require.NoError(t, err)
	require.Equal(t, *tChar, "hello")
	require.Equal(t, *tText, "world")
}

func testSelectRows(t *testing.T, _testMapper *test.TestMapper) {
	tChar, tText, err := _testMapper.SelectRows(363, 364)
	require.NoError(t, err)
	for _, v := range tChar {
		//require.Equal(t, v, "hello")
		t.Log(v)
	}
	for _, v := range tText {
		//require.Equal(t, v, "world")
		t.Logf("scanner: %+v", v)
	}
}

func testSelectRowsPointer(t *testing.T, _testMapper *test.TestMapper) {
	tChar, tText, err := _testMapper.SelectRowsPointer(pointer.ToInt(47), pointer.ToInt(50))
	require.NoError(t, err)
	for _, v := range tChar {
		require.Equal(t, *v, "hello")
	}
	for _, v := range tText {
		require.Equal(t, *v, "world")
	}
}

func testSelectStruct(t *testing.T, _testMapper *test.TestMapper) {
	item, err := _testMapper.SelectStruct(47)
	_ = item
	require.NoError(t, err)
	//d, err := json.MarshalIndent(item, "", "\t")
	require.NoError(t, err)
	//fmt.Println(string(d))
	
	item2, err := _testMapper.SelectStructPointer(47)
	_ = item2
	require.NoError(t, err)
	//d, err = json.MarshalIndent(item2, "", "\t")
	require.NoError(t, err)
	//fmt.Println(string(d))
}

func testSelectStructs(t *testing.T, _testMapper *test.TestMapper) {
	item, err := _testMapper.SelectStructs(47)
	_ = item
	require.NoError(t, err)
	//d, err := json.MarshalIndent(item, "", "\t")
	require.NoError(t, err)
	//fmt.Println(string(d))
	
	item2, err := _testMapper.SelectStructsPointer(47)
	_ = item2
	require.NoError(t, err)
	//d, err = json.MarshalIndent(item2, "", "\t")
	require.NoError(t, err)
	//fmt.Println(string(d))
}
