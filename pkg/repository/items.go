package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
)

type ItemsRepository struct {
	db *sqlx.DB
}

func NewItemsRepository(db *sqlx.DB) Items {
	return &ItemsRepository{db: db}
}

func (r *ItemsRepository) GetAllItems(userId int) ([]models.Item, error) {
	var items []models.Item
	query := `select id, url, title from bookmarks_items where user_id=$1`

	err := r.db.Select(&items, query, userId)
	return items, err
}

func (r *ItemsRepository) AddItem(item *models.Item) error {
	query := `insert into bookmarks_items (user_id, url, title) values($1, $2, $3)`
	_, err := r.db.Exec(query, item.UserId, item.Url, item.Title)
	return err
}

func (r *ItemsRepository) DeleteItem(id, userId int) error {
	query := `delete from bookmarks_items where id=$1 and user_id=$2`
	_, err := r.db.Exec(query, id, userId)
	return err
}

func (r *ItemsRepository) GetItemById(id, userId int) (models.Item, error) {
	query := `select id, url, title from bookmarks_items where id=$1 and user_id=$2`

	var item models.Item
	err := r.db.Get(&item, query, id, userId)
	return item, err
}
