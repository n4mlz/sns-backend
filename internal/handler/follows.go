package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setFollowsRoutesFrom(r *gin.RouterGroup) {
	users := r.Group("/follows")
	users.Use(authMiddleware())
	{
		users.PUT("/follow", models.FollowUser)
		users.PUT("/unfollow", models.UnfollowUser)
		users.GET("/followed", models.Followed)
	}
}
