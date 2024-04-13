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

	userId := userDomain.UserId(ctx.GetString("userId"))
	userName := userDomain.UserName(request.UserName)
	displayName := userDomain.DisplayName(request.DisplayName)
	biography := userDomain.Biography(request.Biography)

	user, err := userDomain.Factory.NewUser(userId, userName, displayName, biography)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.SaveUser()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := ProfileDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
		Biography:   user.Biography.String(),
	}

	ctx.JSON(http.StatusOK, response)
}
