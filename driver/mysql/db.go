package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"net"
)

const MySQL = "mysql"

func InitDB(dsn string, sshClient *ssh.Client) (db *sql.DB, err error) {
	if sshClient != nil {
		mysql.RegisterDialContext("tcp", func(ctx context.Context, addr string) (net.Conn, error) {
			return sshClient.Dial("tcp", addr)
		})
	}

	db, err = sql.Open(MySQL, dsn)
	if err != nil {
		err = fmt.Errorf("mysql connnet error: %s", err)
		return
	}
	return
}
