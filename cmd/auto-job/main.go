package main

import (
	"fmt"

	"github.com/stephenjhwang/auto-job/pkg/chromedp"
	"github.com/stephenjhwang/auto-job/pkg/scraper"
)

func main() {

	fmt.Println("Hello World!")
	ctx, cancel := chromedp.Initialize()
	defer cancel()

	var s scraper.Scraper
	ls := scraper.LinkedinScraper{"linkedin"}
	s = &ls
	s.Scrape(ctx)
}
