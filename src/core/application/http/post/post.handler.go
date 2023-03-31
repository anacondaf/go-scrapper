package post

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
)

type PostHandler struct {
	postService  *service.PostService      `di.inject:"postService"`
	cacheService persistence.ICacheService `di.inject:"cache"`
	producer     *rabbitmq.Producer        `di.inject:"producer"`
}

func (h *PostHandler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post, err := h.postService.CreatePost(c)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(post)
	}
}

func (h *PostHandler) GetPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var postsDto []models.Post

		_, err := h.cacheService.GetOrSet(context.Background(), "posts", 0, &postsDto, persistence.Callback(func(...interface{}) (interface{}, error) {
			post, err := h.postService.GetPosts()
			if err != nil {
				return nil, err
			}

			return post, nil
		}))

		if err != nil {
			return err
		}

		err = h.producer.Publish(context.Background(), "hello", rabbitmq.NewMessage("post", postsDto))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		err = h.producer.Publish(context.Background(), "hello", rabbitmq.NewMessage("hello_message", postsDto))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(postsDto)
	}
}
