package service

import (
	"github.com/gofiber/fiber/v2"
)

type IBlogService interface {
	CreatePost(c *fiber.Ctx) error
}
