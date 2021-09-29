package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

// There should always be a naive implementation that we can test locally
// with no external dependancies.
type naiveScraper struct {
	*colly.Collector
	url string
}

func (s *naiveScraper) Scrape(req Request) {
	// Find each job result link
	s.OnHTML(`.result`, func(e *colly.HTMLElement) {
		s.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	// Follow through to next page
	s.OnHTML(`a[aria-label="Next"]`, func(e *colly.HTMLElement) {
		s.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	// Print the Title of the Job
	s.OnHTML(`.jobsearch-JobInfoHeader-title`, func(e *colly.HTMLElement) {
		fmt.Printf("Job Found: %s\n", e.Text)
	})

	// Callback before making a request
	s.OnRequest(func(r *colly.Request) {})

	// Start scraping
	s.Visit(s.url)
}

func newNaiveScraper() *naiveScraper {
	scraper := &naiveScraper{
		Collector: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0"),
			// colly.AllowedDomains("indeed.com", "uk.indeed.com"),
			// colly.Debugger(&debug.WebDebugger{}),
			// colly.AllowURLRevisit(),
		),

		url: "https://uk.indeed.com/jobs?l=Leeds,+West+Yorkshire",
	}

	// Add a limiter to spread requests out as to not overload the server
	scraper.Limit(&colly.LimitRule{
		DomainGlob:  "*indeed.com*",
		Delay:       2 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	return scraper
}
