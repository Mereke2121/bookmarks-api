package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return
	}

	// "Bearer <token>"
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		return
	}

	c.Set("userId", id)
}
