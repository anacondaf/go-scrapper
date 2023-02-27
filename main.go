package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/internal/colly"
	"log"
)

type BlogStruct struct {
	Url string `json:"url"`
}

func main() {
	var app = fiber.New()
	var goColly = colly.NewColly()

	app.Post("/blogs", func(c *fiber.Ctx) error {
		var url = BlogStruct{}

		err := c.BodyParser(&url)
		if err != nil {
			return err
		}

		blog := goColly.VnExpressCrawler(url.Url)

		return c.Status(fiber.StatusOK).JSON(blog)
	})

	log.Fatalf("Error when running fiber app: %v", app.Listen(":3000"))
}
