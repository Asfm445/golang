package repositories

import (
	"task_manager/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepositoryMongo(t *testing.T) {
	m := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer m.Close()

	m.Run("sucessful register", func(mt *mtest.T) {
		repo := &UserMongoRepo{Collection: mt.Coll}

		user := domain.User{
			ID:       "1",
			Email:    "test@example.com",
			Password: "securepass",
		}

		// Simulate: FindOne returns ErrNoDocuments
		mt.AddMockResponses(
			// FindOne will return no documents (simulate user not existing)
			mtest.CreateCursorResponse(0, "test.users", mtest.FirstBatch),
			// CountDocuments returns 0
			mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
				bson.D{
					{"n", int32(0)}, // "n" must match the expected format
				},
			),
			// InsertOne succeeds
			mtest.CreateSuccessResponse(),
		)

		err := repo.Register(user)
		assert.NoError(mt, err)
	})

	m.Run("user already exists", func(mt *mtest.T) {
		repo := &UserMongoRepo{Collection: mt.Coll}
		user := domain.User{
			ID:       "1",
			Email:    "test@example.com",
			Password: "securepass",
		}

		// Simulate: FindOne returns a user document
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
				bson.D{
					{"_id", user.ID},
					{"email", user.Email},
				},
			),
		)
		err := repo.Register(user)
		assert.Error(mt, err)
		assert.Equal(mt, "user already exists", err.Error())
	})

	m.Run("find user by email", func(mt *mtest.T) {
		repo := &UserMongoRepo{Collection: mt.Coll}
		user := domain.User{
			ID:    "1",
			Email: "test@example.com",
		}
		// Simulate: FindOne returns a user document
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
				bson.D{
					{"_id", user.ID},
					{"email", user.Email},
				},
			),
		)
		foundUser, err := repo.FindByEmail(user.Email)
		assert.NoError(mt, err)
		assert.Equal(mt, user.ID, foundUser.ID)
		assert.Equal(mt, user.Email, foundUser.Email)
	})

	m.Run("promote user to admin", func(mt *mtest.T) {
		repo := &UserMongoRepo{Collection: mt.Coll}
		email := "test@example.com"
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(),
		)
		err := repo.Promote(email)
		assert.NoError(mt, err)

	})

}
