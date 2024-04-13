package interfaces

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/infrastructure/validation"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := validation.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userId", token.Claims["user_id"].(string))
		ctx.Next()
	}
}
