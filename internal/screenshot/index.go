package screenshot

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func newScreenshotter() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[3:], chromedp.NoFirstRun, chromedp.NoDefaultBrowserCheck)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	return chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
}

func fullScreenshot(urlstr string, sel interface{}, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
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

func CaptureDevToWebpage() {
	screenshotCtx, cancel := newScreenshotter()

	defer cancel()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(screenshotCtx, fullScreenshot(`https://dev.to/`, "footer", 50, &buf)); err != nil {
		panic(err)
	}

	if err := os.WriteFile("devto_thumbnail.jpeg", buf, 0777); err != nil {
		panic(err)
	}

	log.Printf("screenshot taken!")
}

func CaptureYoutubeWebpage() {
	screenshotCtx, cancel := newScreenshotter()

	defer cancel()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(screenshotCtx, partlyScreenshot(`https://www.youtube.com/watch?v=HhanFyk7lTk`, `div#player`, &buf)); err != nil {
		panic(err)
	}

	if err := os.WriteFile("youtube_thumbnail.jpeg", buf, 0777); err != nil {
		panic(err)
	}

	log.Printf("screenshot taken!")
}
