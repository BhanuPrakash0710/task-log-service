package main

import (
	"github.com/BhanuPrakash0710/to-do-list-api/internal/middleware"
	"github.com/BhanuPrakash0710/to-do-list-api/models"
	"github.com/BhanuPrakash0710/to-do-list-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	models.Setup()

	router := gin.Default()
	// userRoutes := router.Group("/user")
	// {
	// 	routes.RegisterUserRoutes(userRoutes)
	// }
	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(middleware.JWTAuthMiddleware())
	{
		routes.RegisterTaskRoutes(taskRoutes)
	}
	authRoutes := router.Group("/auth")
	{
		routes.RegisterAuthRoutes(authRoutes)
	}

	router.Run(":8083")
}
