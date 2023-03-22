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
	if len(headerParts) != 2 || headerParts[0] != "Bearer" || headerParts[1] == "" {
		handleError(c, http.StatusUnauthorized, "auth token in wrong format")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Set("userId", id)
}
