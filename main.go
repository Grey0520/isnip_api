package main

import (
	"log"

	"github.com/Grey0520/isnip_api/v2/database"
	"github.com/Grey0520/isnip_api/v2/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	router.SetupRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
