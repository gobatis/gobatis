package batis

import (
	"github.com/gobatis/gobatis/executor"
)

func Select(data any, columns string) executor.Rows {
	return executor.NewSelectColumns(data, columns)
}

func Except(data any, columns string) executor.Rows {
	return executor.NewExceptColumns(data, columns)
}
