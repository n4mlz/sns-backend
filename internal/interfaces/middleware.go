package interfaces

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/infrastructure/validation"
)

var FRONTEND_URL = os.Getenv("FRONTEND_URL")

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

func authMiddlewareNoAbort() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := validation.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.Set("userId", "")
		} else {
			ctx.Set("userId", token.Claims["user_id"].(string))
		}

		ctx.Next()
	}
}

func SetCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			FRONTEND_URL,
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"Origin",
			"Access-Control-Request-Method",
			"Access-Control-Request-Methods",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Max-Age",
			"Access-Control-Allow-Credentials",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))
}
