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
		posts.GET("/timeline", usecases.Timeline)

		post := posts.Group("/:postId")
		{
			post.GET("", usecases.GetPost)
			post.DELETE("", usecases.DeletePost)
			post.GET("/likes", usecases.Likes)
		}

		comments := posts.Group("/comments")
		{
			comments.POST("", usecases.CreateComment)
			comments.DELETE("/:commentId", usecases.DeleteComment)
		}

		replies := posts.Group("/replies")
		{
			replies.POST("", usecases.CreateReply)
			replies.DELETE("/:replyId", usecases.DeleteReply)
		}

		notifications := posts.Group("/notifications")
		{
			notifications.GET("", usecases.GetNotifications)
			notifications.PUT("/confirm", usecases.ConfirmNotifications)
		}
	}
}
