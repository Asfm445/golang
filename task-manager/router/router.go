package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTasks)
	router.POST("/tasks", controllers.PostTasks)
	router.PUT("/tasks/:id", controllers.PutTasks)
	router.DELETE("/tasks/:id", controllers.DeleteTasks)

	return router
}
