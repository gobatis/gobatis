package entity

import "time"

type User struct {
	Id        int        `db:"id"`
	Name      string     `db:"name"`
	Age       int        `db:"age"`
	From      string     `db:"from"`
	Vip       bool       `db:"vip"`
	CreatedAt *time.Time `db:"created_at"`
}
