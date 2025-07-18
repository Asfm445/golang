package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 🔓 Public routes (no auth)
	router.POST("/register", controllers.UserRegistration)
	router.POST("/login", controllers.UserLogin)

	// 🔐 Protected routes

	router.GET("/tasks", middleware.AuthMiddleware("user"), controllers.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware("user"), controllers.GetTasks)
	router.POST("/tasks", middleware.AuthMiddleware("admin"), controllers.PostTasks)
	router.PUT("/tasks/:id", middleware.AuthMiddleware("admin"), controllers.PutTasks)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware("admin"), controllers.DeleteTasks)
	router.PATCH("/promote", middleware.AuthMiddleware("admin"), controllers.Promote)

	return router
}
