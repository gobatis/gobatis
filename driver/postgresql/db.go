package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/crypto/ssh"
	"net"
	"regexp"
)

const PGX = "pgx"

func InitDB(dsn string, sshClient *ssh.Client) (db *sql.DB, err error) {
	if sshClient == nil {
		db, err = sql.Open(PGX, dsn)
		if err != nil {
			err = fmt.Errorf("postgresql connnet error: %s", err)
			return
		}
		return
	} else {
		var config *pgx.ConnConfig

		config, err = pgx.ParseConfig(dsn)
		if err != nil {
			return nil, err
		}

		config.PreferSimpleProtocol = true
		result := regexp.MustCompile("(time_zone|TimeZone)=(.*?)($|&| )").FindStringSubmatch(dsn)
		if len(result) > 2 {
			config.RuntimeParams["timezone"] = result[2]
		}

		config.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshClient.Dial(network, addr)
		}

		db = stdlib.OpenDB(*config)
		return db, nil
	}

}
