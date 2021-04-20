package types

import (
	"database/sql/driver"
)

type Engine struct {
	Test string `json:"test"`
}

func (e Engine) Close() error {
	panic("implement me")
}

func (e Engine) NumInput() int {
	panic("implement me")
}

func (e Engine) Exec(args []driver.Value) (driver.Result, error) {
	panic("implement me")
}

func (e Engine) Query(args []driver.Value) (driver.Rows, error) {
	panic("implement me")
}

func (e Engine) Open(name string) (driver.Conn, error) {
	panic("implement me")
}
