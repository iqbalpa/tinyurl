package main

import (
	"tinyurl/config"
	"tinyurl/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDb()

	app := fiber.New()
	api := app.Group("/api")

	routes.UrlRoutes(api)

	app.Listen(":3000")
}
