package gobatis

import "database/sql"

type Session struct {
	conn *sql.Conn
}
