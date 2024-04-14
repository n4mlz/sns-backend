package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func CreatePost(ctx *gin.Context) {
	var request PostContentDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content := postDomain.Content(request.Content)

	post, err := postDomain.Factory.CreatePostToRepository(user, content)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poster := UserDisplayDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
	}

	response := PostDto{
		PostId:    post.PostId.String(),
		Poster:    poster,
		Content:   post.Content.String(),
		Likes:     0,
		Liked:     false,
		Comments:  0,
		CreatedAt: post.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
