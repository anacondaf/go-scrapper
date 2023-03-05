package webScraping

import "github.com/gocolly/colly"

type IColly interface {
	crawler(url string)
	VnExpressCrawler(url string) *colly.HTMLElement
}
