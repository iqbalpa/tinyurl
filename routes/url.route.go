package routes

import (
	"tinyurl/dto"
	"tinyurl/services"

	"github.com/gofiber/fiber/v2"
)

type UrlRoute struct {
	router fiber.Router
	service *services.UrlService
}

func NewUrlRoute(router fiber.Router, service *services.UrlService) *UrlRoute {
	return &UrlRoute{
		router: router, 
		service: service,
	}
}

func (r *UrlRoute) UrlRoutes() {
	r.router.Route("/url", func(urlRouter fiber.Router) {
		urlRouter.Get("/", r.testHandler)
		urlRouter.Post("/shorten", r.shortenUrlHandler)
	})
}

func (r *UrlRoute) testHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "currently at /api/url/",
	})
}

func (r *UrlRoute) shortenUrlHandler(c *fiber.Ctx) error {
	url := new(dto.UrlRequest)

	if err := c.BodyParser(url); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"shortUrl": url,
	})
}