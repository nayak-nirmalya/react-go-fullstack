package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
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
	PORT := os.Getenv("PORT")

}
