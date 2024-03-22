package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setFollowsRoutesFrom(r *gin.RouterGroup) {
	follows := r.Group("/follows")
	follows.Use(authMiddleware())
	{
		follows.PUT("/follow", models.FollowUser)
		follows.PUT("/unfollow", models.UnfollowUser)
		follows.GET("/followed", models.Followed)
		follows.PUT("/reject", models.RejectUser)
	}
}
