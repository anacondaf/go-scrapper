package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"gorm.io/gorm"
)

type PostService struct {
	scraper *webScraping.WebScraper `di.inject:"webScraper"`
	db      *gorm.DB                `di.inject:"db"`
}

type createPostRequest struct {
	Url string `json:"url"`
}

func (s *PostService) CreatePost(c *fiber.Ctx) (webScraping.Post, error) {

	var url = createPostRequest{}

	err := c.BodyParser(&url)
	if err != nil {
		return webScraping.Post{}, err
	}

	post := s.scraper.VnExpressCrawler(url.Url)

	var postDTO = &models.Post{
		Title:      post.Title,
		PostImages: []models.PostImages{},
	}

	for _, image := range post.Images {
		postDTO.PostImages = append(postDTO.PostImages, models.PostImages{Url: image})
	}

	s.db.Create(postDTO)

	return post, nil
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	var postsDto []models.Post

	s.db.Preload("PostImages").Find(&postsDto)

	return postsDto, nil
}
