package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) InitRoutes() http.Handler {
	mux := gin.New()

	mux.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return mux.Handler()
}
