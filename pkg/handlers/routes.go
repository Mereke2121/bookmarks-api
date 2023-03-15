package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitRoutes() http.Handler {
	mux := gin.New()

	mux.POST("/sign-up", h.SignUp)
	mux.POST("/sign-in", h.SignIn)

	mux.GET("/", h.GetAllItems)
	mux.POST("/", h.AddItem)
	mux.DELETE("/:id", h.RemoveItem)

	return mux.Handler()
}
