package scraper

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var scraper Scraper
var once sync.Once

// Singleton returns our 'default' scraper, with the type defined in the environment variable SCRAPER
func Singleton() Scraper {
	once.Do(func() {
		// TODO(calum): decide on a better way to manage configuration
		s, err := NewScraperFactory().SetType(ScraperType(os.Getenv("SCRAPER"))).New()
		if err != nil {
			log.Fatal(err)
		}
		scraper = s
	})

	return scraper
}

type scraperFactory struct {
	stype ScraperType
}

func (sf *scraperFactory) SetType(stype ScraperType) ScraperFactory {
	sf.stype = stype
	return sf
}

// New is responsible for returning a valid scraper, depending on the scraper type
// set in the factory
func (sf *scraperFactory) New() (Scraper, error) {
	switch sf.stype {
	case SCRAPER_NAIVE, "":
		return newNaiveScraper(), nil

	case SCRAPER_APIFY:
		return newApiyScraper(), nil

	default:
		return nil, fmt.Errorf("ERROR|scraperFactory|ScraperType %s not recognised", sf.stype)
	}
}

func NewScraperFactory() ScraperFactory {
	return &scraperFactory{}
}
