package models

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/model"
	"github.com/n4mlz/sns-backend/internal/repository/query"
)

type UserRequest struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
}

// TODO: 後でここをリファクタリングする
func SaveUser(ctx *gin.Context) {
	var request UserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	query.User.WithContext(context.Background()).Save(&model.User{ID: ctx.GetString("userId"), UserName: request.UserName, DisplayName: request.DisplayName, Biography: &request.Biography})
}
