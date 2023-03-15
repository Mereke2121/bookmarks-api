package handlers

import (
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleError(c *gin.Context, status int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(status, models.ErrorResponse{Message: message})
}
