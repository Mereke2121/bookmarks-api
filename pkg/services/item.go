package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
)

type ItemsService struct {
	repo *repository.Repository
}

func NewItemsService(repo *repository.Repository) Items {
	return &ItemsService{repo: repo}
}

func (s *ItemsService) GetAllItems(userId int) ([]models.Item, error) {
	items, err := s.repo.GetAllItems(userId)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *ItemsService) AddItem(item *models.Item) error {
	err := s.repo.AddItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *ItemsService) DeleteItem(id, userId int) error {
	err := s.repo.DeleteItem(id, userId)
	return err
}

func (s *ItemsService) GetItemById(id, userId int) (models.Item, error) {
	item, err := s.repo.GetItemById(id, userId)
	return item, err
}
