// repositories/real_mongo_collection.go
package repositories

import (
	"context"

	"task_manager/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RealMongoCollection struct {
	collection *mongo.Collection
}

func NewRealMongoCollection(col *mongo.Collection) domain.MongoCollection {
	return &RealMongoCollection{collection: col}
}

func (r *RealMongoCollection) InsertOne(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, doc, opts...)
}

func (r *RealMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) domain.MongoSingleResult {
	return r.collection.FindOne(ctx, filter, opts...)
}

func (r *RealMongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, filter, update, opts...)
}

func (r *RealMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, filter, opts...)
}

func (r *RealMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return r.collection.Find(ctx, filter, opts...)
}
