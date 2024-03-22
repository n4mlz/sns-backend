package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setUsersRoutesFrom(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.Use(authMiddleware())
	{
		user := users.Group("/:userName")
		{
			user.GET("", models.User)
			user.GET("/mutuals", models.MutualFollow)
		}
	}
}
