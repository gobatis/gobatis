package tests

import (
	"context"
	"errors"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestPagingQuery(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	m := getProductsMap()
	c := map[int]string{
		0: Chair,
		1: BluetoothHeadphones,
		2: Smartwatch,
		3: Smartphone,
	}

	for i := 0; i <= 4; i++ {
		var products []*Product
		var count int64

		q := batis.PagingQuery{
			Select: "*",
			Count:  "1",
			From:   "products",
			Where:  "price <= #{price}",
			Order:  "price asc",
			Page:   int64(i),
			Limit:  1,
			Params: map[string]any{
				"price": decimal.NewFromInt(900),
			},
			Scan: func(s batis.PagingScanner) error {
				return s.Scan(&count, &products)
			},
		}

		err = db.Debug().PagingQuery(q).Error
		require.NoError(t, err)
		require.Equal(t, int64(4), count)

		if i < 4 {
			require.Equal(t, 1, len(products))
			for _, v := range products {
				vv := m[v.ProductName]
				vv.Id = v.Id
				vv.AddedDateTime = v.AddedDateTime
				require.Equal(t, c[i], v.ProductName)
				compareProduct(t, v, vv)
			}
		} else {
			require.Equal(t, 0, len(products))
		}
	}
}

func TestPagingQueryAffect(t *testing.T) {
	var products []*Product
	var count int64
	err := db.Affect(0).PagingQuery(batis.PagingQuery{
		Select: "*",
		Count:  "1",
		From:   "products",
		Where:  "",
		Order:  "id desc",
		Page:   0,
		Limit:  2,
		Params: nil,
		Scan: func(scanner batis.PagingScanner) error {
			return scanner.Scan(&count, &products)
		},
	}).Error

	require.True(t, errors.Is(err, batis.ErrNotSupportAffectConstraint))
}

func TestPagingQueryRowsAffected(t *testing.T) {
	var products []*Product
	var count int64
	_, err := db.PagingQuery(batis.PagingQuery{
		Select: "*",
		Count:  "1",
		From:   "products",
		Where:  "",
		Order:  "id desc",
		Page:   0,
		Limit:  2,
		Params: nil,
		Scan: func(scanner batis.PagingScanner) error {
			return scanner.Scan(&count, &products)
		},
	}).RowsAffected()

	require.True(t, errors.Is(err, batis.ErrNoSQLResultExists))
}

func TestPagingQueryContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).PagingQuery(batis.PagingQuery{
		Select: "*",
		Count:  "1",
		From:   "products",
		Where:  "",
		Order:  "id desc",
		Page:   0,
		Limit:  2,
		Params: nil,
		Scan: func(scanner batis.PagingScanner) error {
			return scanner.Scan(nil, nil)
		},
	}).Error)
}

func TestPagingTraceId(t *testing.T) {
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").PagingQuery(batis.PagingQuery{
			Select: "*",
			Count:  "1",
			From:   "products2",
			Where:  "",
			Order:  "id desc",
			Page:   0,
			Limit:  2,
			Params: nil,
			Scan: func(scanner batis.PagingScanner) error {
				return scanner.Scan(nil, nil)
			},
		})
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).PagingQuery(batis.PagingQuery{
			Select: "*",
			Count:  "1",
			From:   "products2",
			Where:  "",
			Order:  "id desc",
			Page:   0,
			Limit:  2,
			Params: nil,
			Scan: func(scanner batis.PagingScanner) error {
				return scanner.Scan(nil, nil)
			},
		})
		w.expectTraceId(t, "ctx")
	}
}

func TestPagingQueryColumnTag(t *testing.T) {

	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()

	m := getProductsMap()
	p := m[TV]
	i := &ProductJ{
		JId:              p.Id,
		JProductName:     p.ProductName,
		JDescription:     p.Description,
		JPrice:           p.Price,
		JWeight:          p.Weight,
		JStockQuantity:   p.StockQuantity,
		JIsAvailable:     p.IsAvailable,
		JManufactureDate: p.ManufactureDate,
		JAddedDateTime:   p.AddedDateTime,
	}

	s := db.Session(&batis.Session{ColumnTag: "json"})

	var n []*ProductJ
	var count int
	require.NoError(t, s.PagingQuery(batis.PagingQuery{
		Select: "*",
		Count:  "1",
		From:   "products",
		Where:  "product_name = #{name}",
		Order:  "id desc",
		Page:   0,
		Limit:  2,
		Params: map[string]any{
			"name": p.ProductName,
		},
		Scan: func(scanner batis.PagingScanner) error {
			return scanner.Scan(&count, &n)
		},
	}).Error)
	
	require.True(t, *n[0].JId > 0)
	n[0].JId = nil

	compareProductJ(t, i, n[0])
}
