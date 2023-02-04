package main

import "github.com/gofiber/fiber/v2"

func main() {
	//screenshot.CaptureYoutubeWebpage()
	//screenshot.CaptureDevToWebpage()

	//c := colly.NewCollector()
	//
	//// Before making a request print "Visiting ..."
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL.String())
	//})
	//
	//c.OnResponse(func(response *colly.Response) {
	//	fmt.Println(string(response.Body))
	//})
	//
	//// On every a element which has href attribute call callback
	////c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	////	link := e.Attr("href")
	////	// Print link
	////	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	////	// Visit link found on page
	////	// Only those links are visited which are in AllowedDomains
	////	//c.Visit(e.Request.AbsoluteURL(link))
	////})
	//
	//err := c.Visit("https://hackerspaces.org/")
	//if err != nil {
	//	return
	//}

	var app = fiber.New()

	app.Get("")

}
