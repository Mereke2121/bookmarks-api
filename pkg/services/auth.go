package services

import (
	"github.com/bookmarks-api/models"
	"github.com/bookmarks-api/pkg/repository"
	"github.com/pkg/errors"
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
	return id, err
}

func (s *AuthService) Authorize(authData *models.Authorization) (string, error) {
	authData.Password = generatePassword(authData.Password)
	id, err := s.repo.GetUserId(authData.Email, authData.Password)
	if err != nil {
		return "", errors.Wrap(err, "get user id from repository")
	}

	token, err := CreateToken(strconv.Itoa(id))
	if err != nil {
		return "", errors.Wrap(err, "create token by user id")
	}
	return token, nil
}

func (s *AuthService) ParseToken(token string) (string, error) {
	return VerifyToken(token)
}
