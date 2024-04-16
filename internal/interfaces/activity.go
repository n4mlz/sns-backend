package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/usecases"
)

func setActivityRoutesFrom(r *gin.RouterGroup) {
	r.GET("/timeline", usecases.Timeline)
	r.Use(authMiddleware())
}
