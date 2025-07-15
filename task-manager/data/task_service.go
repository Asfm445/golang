package data

import (
	"errors"
	"task_manager/models"
)

func FindTaskByID(id string) (models.Task, error) {
	for _, task := range models.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func AppendToTasks(newTask models.Task) {
	models.Tasks = append(models.Tasks, newTask)
}

func UpdateTask(id string, updatedTask models.Task) bool {
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks[i] = updatedTask
			return true
		}
	}
	return false
}

func DeleteTask(id string) bool {
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return true
		}
	}
	return false
}
