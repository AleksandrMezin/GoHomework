package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// News представляет данные о новости в базе данных
type News struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"pubDate"`
	SourceURL   string    `json:"sourceURL"`
}

// InitDB инициализирует базу данных MongoDB
func InitDB(connectionString, dbName string) error {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("Error pinging MongoDB: %v", err)
	}

	db = client.Database(dbName)
	log.Println("Connected to MongoDB")

	return nil
}

// SaveNews сохраняет новость в базе данных
func SaveNews(news News) {
	_, err := db.Collection("news").InsertOne(context.Background(), news)
	if err != nil {
		log.Printf("Error saving news to MongoDB: %v", err)
	}
}

// GetNews возвращает заданное количество последних новостей из базы данных
func GetNews(count int) ([]News, error) {
	var newsList []News

	limit := int64(count) // Преобразуем count в int64

	cursor, err := db.Collection("news").Find(context.Background(), nil, &options.FindOptions{
		Sort:  map[string]int{"pubdate": -1},
		Limit: &limit, // Передаем указатель на int64
	})
	if err != nil {
		return nil, fmt.Errorf("Error fetching news from MongoDB: %v", err)
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &newsList)
	if err != nil {
		return nil, fmt.Errorf("Error reading news from MongoDB: %v", err)
	}

	return newsList, nil
}
