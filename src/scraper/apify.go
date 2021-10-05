package scraper

import (
	"blipblop/src/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	// Salary and Contract type
	//jobsearch-JobMetadataHeader-item
	// Use this to extend the output of the scraper and grab other data, modify existing data etc.
	extendOutput := "($) => {return {\"companyImage\": $('.jobsearch-CompanyAvatar-image').attr('src'), \"metadata\": $('.jobsearch-JobMetadataHeader-item > span').text()}}"

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
		fmt.Printf("ERROR|apifyScraper.Scrape(%+v)|failed to make request to apify|%v", req, err)
		return
	}

	// Wait for response then close
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR|apifyScraper.Scrape(%+v)|failed to read response from apify|%v", req, err)
		return
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(string(body))
		fmt.Printf("Elapsed time: %s\n", time.Since(start))
		return
	}

	var jobs []storage.Job
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		fmt.Printf("ERROR|apifyScraper.Scrape(%+v)|failed to unmarshal indeed response to map[string]interface{}|%v", req, err)
		return
	}

	storage.DB().AddJob(jobs...)
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
