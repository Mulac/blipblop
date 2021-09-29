package main

import (
	"fmt"
	"scraper/scraper"
)

func main() {
	fmt.Println("Scraping Jobs in Leeds...")

	scraper.Singleton().Scrape(scraper.ScrapeRequest{
		Country:  "GB",
		Location: "Leeds",
		MaxItems: 10,
	})
}
