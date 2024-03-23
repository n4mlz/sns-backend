package models

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/model"
	"github.com/n4mlz/sns-backend/internal/repository/query"
	"github.com/rs/xid"
)

func CreatePost(ctx *gin.Context) {
	var request PostContentSchema
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.GetString("userId")

	newPost := &model.Post{
		ID:      xid.New().String(),
		UserID:  userId,
		Content: request.Content,
	}

	query.Post.WithContext(context.Background()).Save(newPost)

	user, _ := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(userId)).Take()

	poster := UserDisplaySchema{
		UserName:    user.UserName,
		DisplayName: user.DisplayName,
	}

	response := PostSchema{
		PostId:    newPost.ID,
		Poster:    poster,
		Content:   newPost.Content,
		Likes:     0,
		Liked:     false,
		Comments:  0,
		CreatedAt: newPost.CreatedAt}
	ctx.JSON(http.StatusOK, response)
}
