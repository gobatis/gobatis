package batis

import "github.com/gobatis/gobatis/executor"

func Param(name string, value any) executor.NameValue {
	return executor.NameValue{Name: name, Value: value}
}
