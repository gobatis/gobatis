package batis

import "github.com/gobatis/gobatis/reflects"

func Select(data any, columns string) reflects.Rows {
	return reflects.NewSelectColumns(data, columns)
}

func Except(data any, columns string) reflects.Rows {
	return reflects.NewExceptColumns(data, columns)
}
