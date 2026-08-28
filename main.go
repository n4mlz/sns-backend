package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/query"
	"github.com/n4mlz/sns-backend/internal/infrastructure/s3"
	"github.com/n4mlz/sns-backend/internal/infrastructure/validation"
	"github.com/n4mlz/sns-backend/internal/interfaces"
	"github.com/n4mlz/sns-backend/internal/monitoring"
)

func main() {
	r := gin.Default()
	if err := r.SetTrustedProxies([]string{"127.0.0.1", "::1", "172.16.0.0/12"}); err != nil {
		log.Fatal(err)
	}
	r.Use(monitoring.Middleware())
	r.Use(interfaces.SecurityHeaders())
	r.Use(interfaces.RequestSizeLimit())
	r.Use(interfaces.RateLimit())
	monitoring.Routes(r)
	interfaces.SetCors(r)
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
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	if err := sqlDB.Ping(); err != nil {
		return
	}

	query.SetDefault(db)

	userRepository := &repository.UserRepository{}

	s3app, err := s3.NewS3App()
	if err != nil {
		return
	}
	userImageRepository := s3app

	userFactory := userDomain.NewUserFactory(userRepository, userImageRepository)
	userDomain.SetDefaultUserFactory(userFactory)

	userService := userDomain.NewUserService(userRepository)
	userDomain.SetDefaultUserService(userService)

	postRepository := &repository.PostRepository{}

	postFactory := postDomain.NewPostFactory(postRepository)
	postDomain.SetDefaultPostFactory(postFactory)

	postService := postDomain.NewPostService(postRepository)
	postDomain.SetDefaultPostService(postService)

	h.SetupRoutes()
	monitoring.SetReady(true)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           h.Router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
