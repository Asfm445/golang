package repositories

import (
	"context"
	"errors"
	"task_manager/domain"
	"task_manager/infrastructure"

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

	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	_, err = r.Collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserMongoRepo) Login(email, password string) (string, error) {
	user, err := r.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if !infrastructure.CheckPasswordHash(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return infrastructure.GenerateToken(user.ID, user.Email, user.Role)
}

func (r *UserMongoRepo) Promote(email string) error {
	_, err := r.Collection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func (r *UserMongoRepo) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}
