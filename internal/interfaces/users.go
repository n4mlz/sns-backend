package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setUsersRoutesFrom(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.Use(authMiddleware())
	{
		user := users.Group("/:userName")
		{
			user.GET("", usecases.User)
			user.GET("/mutuals", usecases.MutualFollow)
		}
	}
}
