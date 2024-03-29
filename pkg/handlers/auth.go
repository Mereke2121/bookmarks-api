package handlers

import (
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.AddUser(user)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "service failure")
		return
	}
	c.JSON(http.StatusOK, models.AddUserResponse{
		Id: id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var authData models.Authorization
	if err := c.BindJSON(&authData); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorize(&authData)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.AuthorizationResponse{
		Token: token,
	})
}
