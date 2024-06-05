package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ToDo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!")
	}

	PORT := os.Getenv("PORT")

	todos := []ToDo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World! 🚀"})
	})

	// Get all ToDos
	app.Get("/api/v1/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Get a ToDo
	app.Get("/api/v1/todos/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID!"})
		}

		if id <= 0 || id > len(todos) {
			return c.Status(404).JSON(fiber.Map{"error": "ToDo Not Found!"})
		}

		return c.Status(200).JSON(todos[id-1])
	})

	// Create a ToDo
	app.Post("/api/v1/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ToDo Body is Required!"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update a ToDo
	app.Patch("/api/v1/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "ToDo Not Found!"})
	})

	// Delete a ToDo
	app.Delete("/api/v1/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}

		return c.Status(400).JSON(fiber.Map{"success": false})
	})

	log.Fatal(app.Listen(":" + PORT))
}
