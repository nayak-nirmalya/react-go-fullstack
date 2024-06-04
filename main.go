package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ToDo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	todos := []ToDo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World! 🚀"})
	})

	// Get all ToDos
	app.Get("/api/v1/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
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
	})

	log.Fatal(app.Listen(":8080"))
}
