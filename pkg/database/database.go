package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool(db_dsn string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), db_dsn)
	if err != nil {
		panic(err)
	}
	return pool
}

func CloseDB(conn *pgxpool.Pool) {
	conn.Close()
}
