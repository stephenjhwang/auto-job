package scraper

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type LinkedinScraper struct {
	Name string
}

func (ls *LinkedinScraper) Scrape(ctx context.Context) {
	fmt.Println("Hello world from", ls.Name)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.linkedin.com/jobs/search/?geoId=102221843&keywords=software engineer`),
		chromedp.Text(`.jobs-search-results-list li`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))

}
