package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/kainguyen/go-scrapper/internal/screenshotter"
	"log"
	"os"
)

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`div#primary`),
		chromedp.FullScreenshot(res, quality),
	}
}

func partlyScreenshot(urlstr string, sel interface{}, buf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.Screenshot(sel, buf),
	}
}

func CaptureDevToWebpage(ctx context.Context, cancelFunc context.CancelFunc) {
	defer cancelFunc()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://dev.to/`, 50, &buf)); err != nil {
		panic(err)
	}

	if err := os.WriteFile("devto_thumbnail.jpeg", buf, 0777); err != nil {
		panic(err)
	}

	log.Printf("screenshot taken!")
}

func CaptureYoutubeWebpage(ctx context.Context, cancelFunc context.CancelFunc) {
	defer cancelFunc()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, partlyScreenshot(`https://www.youtube.com/watch?v=HhanFyk7lTk`, `div#player`, &buf)); err != nil {
		panic(err)
	}

	if err := os.WriteFile("youtube_thumbnail.jpeg", buf, 0777); err != nil {
		panic(err)
	}

	log.Printf("screenshot taken!")
}

func main() {
	newScreenshotterCtx, cancel := screenshotter.NewScreenshotter()

	//CaptureDevToWebpage(newScreenshotterCtx, cancel)
	CaptureYoutubeWebpage(newScreenshotterCtx, cancel)
}
