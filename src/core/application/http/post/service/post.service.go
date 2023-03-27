package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostService struct {
	scraper      *webScraping.WebScraper `di.inject:"webScraper"`
	db           *gorm.DB                `di.inject:"db"`
	cacheService *redis.Client           `di.inject:"cache"`
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

	ctx := context.Background()

	_, err := s.cacheService.Get(ctx, "posts").Result()

	switch {
	case err == redis.Nil:
		{
			fmt.Println("cache is empty")

			s.db.Preload("PostImages").Find(&postsDto)

			s.cacheService.HMSet(ctx, "posts", map[string]interface{}{
				"key1": "value2",
				"key2": "value2",
			})
		}
	case err != nil:
		{
			fmt.Println("Get failed", err)
			return nil, err
		}
	}

	return postsDto, nil
}
