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
func (col *Colly) Crawler(url string) *colly.HTMLElement {
	col.Collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	col.Collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	col.Collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	var htmlElems *colly.HTMLElement

	col.Collector.OnHTML("div.crayons-article__main > div[data-article-id=\"1371452\"]", func(e *colly.HTMLElement) {
		fmt.Println(len(e.ChildAttrs("img", "src")))
	})

	col.Collector.Visit(url)

	return htmlElems
}
