# 实体类

## 定义示例

``` go
type User struct {
    Name string         `sql:"name"`
    City *string        `sql:"city"`
    Area sql.NullString `sql:"area"`
}
```

## 基础数据类型

* sql.Scanner（最高优先级）;
* int, int8, int16, int32, int64;
* uint, uint8, uint16, uint32, uint64;
* float32, float64;
* string;
* bool;
* time.Time, time.Duration;
* decimal.Decimal;

