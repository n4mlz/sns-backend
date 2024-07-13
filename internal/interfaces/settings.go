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
			profile.PUT("/userName", usecases.SaveUserName)
			profile.PUT("/icon", usecases.SaveIcon)
			profile.PUT("/bgImage", usecases.SaveBgImage)
		}

		account := settings.Group("/account")
		{
			account.DELETE("", usecases.DeleteUser)
		}
	}
}
