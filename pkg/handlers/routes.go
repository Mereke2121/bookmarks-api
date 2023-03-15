package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitRoutes() http.Handler {
	mux := gin.New()

	mux.POST("/sign-up", h.SignUp)
	mux.POST("/sign-in", h.SignIn)

	api := mux.Group("/api", h.UserIdentity)
	{
		api.GET("/", h.GetAllItems)
		api.POST("/", h.AddItem)
		api.DELETE("/:id", h.RemoveItem)
	}

	return mux.Handler()
}
