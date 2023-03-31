package route

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
)

func Route(routeVer fiber.Router) {
	var postHandler = di.GetInstance("postHandler").(*post.PostHandler)

	postRouter := routeVer.Group("posts")
	postRouter.Post("/", postHandler.CreatePost())
	postRouter.Get("/", postHandler.GetPosts())

	fmt.Println("Setup Endpoint Routing Success!")
}
