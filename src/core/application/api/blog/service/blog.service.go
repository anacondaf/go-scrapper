package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

type blogStruct struct {
	Url string `json:"url"`
}

func CreatePost(c *fiber.Ctx, colly *webScraping.Colly) (webScraping.BlogContent, error) {
	var url = blogStruct{}

	err := c.BodyParser(&url)
	if err != nil {
		return webScraping.BlogContent{}, err
	}

	blog := colly.VnExpressCrawler(url.Url)
	return blog, nil
}
