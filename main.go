package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/handler"
	"github.com/n4mlz/sns-backend/internal/repository"
	"github.com/n4mlz/sns-backend/internal/repository/query"
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

	db, err := repository.NewRepository()
	if err != nil {
		return
	}

	query.SetDefault(db)

	h.SetupRoutes()

	h.Router.Run(":8080")
}
