package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gozelle/spew"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestNamer(t *testing.T) {
	n := &Namer{}
	a := n.ReservedName(`"insert"`)
	require.Equal(t, a, `"insert"`)

	b := n.TableName("public.products")
	require.Equal(t, b, `"public"."products"`)
}

func TestLink(t *testing.T) {
	conn, err := pgx.Connect(context.Background(), "postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var count int64
	err = conn.QueryRow(context.Background(), `select count(1) from products where id in ?`, []int64{1}).Scan(&count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	spew.Json(count)
}

func TestDB(t *testing.T) {

	d := Open("postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable")
	db, err := d.DB()
	require.NoError(t, err)

	var count int64
	err = db.QueryRow(`select count(1) from products where id in ($1,$2)`, 1, '2').Scan(&count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	spew.Json(count)
}
