package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/models"
)

func setPostsRoutesFrom(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	posts.Use(authMiddleware())
	{
		posts.POST("", models.CreatePost)
	}
}
