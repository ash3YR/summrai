package db

import (
	"context"
	"fmt"
	"log"

	"summrai-backend/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(cfg *config.Config) *pgxpool.Pool {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	dbpool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = dbpool.Ping(context.Background())

	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")

	return dbpool
}