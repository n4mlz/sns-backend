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

func DeletePost(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := postDomain.Factory.GetPost(postDomain.PostId(ctx.Param("postId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postDomain.Factory.DeletePostFromRepository(user, post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func LikePost(ctx *gin.Context) {
	var request PostIdDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := postDomain.Factory.GetPost(postDomain.PostId(request.PostId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = post.Like(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := PostIdDto{
		PostId: post.PostId.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func UnlikePost(ctx *gin.Context) {
	var request PostIdDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := postDomain.Factory.GetPost(postDomain.PostId(request.PostId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = post.Unlike(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := PostIdDto{
		PostId: post.PostId.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func Likes(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(ctx.Param("postId"))

	post, err := postDomain.Factory.GetPost(postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allLikers, err := post.GetLikers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutuals, err := user.MutualFollows()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutualSet := NewSet()
	for _, mutual := range mutuals {
		mutualSet.Add(mutual.UserId)
	}

	var response []UserDto
	for _, user := range allLikers {
		if mutualSet.Contains(user.UserId) {
			response = append(response, UserDto{
				UserName:        user.UserName.String(),
				DisplayName:     user.DisplayName.String(),
				Biography:       user.Biography.String(),
				CreatedAt:       user.CreatedAt,
				FollowingStatus: userDomain.MUTUAL,
			})
		}
	}

	ctx.JSON(http.StatusOK, response)
}
