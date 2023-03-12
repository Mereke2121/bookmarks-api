package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	query := `select id, url, title from bookmarks_items`

	err := r.db.Select(&items, query)
	if err != nil {
		return items, err
	}
	return items, nil
}
