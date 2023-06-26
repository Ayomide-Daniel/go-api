package main

import (
	"os"

	"example.com/rest-api/configs"
	"example.com/rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	configs.LoadEnv()
	configs.ConnectToDB()
	configs.MigrateDB()
}

func main() {
	// Set up app
	app := fiber.New()

	// Routes
	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/v1/users", controllers.UserIndex)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
