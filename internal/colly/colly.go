package colly

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type Colly struct {
	Collector *colly.Collector
}

func NewColly() *Colly {
	collector := colly.NewCollector()
	return &Colly{Collector: collector}
}

// Crawler /*
func (col *Colly) crawler(url string) {
	col.Collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	col.Collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	col.Collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
}

type BlogContent struct {
	Title  string
	Images []string
}

func (col *Colly) VnExpressCrawler(url string) BlogContent {
	col.crawler(url)

	var blog = BlogContent{
		Title:  "",
		Images: []string{},
	}

	col.Collector.OnHTML("section.section.page-detail.top-detail > div.container div.sidebar-1", func(e *colly.HTMLElement) {
		blog.Title = e.ChildText("h1")

		e.ForEach("article figure", func(_ int, elements *colly.HTMLElement) {
			var url = e.ChildAttr("div.fig-picture > picture > img", "data-src")
			blog.Images = append(blog.Images, url)
		})
	})

	col.Collector.Visit(url)

	return blog
}
