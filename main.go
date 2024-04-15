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

	userRepository := &repository.UserRepository{}

	userFactory := userDomain.NewUserFactory(userRepository)
	userDomain.SetDefaultUserFactory(userFactory)

	userService := userDomain.NewUserService(userRepository)
	userDomain.SetDefaultUserService(userService)

	postRepository := &repository.PostRepository{}

	postFactory := postDomain.NewPostFactory(postRepository)
	postDomain.SetDefaultPostFactory(postFactory)

	postService := postDomain.NewPostService(postRepository)
	postDomain.SetDefaultPostService(postService)

	h.SetupRoutes()

	h.Router.Run(":8080")
}
