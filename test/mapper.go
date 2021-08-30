package test

import (
	"context"
	"database/sql"
)

type TestMapper struct {
	SelectInsert                     func(entity Entity) (id int, err error)
	SelectInsertPointer              func(entity *EntityPointer) (id int32, err error)
	SelectInsertForeachSlice         func(testEntity Entity, enums []string) (id int16, err error)
	SelectInsertForeachSlicePointer  func(testEntity *EntityPointer, enums *[][]*string) (id int32, err error)
	SelectInsertForeachMap           func(testEntity Entity, enums map[string][]string) (id uint, err error)
	SelectInsertForeachMapPointer    func(testEntity *EntityPointer, enums *map[string][]*string) (id uint16, err error)
	SelectInsertForeachStruct        func(testEntity Entity) (id uint32, err error)
	SelectInsertForeachStructPointer func(testEntity *EntityPointer) (id uint64, err error)
	SelectInsertContextTx            func(ctx context.Context, tx *sql.Tx, testEntity Entity) (int int, err error)
	Insert                           func(name string, tags ...string) (rows int64, err error)
	SelectRow                        func(id int) (t_char, t_text string, err error)
	SelectRowPointer                 func(id *int) (t_char, t_text *string, err error)
	SelectRows                       func(start, end int) (t_char []string, t_text []sql.NullString, err error)
	SelectRowsPointer                func(start, end *int) (t_char, t_text []*string, err error)
	SelectStruct                     func(id int) (entity Entity, err error)
	SelectStructPointer              func(id int) (entity *Entity, err error)
	SelectStructs                    func(id int) (entity []Entity, err error)
	SelectStructsPointer             func(id int) (entity []*Entity, err error)
}

type StmtMapper struct {
	TestInsertStmt    func(user *User) error
	TestInsertStmt2   func(user *User) error
	TestQueryStmt     func(name string, age int64) ([]*User, error)
	TestQueryStmt2    func(name string, age int64) ([]*User, error)
	InsertStringArray func(user *User) (err error)
	GetStringArray    func(name string) (user *User, err error)
}
