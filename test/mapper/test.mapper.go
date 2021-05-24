package mapper

import (
	"context"
	"database/sql"
	"github.com/gobatis/gobatis/test/entity"
)

type TestMapper struct {
	SelectInsert                     func(entity entity.TestEntity) (id int, err error)
	SelectInsertPointer              func(entity *entity.TestEntityPointer) (id int8, err error)
	SelectInsertForeachSlice         func(testEntity entity.TestEntity, enums []string) (id int16, err error)
	SelectInsertForeachSlicePointer  func(testEntity *entity.TestEntityPointer, enums *[][]*string) (id int32, err error)
	SelectInsertForeachMap           func(testEntity entity.TestEntity, enums map[string][]string) (id uint, err error)
	SelectInsertForeachMapPointer    func(testEntity *entity.TestEntityPointer, enums *map[string][]*string) (id uint16, err error)
	SelectInsertForeachStruct        func(testEntity entity.TestEntity) (id uint32, err error)
	SelectInsertForeachStructPointer func(testEntity *entity.TestEntityPointer) (id uint64, err error)
	SelectInsertContextTx            func(ctx context.Context, tx *sql.Tx, testEntity entity.TestEntity) (int int, err error)
	Insert                           func(name string, tags ...string) (rows int64, err error)
	//Update                           func(id int64, entity entity.TestEntity) (rows int64, err error)
	//Delete                           func(id int64) (rows int64, err error)
	SelectRow         func(id int) (t_char, t_text string, err error)
	SelectRowPointer  func(id *int) (t_char, t_text *string, err error)
	SelectRows        func(start, end int) (t_char, t_text []string, err error)
	SelectRowsPointer func(start, end *int) (t_char, t_text []*string, err error)
	//SelectStruct                     func() (tInt int, tChar string, tDecimal decimal.Decimal, tTime time.Time, tInterval time.Duration, err error)
	//SelectStructPointers             func() (tInt *int, tChar *string, tDecimal *decimal.Decimal, tTime *time.Time, tInterval *time.Duration, err error)
}
