package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setFollowsRoutesFrom(r *gin.RouterGroup) {
	follows := r.Group("/follows")
	follows.Use(authMiddleware())
	{
		follows.PUT("/follow", usecases.FollowUser)
		follows.PUT("/unfollow", usecases.UnfollowUser)
		follows.GET("/requested", usecases.RequestedUsers)
		follows.PUT("/reject", usecases.RejectUser)
	}
}
