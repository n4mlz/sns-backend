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

	user, err := userDomain.Factory.SaveUserToRepository(userId, userName, displayName, biography)
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

func SaveIcon(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := userDomain.UserId(ctx.GetString("userId"))
	user, err := userDomain.Factory.GetUser(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	img, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer img.Close()

	err = user.SaveIcon(img)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func SaveBgImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := userDomain.UserId(ctx.GetString("userId"))
	user, err := userDomain.Factory.GetUser(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	img, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer img.Close()

	err = user.SaveBgImage(img)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
