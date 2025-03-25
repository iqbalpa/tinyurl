package main

import (
	"tinyurl/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	routes.UrlRoutes(api)

	app.Listen(":3000")
}
