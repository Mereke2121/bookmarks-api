package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
	"strconv"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) Authorization {
	return &AuthService{repo: repo}
}

func (s *AuthService) AddUser(user *models.User) (int, error) {
	user.Password = generatePassword(user.Password)
	id, err := s.repo.AddUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthService) Authorize(authData *models.Authorization) (string, error) {
	authData.Password = generatePassword(authData.Password)
	id, err := s.repo.GetUserId(authData.Email, authData.Password)
	if err != nil {
		return "", err
	}

	token, err := CreateToken(strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) ParseToken(token string) (string, error) {
	return VerifyToken(token)
}
