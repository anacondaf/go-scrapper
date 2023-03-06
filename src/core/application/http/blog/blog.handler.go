package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/http/blog/service"
)

type Handler struct {
	blogService *service.BlogService `di.inject:"blogService"`
}

func (h *Handler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		blog, err := h.blogService.CreatePost(c)
		if err != nil {
			return nil
		}

		return c.Status(fiber.StatusOK).JSON(blog)
	}
}
