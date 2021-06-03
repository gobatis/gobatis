package entity

import "time"

type User struct {
	Id        int        `sql:"id"`
	Name      string     `sql:"name"`
	Age       int        `sql:"age"`
	From      string     `sql:"from"`
	Vip       bool       `sql:"vip"`
	CreatedAt *time.Time `sql:"created_at"`
}
