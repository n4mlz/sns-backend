package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func Timeline(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeline, err := postDomain.Factory.GetPostsByVisibleUsers(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []PostDto
	for _, post := range timeline {
		poster := UserDisplayDto{
			UserName:    post.Poster.UserName.String(),
			DisplayName: post.Poster.DisplayName.String(),
		}

		response = append(response, PostDto{
			PostId:    post.PostId.String(),
			Poster:    poster,
			Content:   post.Content.String(),
			CreatedAt: post.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
