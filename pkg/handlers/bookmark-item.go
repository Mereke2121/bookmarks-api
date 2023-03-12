package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	items, err := h.service.GetAllItems()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {}

func (h *Handler) RemoveItem(c *gin.Context) {}
