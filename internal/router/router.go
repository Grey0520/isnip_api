package router

import (
	// "github.com/Grey0520/isnip_api/v2/database"
	// "github.com/Grey0520/isnip_api/v2/models"
	"github.com/Grey0520/isnip_api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", Hello)
	app.Post("/login", handler.Login)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}
