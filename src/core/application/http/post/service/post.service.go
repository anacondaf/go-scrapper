package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/domain/enums"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"gorm.io/gorm"
)

type PostService struct {
	scraper      *webScraping.WebScraper   `di.inject:"webScraper"`
	db           *gorm.DB                  `di.inject:"db"`
	cacheService persistence.ICacheService `di.inject:"cache"`
	producer     *rabbitmq.Producer        `di.inject:"producer"`
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

	err = s.cacheService.Delete(context.Background(), enums.POST_KEY)
	if err != nil {
		return webScraping.Post{}, err
	}

	return post, nil
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	var postsDto []models.Post

	s.db.Preload("PostImages").Find(&postsDto)

	//errChannel := make(chan error)
	//stopCh := make(chan struct{})

	//go func(errChannel chan error, stopCh <-chan struct{}) {
	//
	//	return
	//}(errChannel, stopCh)

	//err := <-errChannel

	//err := s.producer.Publish(context.Background(), "hello", postsDto)
	//if err != nil {
	//	return nil, err
	//}

	return postsDto, nil
}

func (s *PostService) TestMessage() (string, error) {
	err := s.producer.Publish(context.Background(), "hello", "Hello")
	if err != nil {
		return "", err
	}

	return "HelloWorld", nil
}
