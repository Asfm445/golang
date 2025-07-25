package repositories

import (
	"context"
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepo struct {
	Collection *mongo.Collection
}

func NewUserMongoRepo(db *mongo.Database) *UserMongoRepo {
	return &UserMongoRepo{
		Collection: db.Collection("users"),
	}
}

func (r *UserMongoRepo) Register(user domain.User) error {
	filter := bson.M{"$or": []bson.M{{"_id": user.ID}, {"email": user.Email}}}
	err := r.Collection.FindOne(context.TODO(), filter).Err()
	if err == nil {
		return errors.New("user already exists")
	}

	count, err := r.Collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	_, err = r.Collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserMongoRepo) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func (r *UserMongoRepo) Promote(email string) error {
	_, err := r.Collection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}
