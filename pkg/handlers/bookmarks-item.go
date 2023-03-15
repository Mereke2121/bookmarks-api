package handlers

import (
	"errors"
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, "get user id from header")
		return
	}

	items, err := h.service.GetAllItems(userId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "get all items by user id")
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, "get user id from header")
		return
	}

	var item models.Item
	err = c.BindJSON(&item)
	if err != nil {
		handleError(c, http.StatusBadRequest, "parse item from json to structure")
		return
	}
	item.UserId = userId

	err = h.service.AddItem(&item)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "add item by user id")
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func (h *Handler) RemoveItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, "get user id from header")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "parse int for user id")
		return
	}

	err = h.service.DeleteItem(id, userId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "delete item by user id")
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
