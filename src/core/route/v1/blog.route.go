package route

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kainguyen/go-scrapper/src/core/handler/blog"
)

func BlogRoute(blogRouter fiber.Router, blogHandler handler.BlogHandler) {
	blogRouter.Post("/", blogHandler.CreatePost())
}
