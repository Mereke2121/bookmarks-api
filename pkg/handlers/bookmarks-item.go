package handlers

import (
	"errors"
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	userId, _ := getUserId(c)

	items, err := h.service.GetAllItems(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	userId, _ := getUserId(c)

	var item models.Item
	err := c.BindJSON(&item)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	item.UserId = userId

	err = h.service.AddItem(&item)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func (h *Handler) RemoveItem(c *gin.Context) {
	userId, _ := getUserId(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteItem(id, userId)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
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
