package database

import (
	"fmt"
	"os"
	"ultradex/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgres initializes and returns a PostgreSQL connection
func NewPostgres(cfg config.FlyConfig) (*gorm.DB, error) {
	dsn := cfg.Env.DatabaseURL
	if os.Getenv("ENV") == "prod" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
