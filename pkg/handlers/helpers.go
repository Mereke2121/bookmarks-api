package handlers

import (
	"errors"
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleError(c *gin.Context, status int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(status, models.ErrorResponse{Message: message})
}

func getUserId(c *gin.Context) (int, error) {
	// TODO: вынести в константы такие значения как userId
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("convert id from interface to int")
	}

	return idInt, nil
}
