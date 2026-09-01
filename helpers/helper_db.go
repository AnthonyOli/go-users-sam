package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func buildPostgresURL() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user, pass, host, port, name,
	)
}

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, buildPostgresURL())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
