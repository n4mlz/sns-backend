package models

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/model"
	"github.com/n4mlz/sns-backend/internal/repository/query"
)

func SaveProfile(ctx *gin.Context) {
	var request ProfileSchema
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidUserName(request.UserName) || !isValidDisplayName(request.DisplayName) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userId := ctx.GetString("userId")

	newUser := &model.User{
		ID:          userId,
		UserName:    request.UserName,
		DisplayName: request.DisplayName,
		Biography:   request.Biography}

	query.User.WithContext(context.Background()).Save(newUser)

	newFollow := &model.Follow{
		ID:              userId,
		FollowerUserID:  userId,
		FollowingUserID: userId}

	query.Follow.WithContext(context.Background()).Save(newFollow)

	response := ProfileSchema{
		UserName:    newUser.UserName,
		DisplayName: newUser.DisplayName,
		Biography:   newUser.Biography}

	ctx.JSON(http.StatusOK, response)
}
