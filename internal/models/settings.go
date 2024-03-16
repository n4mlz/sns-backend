package models

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/model"
	"github.com/n4mlz/sns-backend/internal/repository/query"
)

var (
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 16
)

type UserRequest struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
}

func isValidUserName(s string) bool {
	pattern := regexp.MustCompile(fmt.Sprintf(`^[A-Za-z0-9_-]{%d,%d}$`, MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH))
	return pattern.MatchString(s)
}

func isValidDisplayName(s string) bool {
	return len(s) != 0
}

func SaveUser(ctx *gin.Context) {
	var request UserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if !isValidUserName(request.UserName) || !isValidDisplayName(request.DisplayName) {
		ctx.JSON(http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	newUser := &model.User{
		ID:          ctx.GetString("userId"),
		UserName:    request.UserName,
		DisplayName: request.DisplayName,
		Biography:   &request.Biography}

	query.User.WithContext(context.Background()).Save(newUser)
}
