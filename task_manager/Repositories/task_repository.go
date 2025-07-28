package repositories

import (
	"context"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryMongo struct {
	collection *mongo.Collection
}

func NewTaskRepositoryMongo(db *mongo.Database) *TaskRepositoryMongo {
	return &TaskRepositoryMongo{
		collection: db.Collection("tasks"),
	}
}

func (r *TaskRepositoryMongo) Insert(task domain.Task) error {
	_, err := r.collection.InsertOne(context.TODO(), task)
	return err
}

func (r *TaskRepositoryMongo) FindByID(id string) (domain.Task, error) {
	filter := bson.M{"id": id}
	var task domain.Task
	err := r.collection.FindOne(context.TODO(), filter).Decode(&task)
	return task, err
}

func (r *TaskRepositoryMongo) Update(id string, task domain.Task) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": task}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *TaskRepositoryMongo) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (r *TaskRepositoryMongo) FindAll() ([]domain.Task, error) {
	ctx := context.TODO()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
