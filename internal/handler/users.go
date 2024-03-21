package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setUsersRoutesFrom(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.Use(authMiddleware())
	{
		users.PUT("/follow", models.FollowUser)
		users.PUT("/unfollow", models.UnfollowUser)

		user := users.Group("/:userName")
		{
			user.GET("", models.GetUser)
			user.GET("/mutuals", models.GetMutualFollow)
		}
	}
}
