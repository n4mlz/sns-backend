package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/handler"
	"github.com/n4mlz/sns-backend/validation"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true

	h := handler.NewHandler(r)

	firebaseApp, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}

	h.SetupRoutes(firebaseApp)

	h.Router.Run(":8080")
}
