package main

import (
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World! 🚀"})
	})

	log.Fatal(app.Listen(":8080"))
}
