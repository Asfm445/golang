package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func UserRegistration(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// User registration logic
	err := data.UserRegistration(user)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})

}

func Promote(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if user.Email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}
	err := data.PromoteUser(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User promoted successfully"})
}

func UserLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	jwtToken, err := data.UserLogin(user.Email, user.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
}

func GetTasks(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		task, err := data.FindTaskByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, task)
		return
	}

	tasks, err := data.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not fetch tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTasks(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	if err := data.InsertTask(newTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not create task"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTask)
}

func PutTasks(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func DeleteTasks(c *gin.Context) {
	id := c.Param("id")

	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
}
