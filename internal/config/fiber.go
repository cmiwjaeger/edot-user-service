package config

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Default to 500 Internal Server Error
		code := fiber.StatusInternalServerError

		// Check if the error is a Fiber error (which contains a status code)
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// Log the error
		log.Printf("Error: %v", err)

		// Respond with the error message and status code
		return c.Status(code).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

}
