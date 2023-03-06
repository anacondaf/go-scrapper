package webScraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type BlogContent struct {
	Title  string
	Images []string
}

type WebScraper struct {
	*colly.Collector
}

func NewWebScraper(opts ...func(c *colly.Collector)) *WebScraper {
	collector := colly.NewCollector(opts...)

	return &WebScraper{collector}
}

func (s *WebScraper) VnExpressCrawler(url string) BlogContent {
	s.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	s.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	s.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	var blog = BlogContent{
		Title:  "",
		Images: []string{},
	}

	s.OnHTML("section.section.page-detail.top-detail > div.container div.sidebar-1", func(e *colly.HTMLElement) {
		blog.Title = e.ChildText("h1")

		e.ForEach("article figure", func(_ int, elements *colly.HTMLElement) {
			var url = e.ChildAttr("div.fig-picture > picture > img", "data-src")
			blog.Images = append(blog.Images, url)
		})
	})

	s.Visit(url)

	return blog
}
