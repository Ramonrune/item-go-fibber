package routes

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	items := v1.Group("/items")

	items.Get("/", handlers.GetItems)
	items.Get("/:id", handlers.GetItem)
	items.Post("/", handlers.CreateItem)
	items.Put("/:id", handlers.UpdateItem)
	items.Delete("/:id", handlers.DeleteItem)
}
