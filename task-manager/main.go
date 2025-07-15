package main

import (
	"task_manager/router"
)

func main() {
	router.SetupRouter().Run("localhost:8080")
}
