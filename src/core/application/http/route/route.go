package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
	"github.com/rs/zerolog"
)

func HttpRoute(routeVer fiber.Router, logger *zerolog.Logger) {
	var postHandler = di.GetInstance("postHandler").(*post.PostHandler)

	postRouter := routeVer.Group("posts")
	postRouter.Post("/", postHandler.CreatePost())
	postRouter.Get("/", postHandler.GetPosts())

	logger.Info().Msg("Setup Endpoint Routing Success!")
}
