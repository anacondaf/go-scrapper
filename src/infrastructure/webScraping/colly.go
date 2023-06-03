package webScraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/rs/zerolog"
)

type Post struct {
	Title  string
	Images []string
}

type WebScraper struct {
	*colly.Collector
	logger *zerolog.Logger
}

func NewWebScraper(config *config.Config, logger *zerolog.Logger, opts ...func(c *colly.Collector)) *WebScraper {
	collector := colly.NewCollector(opts...)

	//storage := &postgres.Storage{
	//	URI:          config.DBUrl,
	//	VisitedTable: "colly_visited",
	//	CookiesTable: "colly_cookies",
	//}
	//
	//collector.SetStorage(storage)

	return &WebScraper{collector, logger}
}

func (s *WebScraper) VnExpressCrawler(url string) Post {
	s.OnRequest(func(request *colly.Request) {
		s.logger.Info().Msg(fmt.Sprintf("Visiting %v", request.URL))
	})

	s.OnError(func(_ *colly.Response, err error) {
		s.logger.Fatal().Err(err).Msg("Something Went Wrong")
	})

	s.OnResponse(func(r *colly.Response) {
		s.logger.Info().Msg(fmt.Sprintf("Visited %v", r.Request.URL))
	})

	var blog = Post{
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
