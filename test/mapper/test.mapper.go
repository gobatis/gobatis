package mapper

import (
	"github.com/gobatis/gobatis/test/entity"
)

type TestMapper struct {
	SelectInsert                    func(entity entity.TestEntity) (id int, err error)
	SelectInsertPointer             func(entity *entity.TestEntityPointer) (id int32, err error)
	SelectInsertForeachSlice        func(testEntity entity.TestEntity, enums []string) (id int16, err error)
	SelectInsertForeachSlicePointer func(testEntity *entity.TestEntityPointer, enums *[][]*string) (id uint16, err error)
	//SelectInsertForeachMap           func(testEntity entity.TestEntity, enums map[string]string) (rows int64, err error)
	//SelectInsertForeachMapPointer    func(testEntity *entity.TestEntity, enums *map[string]string) (rows uint, err error)
	//SelectInsertForeachStruct        func(testEntity entity.TestEntity, enums entity.TestEntity) (rows uint16, err error)
	//SelectInsertForeachStructPointer func(testEntity *entity.TestEntity, enums *entity.TestEntity) (rows uint64, err error)
	//SelectInsertContext              func(ctx context.Context, testEntity entity.TestEntity) (rows uint64, err error)
	//SelectInsertContextPointer       func(ctx *context.Context, testEntity *entity.TestEntity) (rows uint64, err error)
	//SelectInsertTx                   func(ctx *context.Context, tx *sql.Tx, testEntity *entity.TestEntity) (rows uint64, err error)
	//SelectInsertTxPointer            func(ctx *context.Context, tx *sql.Tx, testEntity *entity.TestEntity) (rows uint64, err error)
	//InsertPointer                    func(entity *entity.TestEntity) (rows int8, err error)
	//InsertForeachSlice               func(testEntity entity.TestEntity, enums []string) (rows int16, err error)
	//InsertForeachSlicePointer        func(testEntity entity.TestEntity, enums *[]string) (rows int32, err error)
	//InsertForeachMap                 func(testEntity entity.TestEntity, enums map[string]string) (rows int64, err error)
	//InsertForeachMapPointer          func(testEntity entity.TestEntity, enums *map[string]string) (rows uint, err error)
	//InsertForeachStruct              func(testEntity entity.TestEntity, enums entity.TestEntity) (rows uint16, err error)
	//InsertForeachStructPointer       func(testEntity entity.TestEntity, enums *entity.TestEntity) (rows uint64, err error)
	//Update                           func(id int64, entity entity.TestEntity) (rows int64, err error)
	//UpdatePointer                    func(id *int64, entity *entity.TestEntity) (rows int64, err error)
	//Delete                           func(id int64) (rows int64, err error)
	//SelectRow                        func(id int) (item entity.TestEntity, err error)
	//SelectRowPointer                 func(id *int) (item *entity.TestEntity, err error)
	//SelectRows                       func() (item []entity.TestEntity, err error)
	//SelectRowsPointer                func() (item []*entity.TestEntity, err error)
	//SelectFields                     func() (tInt int, tChar string, tDecimal decimal.Decimal, tTime time.Time, tInterval time.Duration, err error)
	//SelectFieldsPointers             func() (tInt *int, tChar *string, tDecimal *decimal.Decimal, tTime *time.Time, tInterval *time.Duration, err error)
}
