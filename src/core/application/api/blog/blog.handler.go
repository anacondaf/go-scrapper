package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/api/blog/service"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

func CreatePost(colly *webScraping.Colly) fiber.Handler {
	return func(c *fiber.Ctx) error {
		blog, err := service.CreatePost(c, colly)
		if err != nil {
			return nil
		}

		return c.Status(fiber.StatusOK).JSON(blog)
	}
}
