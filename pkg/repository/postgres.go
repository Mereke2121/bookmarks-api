package repository

import (
	"fmt"
	"github.com/bookmarks-api/config"
	"github.com/jmoiron/sqlx"
	"os"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s",
		cfg.DB.User,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
		os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
