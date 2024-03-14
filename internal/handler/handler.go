package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/validation"
)

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

func (h *Handler) SetupRoutes(firebaseApp *validation.FirebaseApp) {
	firebase := h.Router.Group("/firebase")
	firebase.Use(authMiddleware(*firebaseApp))
	{
		firebase.GET("", func(ctx *gin.Context) {
			// some action
		})
	}
}
