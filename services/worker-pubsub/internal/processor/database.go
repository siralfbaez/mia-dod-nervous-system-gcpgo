package processor

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectAlloyDB initializes a high-performance connection pool
func ConnectAlloyDB(ctx context.Context) (*pgxpool.Pool, error) {
	// Staff Pattern: Use connection strings from environment variables
	// Format: postgres://user:password@localhost:5432/dbname
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// pgxpool is the standard for high-concurrency Go/Postgres apps
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	// Configure pool settings for "Pizza Hut" scale
	config.MaxConns = 20
	config.MinConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to AlloyDB: %v", err)
	}

	log.Println("✅ Connected to AlloyDB Metabolism")
	return pool, nil
}