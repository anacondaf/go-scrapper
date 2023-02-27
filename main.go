package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/internal/colly"
)

func main() {
	var app = fiber.New()

	app.Get("youtube/thumbnail", func(c *fiber.Ctx) error {

		type Student struct {
			Name string
		}

		student, err := json.Marshal(Student{
			Name: "A",
		})
		if err != nil {
			return err
		}

		return c.Send(student)
	})

	app.Post("youtube/thumbnail/upload", func(c *fiber.Ctx) error {
		var colly = colly.NewColly()

		type Body struct {
			Url string
		}

		var _body = Body{}

		err := json.Unmarshal(c.Body(), &_body)
		if err != nil {
			return err
		}

		colly.Crawler(_body.Url)
		//
		//fmt.Printf("%+v\n", crawler.Attr("src"))

		type Response struct {
			ImageUrl string
		}

		response, _ := json.Marshal(Response{})
		return c.Send(response)
	})

	app.Listen(":3000")
}
