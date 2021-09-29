package main

import (
	"fmt"
	"scraper/scraper"
)

func main() {
	fmt.Println("Scraping Jobs in Leeds...")

	// The first place to look is in scraper/interface.go as this defines how we will
	// interact with the scraper
	scraper.Singleton().Scrape(scraper.ScrapeRequest{
		Country:  "GB",
		Location: "Leeds",
		MaxItems: 10,
	})
}
