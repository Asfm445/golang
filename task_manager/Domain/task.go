package domain

import "time"

type Task struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}
type TaskRepository interface {
	Insert(task Task) error
	FindByID(id string) (Task, error)
	Update(id string, task Task) error
	Delete(id string) error
	FindAll() ([]Task, error)
}
