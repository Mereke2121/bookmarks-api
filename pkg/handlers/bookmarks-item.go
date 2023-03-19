package handlers

import (
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	items, err := h.service.GetAllItems(userId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var item models.Item
	err = c.BindJSON(&item)
	if err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}
	item.UserId = userId

	err = h.service.AddItem(&item)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func (h *Handler) RemoveItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteItem(id, userId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func (h *Handler) GetItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.service.GetItemById(id, userId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
