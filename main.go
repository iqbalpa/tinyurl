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

	urlRoute := routes.NewUrlRoute(api, urlService)
	urlRoute.UrlRoutes()

	app.Listen(":3000")
}
