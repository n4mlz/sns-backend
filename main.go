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

	firebaseApp, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}

	// db, err := repository.NewRepository()

	h.SetupRoutes(firebaseApp)

	h.Router.Run(":8080")
}
