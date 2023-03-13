package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"gorm.io/gorm"
)

type BlogService struct {
	scraper *webScraping.WebScraper `di.inject:"webScraper"`
	db      *gorm.DB                `di.inject:"db"`
}

type createPostRequest struct {
	Url string `json:"url"`
}

func (s *BlogService) CreatePost(c *fiber.Ctx) (webScraping.Post, error) {

	var url = createPostRequest{}

	err := c.BodyParser(&url)
	if err != nil {
		return webScraping.Post{}, err
	}

	post := s.scraper.VnExpressCrawler(url.Url)

	var postDTO = &models.Post{
		Title: post.Title,
		PostImages: []models.PostImages{
			{Url: "Hihi"},
			{Url: "Hoho"},
		},
	}

	//for _, image := range post.Images {
	//	postDTO.PostImages = append(postDTO.PostImages, models.PostImages{Url: image})
	//}

	s.db.Create(postDTO)

	fmt.Printf("%+v\n", postDTO)

	return post, nil
}
