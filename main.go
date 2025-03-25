package main

import (
	"tinyurl/config"
	"tinyurl/repository"
	"tinyurl/routes"
	"tinyurl/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDb()
	urlRepository := repository.NewUrlRepository(db)
	urlService := services.NewUrlService(urlRepository)

	app := fiber.New()
	api := app.Group("/api")

	routes.UrlRoutes(api)

	app.Listen(":3000")
}
