package scraper

// Request is likely to change in the future but for now we can overfit it
// to querying indeed.
type Request struct {
	Query    string
	Country  string
	Location string
	MaxItems int
}

// Scraper is a facade for interacting with any kind of scraper manager.
//
// The purpose of a facade is to hide the underlying implementation from the user,
// such that they do not need to know nor care about of it's complexities, allowing
// us to also switch out different implementations in the future with relative ease.
//
// The naive.go implementation simply uses colly to scrape the website ourselves with no
// external dependancies.  Other implementations may use someone elses API like apify or
// scraper.io
type Scraper interface {
	Scrape(Request)
}

type ScraperType string

const (
	SCRAPER_NAIVE = "naive"
	SCRAPER_APIFY = "apify"
)

// ScraperFactory is the tool that allows us to create a new scraper with a different type.
type ScraperFactory interface {
	SetType(ScraperType) ScraperFactory
	New() (Scraper, error)
}
