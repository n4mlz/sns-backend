package interfaces

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
	api := h.Router.Group("/api")
	{
		setSettingsRoutesFrom(api)
		setUsersRoutesFrom(api)
		setFollowsRoutesFrom(api)
	}
}
