package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func SaveProfile(ctx *gin.Context) {
	var request ProfileDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userName := userDomain.UserName(request.UserName)
	displayName := userDomain.DisplayName(request.DisplayName)

	if !userName.IsValid() || !displayName.IsValid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userId := userDomain.UserId(ctx.GetString("userId"))

	user, err := userDomain.Factory.NewUser(userId, userName, displayName, request.Biography)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.SaveUser()

	response := ProfileDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
		Biography:   user.Biography}

	ctx.JSON(http.StatusOK, response)
}
