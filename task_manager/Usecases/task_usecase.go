package usecases

import "task_manager/domain"

type TaskUseCaseInterface interface {
	CreateTask(task domain.Task) error
	GetTask(id string) (domain.Task, error)
	UpdateTask(id string, task domain.Task) error
	DeleteTask(id string) error
	ListTasks() ([]domain.Task, error)
}

type TaskUseCase struct {
	repo domain.TaskRepository
}

func NewTaskUseCase(r domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: r}
}

func (uc *TaskUseCase) CreateTask(task domain.Task) error {
	return uc.repo.Insert(task)
}

func (uc *TaskUseCase) GetTask(id string) (domain.Task, error) {
	return uc.repo.FindByID(id)
}

func (uc *TaskUseCase) UpdateTask(id string, task domain.Task) error {
	return uc.repo.Update(id, task)
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	return uc.repo.Delete(id)
}

func (uc *TaskUseCase) ListTasks() ([]domain.Task, error) {
	return uc.repo.FindAll()
}
