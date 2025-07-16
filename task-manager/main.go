package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.ConnectToDb()
	router.SetupRouter().Run("localhost:8081")
}
