package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		handleError(c, http.StatusUnauthorized, "get authorization token from header")
		return
	}

	// "Bearer <token>"
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		handleError(c, http.StatusBadRequest, "auth token in wrong format")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		handleError(c, http.StatusBadRequest, "parse jwt token")
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		handleError(c, http.StatusBadRequest, "parse int for user id")
		return
	}

	c.Set("userId", id)
}
