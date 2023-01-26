package scraper

import (
	"context"
	"fmt"
)

type Scraper struct {
	name string
}

func (s *Scraper) Scrape(ctx context.Context) {
	fmt.Println("Hello world from", s.name)
}
