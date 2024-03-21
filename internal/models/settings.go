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

	newUser := &model.User{
		ID:          ctx.GetString("userId"),
		UserName:    request.UserName,
		DisplayName: request.DisplayName,
		Biography:   &request.Biography}

	query.User.WithContext(context.Background()).Save(newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"userName":    newUser.UserName,
		"displayName": newUser.DisplayName,
		"biography":   newUser.Biography})
}
