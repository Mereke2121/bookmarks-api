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

func (r *Repository) AddItem(item *models.Item) error {
	query := `insert into bookmarks_items (url, title) values($1, $2)`
	_, err := r.db.Exec(query, item.Url, item.Title)
	return err
}

func (r *Repository) DeleteItem(id int) error {
	query := `delete from bookmarks_items where id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) AddUser(user *models.User) (int, error) {
	query := `insert into users (username, email, password_hash) values($1, $2, $3) returning id`

	var id int
	err := r.db.QueryRow(query, user.UserName, user.Email, user.Password).Scan(&id)

	return id, err
}
