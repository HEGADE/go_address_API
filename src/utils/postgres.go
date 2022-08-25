package utils

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// old connection using pgx
// func GetConnection(ctx context.Context, username, password, host, port string) *pgxpool.Pool {
// 	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/address?statement_cache_mode=describe&sslmode=disable", username, password, host, port)
// 	timeoutctx, cancel := context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()
// 	poolConfig, err := pgxpool.ParseConfig(dsn)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	poolConfig.MaxConns = 10
// 	poolConfig.MinConns = 1
// 	conn, err := pgxpool.ConnectConfig(timeoutctx, poolConfig)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	return conn
// }

type pgConfig struct {
	user     string
	host     string
	port     int
	password string
	database string
}

func GormConnection(ctx context.Context) (*gorm.DB, error) {
	config := pgConfig{
		user:     "postgres",
		host:     "localhost",
		port:     5432,
		password: "1234",
		database: "address",
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.host, config.user, config.password, config.database, config.port,
	)
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	sqlDB, err := gormDB.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)
	return gormDB, err
}
