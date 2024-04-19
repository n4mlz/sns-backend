package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setSettingsRoutesFrom(r *gin.RouterGroup) {
	settings := r.Group("/settings")
	settings.Use(authMiddleware())
	{
		settings.PUT("", usecases.SaveProfile)
		settings.PUT("/icon", usecases.SaveIcon)
		settings.PUT("/bgimage", usecases.SaveBgImage)
	}
}
