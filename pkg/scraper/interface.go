package scraper

import "context"

type Scraper interface {
	Scrape(context.Context)
}
