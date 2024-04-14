package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/query"
	"github.com/n4mlz/sns-backend/internal/infrastructure/validation"
	"github.com/n4mlz/sns-backend/internal/interfaces"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true

	h := interfaces.NewHandler(r)

	err := validation.InitFirebaseApp()
	if err != nil {
		return
	}

	db, err := repository.NewRepository()
	if err != nil {
		return
	}

	query.SetDefault(db)

	userFactory := userDomain.NewUserFactory(&repository.UserRepository{})
	userDomain.SetDefaultUserFactory(userFactory)

	userService := userDomain.NewUserService(&repository.UserRepository{})
	userDomain.SetDefaultUserService(userService)

	postFactory := postDomain.NewPostFactory(&repository.PostRepository{})
	postDomain.SetDefaultPostFactory(postFactory)

	h.SetupRoutes()

	h.Router.Run(":8080")
}
