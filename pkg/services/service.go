package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
	"log"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllItems() ([]models.Item, error) {
	items, err := s.repo.GetAllItems()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *Service) AddItem(item models.Item) error {
	err := s.repo.AddItem(item)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
