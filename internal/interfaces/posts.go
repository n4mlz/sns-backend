package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setPostsRoutesFrom(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	posts.Use(authMiddleware())
	{
		posts.POST("", usecases.CreatePost)
		posts.PUT("/like", usecases.LikePost)
		posts.PUT("/unlike", usecases.UnlikePost)

		post := posts.Group("/:postId")
		{
			post.DELETE("", usecases.DeletePost)
			post.GET("/likes", usecases.Likes)
		}
	}
}
