package screenshotter

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func NewScreenshotter() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[3:], chromedp.NoFirstRun, chromedp.NoDefaultBrowserCheck)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return taskCtx, cancel
}
