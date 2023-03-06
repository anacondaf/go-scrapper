package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

type BlogService struct {
}

type createPostRequest struct {
	Url string `json:"url"`
}

func (s *BlogService) CreatePost(c *fiber.Ctx) (webScraping.BlogContent, error) {
	var url = createPostRequest{}

	err := c.BodyParser(&url)
	if err != nil {
		return webScraping.BlogContent{}, err
	}

	blog := webScraping.VnExpressCrawler(url.Url)
	return blog, nil
}
