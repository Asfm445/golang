// file: task_controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTaskUseCase implements the usecase with testify.Mock
type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUseCase) ListTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTask_Success(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	task := domain.Task{ID: "1", Title: "New Task"}
	mockUC.On("CreateTask", task).Return(nil)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	body, _ := json.Marshal(task)
	ctx.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.CreateTask(ctx)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestCreateTask_BadRequest(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte("bad json")))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.CreateTask(ctx)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	task := domain.Task{ID: "1", Title: "Test Task"}
	mockUC.On("GetTask", "1").Return(task, nil)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	controller.GetTaskByID(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	mockUC.On("GetTask", "1").Return(domain.Task{}, errors.New("not found"))

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	controller.GetTaskByID(ctx)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteTask_Success(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	mockUC.On("DeleteTask", "1").Return(nil)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	controller.DeleteTask(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockUC := new(MockTaskUseCase)
	controller := NewTaskController(mockUC)

	mockUC.On("DeleteTask", "1").Return(errors.New("not found"))

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	controller.DeleteTask(ctx)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
