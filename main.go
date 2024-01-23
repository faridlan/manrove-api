// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route for the root path
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mangrove!")
	})

	// Define a route for an API endpoint
	app.Get("/api/greet/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{"message": "Hello, " + name + "!"})
	})

	// Start the Fiber app on port 3030
	err := app.Listen(":3030")
	if err != nil {
		panic(err)
	}
}
