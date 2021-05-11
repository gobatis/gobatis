package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id        int64           `sql:"id"`
	Name      string          `sql:"name"`
	Width     int             `sql:"width"`
	Height    float32         `sql:"height"`
	Price     decimal.Decimal `sql:"price"`
	CreatedAt *time.Time      `sql:"created_at"`
}
