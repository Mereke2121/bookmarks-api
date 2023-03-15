package services

import (
	"fmt"
	"github.com/bookmarks-api/models"
	"github.com/dgrijalva/jwt-go"
	"strconv"
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

	token, err := CreateToken(strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)
		return userId, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
