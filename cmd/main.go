package main

import (
	"wedding-invitation/config"
	"wedding-invitation/internal/handler"
	"wedding-invitation/internal/middleware"
	"wedding-invitation/internal/repository"
	"wedding-invitation/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB)
	photoRepo := repository.NewPhotoRepository(config.DB)
	commentRepo := repository.NewCommentRepository(config.DB)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo)
	photoUsecase := usecase.NewPhotoUsecase(photoRepo)
	commentUsecase := usecase.NewCommentUsecase(commentRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)
	photoHandler := handler.NewPhotoHandler(photoUsecase)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	// Initialize the router
	router := gin.Default()

	// User routes
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// Protected routes
	auth := router.Group("/").Use(middleware.AuthMiddleware())
	{
		auth.POST("/photos", photoHandler.PostPhoto)
		auth.GET("/photos", photoHandler.GetPhotos)
		auth.POST("/comments", commentHandler.PostComment)
		auth.GET("/comments/:photo_id", commentHandler.GetComments)
	}

	// Start the server
	router.Run(":8080")
}
