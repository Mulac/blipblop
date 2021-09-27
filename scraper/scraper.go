package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func scrapeIndeedSearch(searchURL string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0"),
		// colly.AllowedDomains("indeed.com", "uk.indeed.com"),
		// colly.Debugger(&debug.WebDebugger{}),
		// colly.AllowURLRevisit(),
	)

	// Add a limiter to spread requests out as to not overload the server
	c.Limit(&colly.LimitRule{
		Delay:       2 * time.Second,
		RandomDelay: 1 * time.Second,
	})

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

	// Callback before making a request
	c.OnRequest(func(r *colly.Request) {})

	// Start scraping
	c.Visit(searchURL)
}

func main() {
	fmt.Println("Scraping Jobs in Leeds...")
	scrapeIndeedSearch("https://uk.indeed.com/jobs?l=Leeds,+West+Yorkshire")
}
