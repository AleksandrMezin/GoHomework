package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoDBStorage(connectionString, dbName, collectionName string) (*MongoDBStorage, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &MongoDBStorage{
		client:     client,
		database:   db,
		collection: collection,
	}, nil
}

func (ms *MongoDBStorage) Posts() ([]storage.Post, error) {
	cursor, err := ms.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error querying posts: %v", err)
	}
	defer cursor.Close(context.Background())

	var posts []storage.Post
	for cursor.Next(context.Background()) {
		var post storage.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, fmt.Errorf("error decoding post: %v", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ms *MongoDBStorage) AddPost(post storage.Post) error {
	_, err := ms.collection.InsertOne(context.Background(), post)
	if err != nil {
		return fmt.Errorf("error inserting post: %v", err)
	}
	return nil
}

func (ms *MongoDBStorage) UpdatePost(post storage.Post) error {
	_, err := ms.collection.ReplaceOne(context.Background(), bson.M{"_id": post.ID}, post)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}
	return nil
}

func (ms *MongoDBStorage) DeletePost(post storage.Post) error {
	_, err := ms.collection.DeleteOne(context.Background(), bson.M{"_id": post.ID})
	if err != nil {
		return fmt.Errorf("error deleting post: %v", err)
	}
	return nil
}
