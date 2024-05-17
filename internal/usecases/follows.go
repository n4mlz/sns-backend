package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func FollowUser(ctx *gin.Context) {
	var request UserNameDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUserName := userDomain.UserName(request.UserName)

	if !userDomain.Service.IsExistUserName(targetUserName) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	sourceUser, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sourceUser.Follow(targetUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := UserNameDto{
		UserName: targetUser.UserName.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func UnfollowUser(ctx *gin.Context) {
	var request UserNameDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUserName := userDomain.UserName(request.UserName)

	if !targetUserName.IsValid() {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	sourceUser, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sourceUser.Unfollow(targetUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := UserNameDto{
		UserName: targetUser.UserName.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func RejectUser(ctx *gin.Context) {
	var request UserNameDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUserName := userDomain.UserName(request.UserName)

	if !userDomain.Service.IsExistUserName(targetUserName) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	sourceUser, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !targetUser.IsFollowing(sourceUser) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "not requested"})
		return
	}

	err = targetUser.Unfollow(sourceUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := UserNameDto{
		UserName: targetUser.UserName.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func RequestedUsers(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followRequestUsers, err := user.FollowRequests()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []UserDto
	for _, followRequestUser := range followRequestUsers {
		response = append(response, UserDto{
			UserName:        followRequestUser.UserName.String(),
			DisplayName:     followRequestUser.DisplayName.String(),
			Biography:       followRequestUser.Biography.String(),
			CreatedAt:       followRequestUser.CreatedAt,
			FollowingStatus: userDomain.FOLLOWED,
			IconUrl:         followRequestUser.IconUrl.String(),
			BgImageUrl:      followRequestUser.BgImageUrl.String(),
		})
	}

	ctx.JSON(http.StatusOK, response)
}
