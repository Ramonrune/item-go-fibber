package config

import (
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		},
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "API",
		ReadTimeout:   60,
	}
}
