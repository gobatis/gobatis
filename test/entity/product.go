package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id        int64           `sql:"id"`
	Name      string          `sql:"name"`
	Age       int             `sql:"age"`
	Height    float32         `sql:"height"`
	Price     decimal.Decimal `sql:"price"`
	Price2    *decimal.Decimal `sql:"price"`
	CreatedAt *time.Time      `sql:"created_at"`
}
