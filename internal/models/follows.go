package models

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/repository/model"
	"github.com/n4mlz/sns-backend/internal/repository/query"
	"github.com/rs/xid"
)

func FollowUser(ctx *gin.Context) {
	var request UserNameSchema
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isExistUser(request.UserName) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	fromUserId := ctx.GetString("userId")
	toUserId := userNameToUserId(request.UserName)

	if isFollowing(fromUserId, toUserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "already following"})
		return
	}

	newFollow := &model.Follow{
		ID:              xid.New().String(),
		FollowerUserID:  fromUserId,
		FollowingUserID: toUserId}

	query.Follow.WithContext(context.Background()).Save(newFollow)

	ctx.JSON(http.StatusOK, gin.H{
		"followingUserName": request.UserName,
	})
}

func UnfollowUser(ctx *gin.Context) {
	var request UserNameSchema
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isExistUser(request.UserName) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	fromUserId := ctx.GetString("userId")
	toUserId := userNameToUserId(request.UserName)

	if !isFollowing(fromUserId, toUserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "not following"})
		return
	}

	query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(fromUserId)).Where(query.Follow.FollowingUserID.Eq(toUserId)).Delete()

	ctx.JSON(http.StatusOK, gin.H{
		"unfollowingUserName": request.UserName,
	})
}

func Followed(ctx *gin.Context) {
	userId := ctx.GetString("userId")

	FollowRequestList, err := getFollowRequestUserIdList(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	FollowRequests, _ := query.User.WithContext(context.Background()).Where(query.User.ID.In(FollowRequestList...)).Find()

	var response []UserDataSchema
	for _, user := range FollowRequests {
		response = append(response, UserDataSchema{
			UserName:        user.UserName,
			DisplayName:     user.DisplayName,
			Biography:       user.Biography,
			CreatedAt:       user.CreatedAt,
			FollowingStatus: FOLLOWED,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
