package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

// @Sam Barnes's Apify API KEY
const API_KEY = "apify_api_P7Tb03JwTcEck64l8f6LkmtyYjPD2S3qWaH6"

type apifyScraper struct {
	apifyUrl    string
	scrapeQueue chan Request
}

func (s apifyScraper) Scrape(req Request) {
	start := time.Now()

	// Map to interface to allow us to store multiple different types
	type Map map[string]interface{}

	// Use this to extend the output of the scraper and grab other data, modify existing data etc.
	extendOutput := "($) => {return {\"companyImage\": $('.jobsearch-CompanyAvatar-image').attr('src')}}"

	// Use Apify proxy settings (current) or we can use our own proxy here
	proxyConfig := Map{"useApifyProxy": true}

	// Data in json format to be sent in POST request
	postbody, _ := json.Marshal(Map{
		"position":             req.Query,
		"country":              req.Country,
		"location":             req.Location,
		"maxItems":             req.MaxItems,
		"maxConcurrency":       10,
		"extendOutputFunction": extendOutput,
		"proxyConfiguration":   proxyConfig})

	responseBody := bytes.NewBuffer(postbody)

	// Post with the responseBody, settings the content type to accept json
	resp, err := http.Post(s.apifyUrl, "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	// Wait for response then close
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Output the response
	sb := string(body)
	fmt.Println(sb)

	// Calculate length of execution (for viability)
	duration := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", duration)
}

func (s apifyScraper) GoScrape(req Request) {
	s.scrapeQueue <- req
}

// scrape is a super dumb worker that picks up requests from the queue and Scrapes them
func (s apifyScraper) scrape() {
	for {
		req := <-s.scrapeQueue
		s.Scrape(req)
	}
}

func (s apifyScraper) runDeamon() {
	// TODO(calum): let the number of workers be set via config
	for i := 0; i < runtime.NumCPU(); i++ {
		// TODO(calum): give each worker a different API key
		go s.scrape()
	}
}

func newApifyScraper() *apifyScraper {
	scraper := &apifyScraper{
		apifyUrl:    "https://api.apify.com/v2/acts/hynekhruska~indeed-scraper/run-sync-get-dataset-items?token=" + API_KEY,
		scrapeQueue: make(chan Request, 10000),
	}

	scraper.runDeamon()
	return scraper
}
