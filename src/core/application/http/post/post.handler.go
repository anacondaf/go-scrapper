package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
)

type Handler struct {
	postService *service.PostService `di.inject:"postService"`
}

func (h *Handler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post, err := h.postService.CreatePost(c)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(post)
	}
}

func (h *Handler) GetPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post, err := h.postService.GetPosts()

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(post)
	}
}
