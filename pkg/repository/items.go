package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ItemsRepository struct {
	db *sqlx.DB
}

func NewItemsRepository(db *sqlx.DB) Items {
	return &ItemsRepository{db: db}
}

func (r *ItemsRepository) GetAllItems(userId int) ([]models.Item, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "start db")
	}

	query := `select id, url, title from bookmarks_items where user_id=$1`

	var items []models.Item
	err = r.db.Select(&items, query, userId)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "get items from db")
	}

	return items, err
}

func (r *ItemsRepository) AddItem(item *models.Item) error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "start db")
	}

	query := `insert into bookmarks_items (user_id, url, title) values($1, $2, $3)`

	_, err = r.db.Exec(query, item.UserId, item.Url, item.Title)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "insert items into the db")
	}

	return err
}

func (r *ItemsRepository) DeleteItem(id, userId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "start db")
	}

	query := `delete from bookmarks_items where id=$1 and user_id=$2`

	_, err = r.db.Exec(query, id, userId)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "delete items from db")
	}

	return err
}

func (r *ItemsRepository) GetItemById(id, userId int) (models.Item, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return models.Item{}, errors.Wrap(err, "start db")
	}

	query := `select id, url, title from bookmarks_items where id=$1 and user_id=$2`

	var item models.Item
	err = r.db.Get(&item, query, id, userId)
	if err != nil {
		tx.Rollback()
		return models.Item{}, errors.Wrap(err, "get items from db")
	}

	return item, err
}
