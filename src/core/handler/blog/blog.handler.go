package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/api/services/blog"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

type BlogHandler struct {
	blogService *service.BlogService
}

func NewBlogHandler(webScraper *webScraping.Colly) BlogHandler {
	blogService := service.NewBlogService(webScraper)

	return BlogHandler{blogService: blogService}
}

func (controller *BlogHandler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controller.blogService.CreatePost(c)
	}
}
