package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
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

	defer client.Disconnect(context.Background())

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
	app.Get("/api/v1/todos/:id", getToDo)
	app.Post("/api/v1/todos", createToDo)
	app.Patch("/api/v1/todos/:id", updateToDo)
	// app.Delete("/api/v1/todos/:id", deleteToDo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getToDos(c *fiber.Ctx) error {
	var todos []ToDo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo ToDo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func getToDo(c *fiber.Ctx) error {
	var todo ToDo

	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		return err
	}

	return c.JSON(todo)
}

func createToDo(c *fiber.Ctx) error {
	todo := new(ToDo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Body is required",
		})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateToDo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"message": "Todo Updated!"})
}

// func deleteToDo(c *fiber.Ctx) error {}
