package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

type BlogService struct {
	webScraper *webScraping.Colly
}

type blogStruct struct {
	Url string `json:"url"`
}

func NewBlogService(webScraper *webScraping.Colly) *BlogService {
	return &BlogService{webScraper: webScraper}
}

func (s *BlogService) CreatePost(c *fiber.Ctx) error {
	var url = blogStruct{}

	err := c.BodyParser(&url)
	if err != nil {
		return err
	}

	blog := s.webScraper.VnExpressCrawler(url.Url)

	return c.Status(fiber.StatusOK).JSON(blog)
}
