package repository

import (
	"fmt"
	"github.com/bookmarks-api/config"
	"github.com/jmoiron/sqlx"
	"os"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s host=%s port=%s password=%s",
		cfg.DB.User,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
		cfg.DB.Host,
		cfg.DB.Port,
		os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
