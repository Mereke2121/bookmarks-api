package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
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
