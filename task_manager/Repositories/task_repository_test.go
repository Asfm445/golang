// repositories/task_repository_test.go

package repositories

import (
	"context"
	"task_manager/domain"
	"task_manager/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestInsert(t *testing.T) {
	mockCol := mocks.NewMockCollection()

	repo := NewTaskRepositoryMongoFromCollection(mockCol)

	task := domain.Task{
		ID:          "abc",
		Title:       "Test",
		Description: "test desc",
		Status:      "pending",
	}

	mockCol.On("InsertOne", mock.Anything, task).
		Return(&mongo.InsertOneResult{}, nil)

	err := repo.Insert(task)
	assert.Nil(t, err)

	mockCol.AssertExpectations(t)
}

func TestFindByID(t *testing.T) {
	mockCol := mocks.NewMockCollection()
	mockSingleResult := new(mocks.MockSingleResult)

	repo := NewTaskRepositoryMongoFromCollection(mockCol)

	task := domain.Task{
		ID:          "abc",
		Title:       "Test",
		Description: "test desc",
		Status:      "pending",
	}

	mockCol.On("FindOne", mock.Anything, bson.M{"id": "abc"}).
		Return(mockSingleResult)

	mockSingleResult.On("Decode", mock.AnythingOfType("*domain.Task")).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Task)
			*arg = task
		}).Return(nil)

	result, err := repo.FindByID("abc")

	assert.Nil(t, err)
	assert.Equal(t, task, result)

	mockCol.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockCol := mocks.NewMockCollection()
	repo := NewTaskRepositoryMongoFromCollection(mockCol)

	task := domain.Task{
		ID:          "abc",
		Title:       "Updated",
		Description: "updated desc",
		Status:      "done",
	}

	mockCol.On("UpdateOne", mock.Anything, bson.M{"id": "abc"}, mock.Anything).
		Return(&mongo.UpdateResult{}, nil)

	err := repo.Update("abc", task)
	assert.Nil(t, err)

	mockCol.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockCol := mocks.NewMockCollection()
	repo := NewTaskRepositoryMongoFromCollection(mockCol)

	mockCol.On("DeleteOne", mock.Anything, bson.M{"id": "abc"}).
		Return(&mongo.DeleteResult{}, nil)

	err := repo.Delete("abc")
	assert.Nil(t, err)

	mockCol.AssertExpectations(t)
}

func TestFindAll(t *testing.T) {
	mockCol := mocks.NewMockCollection()
	repo := NewTaskRepositoryMongoFromCollection(mockCol)

	expected := []domain.Task{
		{ID: "1", Title: "Task 1", Description: "Desc 1", Status: "pending"},
		{ID: "2", Title: "Task 2", Description: "Desc 2", Status: "done"},
	}

	// Create a mock cursor with Next and Decode behaviors
	mockCursor := new(mocks.MockCursor)
	index := 0
	mockCursor.On("Next", mock.Anything).Return(func(ctx context.Context) bool {
		if index < len(expected) {
			return true
		}
		return false
	}).Run(func(args mock.Arguments) {
		index++
	}).Times(len(expected) + 1)

	mockCursor.On("Decode", mock.AnythingOfType("*domain.Task")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.Task)
		*arg = expected[index-1]
	})

	mockCursor.On("Close", mock.Anything).Return(nil)

	mockCol.On("Find", mock.Anything, bson.M{}).Return(mockCursor, nil)

	result, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	mockCol.AssertExpectations(t)
	mockCursor.AssertExpectations(t)
}
