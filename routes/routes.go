package routes

import (
	"github.com/BhanuPrakash0710/to-do-list-api/controllers"
	"github.com/gin-gonic/gin"
)

// func RegisterUserRoutes(router *gin.RouterGroup) {
// 	router.GET("/", controllers.GetAllUsers)
// 	router.POST("/", controllers.CreateUser)
// 	router.GET("/:id", controllers.GetUserByID)
// 	router.PUT("/:id", controllers.UpdateUserByID)
// 	router.DELETE("/:id", controllers.DeleteUserByID)
// }

func RegisterTaskRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllTasks)
	router.POST("/", controllers.CreateTask)
	router.GET("/:id", controllers.GetTaskByID)
	router.PATCH("/:id", controllers.UpdateTaskByID)
	router.DELETE("/:id", controllers.DeleteTaskByID)
}

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	//router.GET("/logout", controllers.Logout)
	//router.GET("/profile", controllers.GetProfile)
	//router.PUT("/profile", controllers.UpdateProfile)
}
