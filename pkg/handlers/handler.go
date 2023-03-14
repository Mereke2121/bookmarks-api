package handlers

import (
	"github.com/bookmarks-api/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := gin.New()

	mux.POST("/sign-up", h.SignUp)

	mux.GET("/", h.GetAllItems)
	mux.POST("/", h.AddItem)
	mux.DELETE("/:id", h.RemoveItem)

	return mux.Handler()
}
