package usecases

import (
	"task_manager/domain"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

// mock for TaskRepository

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) Insert(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepo) FindByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Update(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepo) FindAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

// start unit tests for TaskUseCase

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	task := domain.Task{ID: "1", Title: "Test Task"}

	mockRepo.On("Insert", task).Return(nil)

	err := uc.CreateTask(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	expected := domain.Task{ID: "1", Title: "Test Task"}
	mockRepo.On("FindByID", "1").Return(expected, nil)

	task, err := uc.GetTask("1")

	assert.NoError(t, err)
	assert.Equal(t, expected, task)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	task := domain.Task{Title: "Updated"}
	mockRepo.On("Update", "1", task).Return(nil)

	err := uc.UpdateTask("1", task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := uc.DeleteTask("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	tasks := []domain.Task{
		{ID: "1", Title: "One"},
		{ID: "2", Title: "Two"},
	}
	mockRepo.On("FindAll").Return(tasks, nil)

	result, err := uc.ListTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}
