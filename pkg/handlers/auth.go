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
	c.JSON(http.StatusOK, models.AddUserResponse{
		Id: id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var authData models.Authorization
	if err := c.BindJSON(&authData); err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, err := h.service.Authorize(&authData)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, models.AuthorizationResponse{
		Token: token,
	})
}
