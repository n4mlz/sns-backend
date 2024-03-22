package models

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/query"
)

func User(ctx *gin.Context) {
	userName := ctx.Param("userName")
	user, err := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq(userName)).Take()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := UserDataSchema{
		UserName:        user.UserName,
		DisplayName:     user.DisplayName,
		Biography:       user.Biography,
		CreatedAt:       user.CreatedAt,
		FollowingStatus: getFollowingStatus(ctx.GetString("userId"), user.ID),
	}

	ctx.JSON(http.StatusOK, response)
}

func MutualFollow(ctx *gin.Context) {
	userName := ctx.Param("userName")

	if !isExistUser(userName) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	userId := userNameToUserId(userName)

	mutualList, err := getMutualFollowUserIdList(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutualUsers, _ := query.User.WithContext(context.Background()).Where(query.User.ID.In(mutualList...)).Find()

	var response []UserDataSchema
	for _, user := range mutualUsers {
		response = append(response, UserDataSchema{
			UserName:        user.UserName,
			DisplayName:     user.DisplayName,
			Biography:       user.Biography,
			CreatedAt:       user.CreatedAt,
			FollowingStatus: getFollowingStatus(ctx.GetString("userId"), user.ID),
		})
	}

	ctx.JSON(http.StatusOK, response)
}
