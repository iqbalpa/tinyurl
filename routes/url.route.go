package routes

import (
	"tinyurl/dto"

	"github.com/gofiber/fiber/v2"
)

func UrlRoutes(api fiber.Router) {
	api.Route("/url", func(urlRouter fiber.Router) {
		urlRouter.Get("/", testHandler)
		urlRouter.Post("/shorten", shortenUrlHandler)
	})
}

func testHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "currently at /api/url/",
	})
}

func shortenUrlHandler(c *fiber.Ctx) error {
	url := new(dto.UrlRequest)

	if err := c.BodyParser(url); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"shortUrl": url,
	})
}