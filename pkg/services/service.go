package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
)

type Authorization interface {
	AddUser(user *models.User) (int, error)
	Authorize(authData *models.Authorization) (string, error)
	ParseToken(token string) (string, error)
}

type Items interface {
	GetAllItems(userId int) ([]models.Item, error)
	AddItem(item *models.Item) error
	DeleteItem(id, userId int) error
}

type Service struct {
	Authorization
	Items
	//repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Items:         NewItemsService(repo),
	}
}
