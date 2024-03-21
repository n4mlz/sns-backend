package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setSettingsRoutesFrom(r *gin.RouterGroup) {
	settings := r.Group("/settings")
	settings.Use(authMiddleware())
	{
		settings.PUT("", models.SaveProfile)
	}
}
