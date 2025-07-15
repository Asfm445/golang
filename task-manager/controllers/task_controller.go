package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

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
	c.IndentedJSON(http.StatusOK, models.Tasks)
}

func PostTasks(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	// Add the new album to the slice.
	data.AppendToTasks(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func PutTasks(c *gin.Context) {
	var updatedTask models.Task
	id := c.Param("id")
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}
	if data.UpdateTask(id, updatedTask) {
		c.IndentedJSON(http.StatusOK, updatedTask)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	// return
}

func DeleteTasks(c *gin.Context) {
	id := c.Param("id")
	if data.DeleteTask(id) {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "task deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
