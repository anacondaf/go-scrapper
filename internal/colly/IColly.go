package colly

import "github.com/gocolly/colly"

type IColly interface {
	Crawler(url string) *colly.HTMLElement
}
