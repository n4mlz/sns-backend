package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

func (h *Handler) SetupRoutes() {
	firebase := h.Router.Group("/firebase")
	firebase.Use(authMiddleware())
	{
		firebase.GET("", func(ctx *gin.Context) {
			// some action
		})
	}
}
