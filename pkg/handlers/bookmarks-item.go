package handlers

import (
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	items, err := h.service.GetAllItems()
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	var item models.Item
	err := c.BindJSON(&item)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.service.AddItem(&item)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}

func (h *Handler) RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteItem(id)
	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, models.ItemResponse{Status: "ok"})
}
