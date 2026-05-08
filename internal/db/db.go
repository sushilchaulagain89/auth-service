package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	dsn := "postgres://app:app123@localhost:5432/authdb"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	Pool = pool
	fmt.Println("Connected to PostgreSQL")
}