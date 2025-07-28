package repositories

import (
	"testing"
	"time"

	"task_manager/domain"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestTaskRepositoryMongo(t *testing.T) {
	m := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	m.Run("Insert", func(mt *mtest.T) {
		repo := &TaskRepositoryMongo{collection: mt.Coll}
		task := domain.Task{ID: "1", Title: "Insert Task"}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := repo.Insert(task)
		assert.NoError(mt, err)
	})
	m.Run("Update", func(mt *mtest.T) {
		repo := &TaskRepositoryMongo{collection: mt.Coll}
		task := domain.Task{ID: "1", Title: "Updated Task"}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := repo.Update("1", task)
		assert.NoError(mt, err)
	})
	m.Run("Delete", func(mt *mtest.T) {
		repo := &TaskRepositoryMongo{collection: mt.Coll}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := repo.Delete("1")
		assert.NoError(mt, err)
	})
	m.Run("returns task by ID", func(mt *mtest.T) {
		repo := &TaskRepositoryMongo{collection: mt.Coll}

		expected := domain.Task{
			ID:          "1",
			Title:       "Found Task",
			Description: "Test Description",
			Status:      "todo",
			DueDate:     time.Date(2025, 7, 28, 12, 0, 0, 0, time.UTC),
		}

		// Create a BSON document simulating a MongoDB document
		doc := bson.D{
			{Key: "id", Value: expected.ID},
			{Key: "title", Value: expected.Title},
			{Key: "description", Value: expected.Description},
			{Key: "status", Value: expected.Status},
			{Key: "due_date", Value: primitive.NewDateTimeFromTime(expected.DueDate)},
		}

		// Return a successful FindOne response
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "tasks.collection", mtest.FirstBatch, doc))

		// Call the function
		result, err := repo.FindByID("1")

		// Assertions
		assert.NoError(mt, err)
		assert.Equal(mt, expected.ID, result.ID)
		assert.Equal(mt, expected.Title, result.Title)
		assert.Equal(mt, expected.Description, result.Description)
		assert.Equal(mt, expected.Status, result.Status)
		assert.WithinDuration(mt, expected.DueDate, result.DueDate, time.Second)
	})
	m.Run("returns all tasks", func(mt *mtest.T) {
		repo := &TaskRepositoryMongo{collection: mt.Coll}

		// Expected tasks
		expectedTasks := []domain.Task{
			{
				ID:          "1",
				Title:       "Task 1",
				Description: "Desc 1",
				Status:      "todo",
				DueDate:     time.Date(2025, 7, 28, 15, 0, 0, 0, time.UTC),
			},
			{
				ID:          "2",
				Title:       "Task 2",
				Description: "Desc 2",
				Status:      "done",
				DueDate:     time.Date(2025, 8, 1, 9, 0, 0, 0, time.UTC),
			},
		}

		// Create BSON docs to mock
		first := bson.D{
			{Key: "id", Value: expectedTasks[0].ID},
			{Key: "title", Value: expectedTasks[0].Title},
			{Key: "description", Value: expectedTasks[0].Description},
			{Key: "status", Value: expectedTasks[0].Status},
			{Key: "due_date", Value: primitive.NewDateTimeFromTime(expectedTasks[0].DueDate)},
		}
		second := bson.D{
			{Key: "id", Value: expectedTasks[1].ID},
			{Key: "title", Value: expectedTasks[1].Title},
			{Key: "description", Value: expectedTasks[1].Description},
			{Key: "status", Value: expectedTasks[1].Status},
			{Key: "due_date", Value: primitive.NewDateTimeFromTime(expectedTasks[1].DueDate)},
		}

		// Mock cursor with two tasks
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "tasks.collection", mtest.FirstBatch, first, second),
			mtest.CreateCursorResponse(0, "tasks.collection", mtest.NextBatch),
		)

		// Call the function
		tasks, err := repo.FindAll()

		// Assertions
		assert.NoError(mt, err)
		assert.Len(mt, tasks, 2)
		assert.Equal(mt, expectedTasks[0].ID, tasks[0].ID)
		assert.Equal(mt, expectedTasks[1].ID, tasks[1].ID)
		assert.WithinDuration(mt, expectedTasks[0].DueDate, tasks[0].DueDate, time.Second)
		assert.WithinDuration(mt, expectedTasks[1].DueDate, tasks[1].DueDate, time.Second)
	})
}
