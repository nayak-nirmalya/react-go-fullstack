package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/v1/todos", getToDos)
	app.Get("/api/v1/todos:id", getToDo)
	app.Post("/api/v1/todos", createToDo)
	app.Patch("/api/v1/todos:id", updateToDo)
	app.Delete("/api/v1/todos:id", deleteToDo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getToDos(c *fiber.Ctx) error   {}
func getToDo(c *fiber.Ctx) error    {}
func createToDo(c *fiber.Ctx) error {}
func updateToDo(c *fiber.Ctx) error {}
func deleteToDo(c *fiber.Ctx) error {}
