package handlers

import (
	"errors"
	"github.com/bookmarks-api/models"
	"github.com/gin-gonic/gin"
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
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.service.AddItem(&item)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (h *Handler) RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("parse int for id"))
		return
	}

	err = h.service.DeleteItem(id)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
