package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	AddUser(user *models.User) (int, error)
	GetUserId(email, password string) (int, error)
}

type Items interface {
	GetAllItems(userId int) ([]models.Item, error)
	AddItem(item *models.Item) error
	DeleteItem(id, userId int) error
}

type Repository struct {
	Authorization
	Items
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Items:         NewItemsRepository(db),
	}
}
