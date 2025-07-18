package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task_manager/models"
)

var client *mongo.Client
var taskCollection *mongo.Collection

// var userCollection *mongo.Collection

func ConnectToDb() {
	// Set options and connect
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Select DB and collection
	taskCollection = client.Database("taskdb").Collection("tasks")
	userCollection = client.Database("taskdb").Collection("users")
	// userCollection = client.Database("taskdb").Collection("users")
	log.Println("✅ Connected to MongoDB and task collection initialized.")
}

func FindTaskByID(id string) (models.Task, error) {
	filter := bson.M{"id": id}
	var result models.Task
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Task{}, err
	}
	return result, nil
}

func InsertTask(newTask models.Task) error {
	_, err := taskCollection.InsertOne(context.TODO(), newTask)
	return err
}

func UpdateTask(id string, updatedTask models.Task) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedTask}

	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteTask(id string) error {
	filter := bson.M{"id": id}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	return err
}

func GetAllTasks() ([]models.Task, error) {
	cursor, err := taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
