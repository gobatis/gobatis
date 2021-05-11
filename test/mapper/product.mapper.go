package mapper

import "github.com/gobatis/gobatis/test/entity"

type ProductMapper struct {
	CreateProduct   func(product *entity.Product) (rows int64, err error)
	GetProductById  func(id int64) (product *entity.Product, err error)
	GetProductsById func(id int64) (product []*entity.Product, err error)
}
