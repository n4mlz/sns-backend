package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/validation"
)

func authMiddleware(firebaseApp validation.FirebaseApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := firebaseApp.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		ctx.Set("userId", token.Claims["user_id"].(string))
		ctx.Next()
	}
}
