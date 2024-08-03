package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// The backend for a social media app using Go fibre
func main() {
	// Create a new fibre app
	app := fiber.New()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("Register")
	})

	app.Get("/api/user", func(c *fiber.Ctx) error {
		return c.SendString("User")
	})

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return c.SendString("Posts")
	})

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
