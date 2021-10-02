package main

import (
	"blipblop/src/scraper"
	"fmt"
)

func main() {
	fmt.Println("Scraping Jobs in Leeds...")

	// The first place to look is in scraper/interface.go
	scraper.Singleton().Scrape(scraper.Request{
		Country:  "GB",
		Location: "Leeds",
		MaxItems: 10,
	})
}
