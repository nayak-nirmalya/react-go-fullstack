package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello, World! 🚀")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File: ", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	// PORT := os.Getenv("PORT")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Error Connection to MongoDB: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error Connection to MongoDB: ", err)
	}

	fmt.Println("Connected to MongoDB Atlas. ☁️")

	collection = client.Database("go-lang-react").Collection("todos")
}
