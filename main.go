package main

import (
	"api/config"
	"api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(config.FiberConfig())

	app.Use(logger.New())
	app.Use(recover.New())

	routes.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Route not found",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
