package main

import (
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/router"
	"task_manager/infrastructure"
	"task_manager/repositories"
	"task_manager/usecases"
)

func main() {
	// Initialize MongoDB connection
	db := infrastructure.InitMongoDB()

	// ==== Repositories ====
	taskRepo := repositories.NewTaskRepositoryMongo(db)
	userRepo := repositories.NewUserMongoRepo(db)

	// ==== Usecases ====
	taskUsecase := usecases.NewTaskUseCase(taskRepo)
	userUsecase := usecases.NewUserUseCase(userRepo)

	// ==== Controllers ====
	taskController := controllers.NewTaskController(*taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// ==== Router ====
	r := router.SetupRouter(taskController, userController)

	// Start server
	r.Run("localhost:8081")
}
