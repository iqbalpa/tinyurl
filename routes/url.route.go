package routes

import "github.com/gofiber/fiber/v2"

func UrlRoutes(api fiber.Router) {
	api.Route("/url", func(urlRouter fiber.Router) {
		urlRouter.Get("/", testHandler)
	})
}

func testHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "currently at /api/url/",
	})
}

