package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/handler"
	"github.com/n4mlz/sns-backend/internal/validation"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true

	h := handler.NewHandler(r)

	err := validation.InitFirebaseApp()
	if err != nil {
		return
	}

	h.SetupRoutes()

	h.Router.Run(":8080")
}
