package handlers

import (
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	id, err := h.service.AddUser(&user)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{
		Id: id,
	})
}
