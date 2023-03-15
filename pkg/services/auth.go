package services

import (
	"github.com/bookmarks-api/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

const (
	signingKey = "39fije9wjfe90"
)

func (s *Service) AddUser(user *models.User) (int, error) {
	user.Password = generatePassword(user.Password)
	id, err := s.repo.AddUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Service) Authorize(authData *models.Authorization) (string, error) {
	authData.Password = generatePassword(authData.Password)
	id, err := s.repo.GetUserId(authData.Email, authData.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(signingKey))
}
