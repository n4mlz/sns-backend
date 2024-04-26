package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setSettingsRoutesFrom(r *gin.RouterGroup) {
	settings := r.Group("/settings")
	settings.Use(authMiddleware())
	{
		profile := settings.Group("/profile")
		{
			profile.GET("", usecases.GetOwnProfile)
			profile.PUT("", usecases.SaveProfile)
			profile.PUT("/icon", usecases.SaveIcon)
			profile.PUT("/bgimage", usecases.SaveBgImage)
		}
	}
}
