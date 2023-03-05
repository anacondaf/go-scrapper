package main

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/kainguyen/go-scrapper/src/core/handler/blog"
	"github.com/kainguyen/go-scrapper/src/core/route/v1"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"log"
)

func main() {
	var app = fiber.New()
	var goColly = webScraping.NewColly()

	v1 := app.Group("/api/v1")

	blogRouter := v1.Group("/blogs")

	var blogHandler = controller.NewBlogHandler(goColly)
	route.BlogRoute(blogRouter, blogHandler)

	log.Fatalf("Error when running fiber app: %v", app.Listen(":3000"))
}
