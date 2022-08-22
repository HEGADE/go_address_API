package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetConnection(ctx context.Context, username, password, host, port string) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/address?statement_cache_mode=describe&sslmode=disable", username, password, host, port)
	timeoutctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1
	conn, err := pgxpool.ConnectConfig(timeoutctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
