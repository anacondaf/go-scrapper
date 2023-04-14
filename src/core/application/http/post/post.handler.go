package post

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"github.com/kainguyen/go-scrapper/src/core/domain/enums"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
)

type PostHandler struct {
	postService  *service.PostService      `di.inject:"postService"`
	cacheService persistence.ICacheService `di.inject:"cache"`
	producer     *rabbitmq.Producer        `di.inject:"producer"`
}

// Create New Post godoc
//
//	@Summary		Create post
//	@Description	Create post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			post	body	service.createPostRequest	true "URL of the website you want to crawl" example(string)
//	@Success		200	{object}	[]models.Post
//	@Router			/posts [post]
func (h *PostHandler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post, err := h.postService.CreatePost(c)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(post)
	}
}

// Get All Posts godoc
//
//	@Summary		Get all posts
//	@Description	Get all posts
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Post
//	@Failure		500	{object}	error
//	@Router			/posts [get]
func (h *PostHandler) GetPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var postsDto []models.Post

		_, err := h.cacheService.GetOrSet(context.Background(), enums.POST_KEY, 0, &postsDto, persistence.Callback(func(...interface{}) (interface{}, error) {
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
