package router

import (
	"task_manager/Delivery/controllers"
	"task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// 🔓 Public Routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// 🔐 Protected Task Routes
	router.GET("/tasks", infrastructure.AuthMiddleware("user"), taskController.GetAllTasks)
	router.GET("/tasks/:id", infrastructure.AuthMiddleware("user"), taskController.GetTaskByID)
	router.POST("/tasks", infrastructure.AuthMiddleware("admin"), taskController.CreateTask)
	router.PUT("/tasks/:id", infrastructure.AuthMiddleware("admin"), taskController.UpdateTask)
	router.DELETE("/tasks/:id", infrastructure.AuthMiddleware("admin"), taskController.DeleteTask)

	// 🔐 User Promotion (optional, still global or inside userController)
	router.PATCH("/promote", infrastructure.AuthMiddleware("admin"), userController.Promote)

	return router
}
