package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

// There should always be a naive implementation that we can test locally
// with no external dependancies.
//
// Do not use this scraper if performance is important, it's not fast.
// We can make a faster one ourselves with raw GoQuery or just use another
// implementation (apify)
type naiveScraper struct {
	*colly.Collector
	url string
}

func (s *naiveScraper) Scrape(req Request) {
	c := s.Clone()
	registerDefaultTriggers(c)
	s.scrape(c, req)
}

func (s *naiveScraper) GoScrape(req Request) {
	c := s.Clone()
	c.Async = true
	registerDefaultTriggers(c)
	s.scrape(c, req)
}

func (s *naiveScraper) scrape(c *colly.Collector, req Request) {
	// This channel allows us to quit scraping prematurely
	quit := make(chan bool)

	// Callbacks
	c.OnRequest(func(r *colly.Request) {
		// If we have exceeded MaxItems then quit scraping
		if req.MaxItems < 0 {
			quit <- true
		}
	})
	c.OnResponse(func(r *colly.Response) {
		req.MaxItems--
	})

	// Start scraping
	go func() {
		c.Visit(s.url)
		quit <- true
	}()

	// This will block until c.Visit() returns or MaxItems is exceeded
	<-quit
}

// This is the 'logic' of the scraper,
// defining the behaviour when elements are found on the page
func registerDefaultTriggers(c *colly.Collector) {
	// Find each job result link
	c.OnHTML(`.result`, func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	// Follow through to next page
	c.OnHTML(`a[aria-label="Next"]`, func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	// Print the Title of the Job
	c.OnHTML(`.jobsearch-JobInfoHeader-title`, func(e *colly.HTMLElement) {
		fmt.Printf("Job Found: %s\n", e.Text)
	})
}

func newNaiveScraper() *naiveScraper {
	scraper := &naiveScraper{
		Collector: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0"),
			// colly.AllowedDomains("indeed.com", "uk.indeed.com"),
			// colly.Debugger(&debug.WebDebugger{}),
			// colly.AllowURLRevisit(),
			// colly.Async(true),
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
